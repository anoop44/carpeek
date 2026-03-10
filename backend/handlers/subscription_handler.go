package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

// RevenueCat webhook event structure
// https://www.revenuecat.com/docs/integrations/webhooks/event-flows
type rcWebhookPayload struct {
	APIVersion string  `json:"api_version"`
	Event      rcEvent `json:"event"`
}

type rcEvent struct {
	ID                       string  `json:"id"`
	Type                     string  `json:"type"`
	AppUserID                string  `json:"app_user_id"`
	ProductID                string  `json:"product_id"`
	EntitlementIDs           []string `json:"entitlement_ids"`
	PeriodType               string  `json:"period_type"`
	ExpirationAtMs           *int64  `json:"expiration_at_ms"`
	IsTrialConversion        bool    `json:"is_trial_conversion"`
	Environment              string  `json:"environment"`
	PresentedOfferingID      string  `json:"presented_offering_id"`
	Currency                 string  `json:"currency"`
	PriceInPurchasedCurrency float64 `json:"price_in_purchased_currency"`
}

// GetSubscriptionStatusHandler returns the user's subscription status
func GetSubscriptionStatusHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	user, err := GetUserFromRequest(w, r, sqlxDB)
	if err != nil {
		jsonError(w, "Authentication required", http.StatusUnauthorized)
		return
	}

	response := models.SubscriptionStatusResponse{
		IsSubscriber: user.IsSubscriber,
	}

	if user.SubscriptionStatus != nil {
		response.SubscriptionStatus = user.SubscriptionStatus
	}

	if user.SubscriptionProvider != nil {
		response.SubscriptionProvider = user.SubscriptionProvider
	}

	if user.SubscriptionExpiresAt != nil {
		formatted := user.SubscriptionExpiresAt.Format(time.RFC3339)
		response.ExpiresAt = &formatted
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RevenueCatWebhookHandler handles incoming RevenueCat webhook events
func RevenueCatWebhookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("RCWebhook", fmt.Errorf("database connection not available"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Verify webhook authorization header
	webhookSecret := os.Getenv("REVENUECAT_WEBHOOK_SECRET")
	if webhookSecret != "" {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer "+webhookSecret {
			utils.LogError("RCWebhook.Auth", fmt.Errorf("invalid webhook authorization"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	// Parse webhook payload
	var payload rcWebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.LogError("RCWebhook.Parse", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := payload.Event

	utils.LogEvent("REVENUECAT", "Webhook received", map[string]interface{}{
		"type":        event.Type,
		"id":          event.ID,
		"app_user_id": event.AppUserID,
		"product_id":  event.ProductID,
		"environment": event.Environment,
	})

	// Skip sandbox events in production (optional)
	// if event.Environment == "SANDBOX" { w.WriteHeader(http.StatusOK); return }

	switch event.Type {
	case "INITIAL_PURCHASE", "RENEWAL", "UNCANCELLATION":
		handleSubscriptionActive(sqlxDB, event)

	case "CANCELLATION":
		// User canceled but still has access until period end
		handleSubscriptionCanceled(sqlxDB, event)

	case "EXPIRATION":
		// Subscription actually expired — remove access
		handleSubscriptionExpired(sqlxDB, event)

	case "BILLING_ISSUE_DETECTED":
		handleBillingIssue(sqlxDB, event)

	case "PRODUCT_CHANGE":
		handleSubscriptionActive(sqlxDB, event)

	default:
		utils.LogEvent("REVENUECAT", "Unhandled event type", map[string]interface{}{
			"type": event.Type,
		})
	}

	w.WriteHeader(http.StatusOK)
}

// --- Internal handlers ---

func handleSubscriptionActive(db *sqlx.DB, event rcEvent) {
	var expiresAt *time.Time
	if event.ExpirationAtMs != nil {
		t := time.Unix(*event.ExpirationAtMs/1000, 0)
		expiresAt = &t
	}

	err := database.UpdateUserSubscription(db, event.AppUserID, event.ProductID, "active", "revenuecat", expiresAt)
	if err != nil {
		utils.LogError("RCWebhook.ActivateSubscription", err)
		return
	}

	// Log event
	user, err := database.GetUserByAnonymousID(db, event.AppUserID)
	if err == nil {
		database.LogSubscriptionEvent(db, user.ID, event.Type, "revenuecat", event.ID,
			fmt.Sprintf("Product: %s, Period: %s", event.ProductID, event.PeriodType))
	}

	utils.LogEvent("REVENUECAT", "Subscription activated", map[string]interface{}{
		"app_user_id": event.AppUserID,
		"product_id":  event.ProductID,
	})
}

func handleSubscriptionCanceled(db *sqlx.DB, event rcEvent) {
	// User canceled but still has access until expiration
	var expiresAt *time.Time
	if event.ExpirationAtMs != nil {
		t := time.Unix(*event.ExpirationAtMs/1000, 0)
		expiresAt = &t
	}

	// Keep is_subscriber = true (access until period end), but mark status as canceled
	err := database.UpdateUserSubscription(db, event.AppUserID, event.ProductID, "canceled", "revenuecat", expiresAt)
	if err != nil {
		utils.LogError("RCWebhook.CancelSubscription", err)
		return
	}

	// Override: user still has access until expiry, keep them as subscriber
	// The is_subscriber will be set to false only on EXPIRATION event
	if expiresAt != nil && expiresAt.After(time.Now()) {
		_, execErr := db.Exec(`UPDATE users SET is_subscriber = TRUE WHERE anonymous_id = $1`, event.AppUserID)
		if execErr != nil {
			utils.LogError("RCWebhook.KeepSubscriberActive", execErr)
		}
	}

	user, err := database.GetUserByAnonymousID(db, event.AppUserID)
	if err == nil {
		database.LogSubscriptionEvent(db, user.ID, "canceled", "revenuecat", event.ID,
			fmt.Sprintf("Product: %s, Expires: %v", event.ProductID, expiresAt))
	}

	utils.LogEvent("REVENUECAT", "Subscription canceled (still active until expiry)", map[string]interface{}{
		"app_user_id": event.AppUserID,
		"expires_at":  expiresAt,
	})
}

func handleSubscriptionExpired(db *sqlx.DB, event rcEvent) {
	err := database.DeactivateUserSubscription(db, event.AppUserID)
	if err != nil {
		utils.LogError("RCWebhook.ExpireSubscription", err)
		return
	}

	user, err := database.GetUserByAnonymousID(db, event.AppUserID)
	if err == nil {
		database.LogSubscriptionEvent(db, user.ID, "expired", "revenuecat", event.ID,
			fmt.Sprintf("Product: %s", event.ProductID))
	}

	utils.LogEvent("REVENUECAT", "Subscription expired", map[string]interface{}{
		"app_user_id": event.AppUserID,
	})
}

func handleBillingIssue(db *sqlx.DB, event rcEvent) {
	var expiresAt *time.Time
	if event.ExpirationAtMs != nil {
		t := time.Unix(*event.ExpirationAtMs/1000, 0)
		expiresAt = &t
	}

	err := database.UpdateUserSubscription(db, event.AppUserID, event.ProductID, "past_due", "revenuecat", expiresAt)
	if err != nil {
		utils.LogError("RCWebhook.BillingIssue", err)
		return
	}

	user, err := database.GetUserByAnonymousID(db, event.AppUserID)
	if err == nil {
		database.LogSubscriptionEvent(db, user.ID, "billing_issue", "revenuecat", event.ID,
			fmt.Sprintf("Product: %s", event.ProductID))
	}
}

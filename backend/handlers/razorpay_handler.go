package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	razorpay "github.com/razorpay/razorpay-go"
)

// Razorpay webhook event structure
type razorpayWebhookPayload struct {
	Entity    string                 `json:"entity"`
	AccountID string                 `json:"account_id"`
	Event     string                 `json:"event"`
	Contains  []string               `json:"contains"`
	Payload   razorpayPayloadContent `json:"payload"`
	CreatedAt int64                  `json:"created_at"`
}

type razorpayPayloadContent struct {
	Subscription struct {
		Entity razorpaySubscriptionEntity `json:"entity"`
	} `json:"subscription"`
	Payment struct {
		Entity razorpayPaymentEntity `json:"entity"`
	} `json:"payment"`
}

type razorpaySubscriptionEntity struct {
	ID                 string                 `json:"id"`
	Entity             string                 `json:"entity"`
	PlanID             string                 `json:"plan_id"`
	CustomerID         string                 `json:"customer_id"`
	Status             string                 `json:"status"`
	CurrentStart       int64                  `json:"current_start"`
	CurrentEnd         int64                  `json:"current_end"`
	EndedAt            int64                  `json:"ended_at"`
	ChargeAt           int64                  `json:"charge_at"`
	StartAt            int64                  `json:"start_at"`
	EndAt              int64                  `json:"end_at"`
	AuthAttempts       int                    `json:"auth_attempts"`
	TotalCount         int                    `json:"total_count"`
	PaidCount          int                    `json:"paid_count"`
	RemainingCount     int                    `json:"remaining_count"`
	ShortURL           string                 `json:"short_url"`
	HasScheduledChange bool                   `json:"has_scheduled_changes"`
	ChangeScheduledAt  int64                  `json:"change_scheduled_at"`
	OfferID            string                 `json:"offer_id"`
	Notes              map[string]interface{} `json:"notes"`
	CreatedAt          int64                  `json:"created_at"`
}

type razorpayPaymentEntity struct {
	ID             string `json:"id"`
	Entity         string `json:"entity"`
	OrderID        string `json:"order_id"`
	SubscriptionID string `json:"subscription_id"`
}

// CreateRazorpaySubscriptionHandler creates a new subscription in Razorpay
func CreateRazorpaySubscriptionHandler(w http.ResponseWriter, r *http.Request) {
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

	keyID := os.Getenv("RAZORPAY_KEY_ID")
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")
	planID := os.Getenv("RAZORPAY_PLAN_ID")

	if keyID == "" || keySecret == "" || planID == "" {
		utils.LogError("Razorpay.Config", fmt.Errorf("Razorpay environment variables missing"))
		jsonError(w, "Subscription service not configured", http.StatusInternalServerError)
		return
	}

	client := razorpay.NewClient(keyID, keySecret)

	data := map[string]interface{}{
		"plan_id":      planID,
		"total_count":  120, // 10 years for monthly
		"quantity":     1,
		"customer_notify": 1,
		"notes": map[string]interface{}{
			"anonymous_id": user.AnonymousID,
		},
	}

	body, err := client.Subscription.Create(data, nil)
	if err != nil {
		utils.LogError("Razorpay.CreateSub", err)
		jsonError(w, "Failed to create subscription", http.StatusInternalServerError)
		return
	}

	// The SDK returns body as map[string]interface{} or similar
	subID, ok := body["id"].(string)
	if !ok {
		utils.LogError("Razorpay.ParseID", fmt.Errorf("failed to parse subscription ID from response"))
		jsonError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := models.CreateRazorpaySubscriptionResponse{
		SubscriptionID: subID,
		KeyID:          keyID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RazorpayWebhookHandler handles incoming Razorpay webhook events
func RazorpayWebhookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("RazorpayWebhook", fmt.Errorf("database connection not available"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.LogError("RazorpayWebhook.Read", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify signature
	webhookSecret := os.Getenv("RAZORPAY_WEBHOOK_SECRET")
	signature := r.Header.Get("X-Razorpay-Signature")

	if webhookSecret != "" {
		if !verifyRazorpaySignature(body, signature, webhookSecret) {
			utils.LogError("RazorpayWebhook.Auth", fmt.Errorf("invalid webhook signature"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Parse payload
	var payload razorpayWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.LogError("RazorpayWebhook.Parse", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sub := payload.Payload.Subscription.Entity
	
	// Extract anonymous_id from notes
	anonymousID := ""
	if sub.Notes != nil {
		if val, ok := sub.Notes["anonymous_id"].(string); ok {
			anonymousID = val
		}
	}

	if anonymousID == "" {
		utils.LogEvent("RAZORPAY", "Webhook received but no anonymous_id in notes", map[string]interface{}{
			"subscription_id": sub.ID,
			"event":           payload.Event,
		})
		w.WriteHeader(http.StatusOK) // Return 200 anyway to stop retries
		return
	}

	utils.LogEvent("RAZORPAY", "Webhook received", map[string]interface{}{
		"event":           payload.Event,
		"subscription_id": sub.ID,
		"anonymous_id":    anonymousID,
	})

	switch payload.Event {
	case "subscription.authenticated", "subscription.activated", "subscription.charged":
		handleRazorpayActive(sqlxDB, anonymousID, sub)
	case "subscription.cancelled", "subscription.expired":
		handleRazorpayDeactivated(sqlxDB, anonymousID, sub)
	}

	w.WriteHeader(http.StatusOK)
}

func verifyRazorpaySignature(body []byte, signature, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return expectedSignature == signature
}

func handleRazorpayActive(db *sqlx.DB, anonymousID string, sub razorpaySubscriptionEntity) {
	var expiresAt *time.Time
	if sub.CurrentEnd != 0 {
		t := time.Unix(sub.CurrentEnd, 0)
		expiresAt = &t
	}

	err := database.UpdateUserSubscription(db, anonymousID, sub.PlanID, sub.Status, "razorpay", expiresAt)
	if err != nil {
		utils.LogError("Razorpay.ActivateSubscription", err)
		return
	}

	user, err := database.GetUserByAnonymousID(db, anonymousID)
	if err == nil {
		database.LogSubscriptionEvent(db, user.ID, "activated", "razorpay", sub.ID,
			fmt.Sprintf("Plan: %s, Next Charge: %v", sub.PlanID, expiresAt))
	}
}

func handleRazorpayDeactivated(db *sqlx.DB, anonymousID string, sub razorpaySubscriptionEntity) {
	err := database.DeactivateUserSubscription(db, anonymousID)
	if err != nil {
		utils.LogError("Razorpay.ExpireSubscription", err)
		return
	}

	user, err := database.GetUserByAnonymousID(db, anonymousID)
	if err == nil {
		database.LogSubscriptionEvent(db, user.ID, "expired", "razorpay", sub.ID,
			fmt.Sprintf("Plan: %s", sub.PlanID))
	}
}

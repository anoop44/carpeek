package handlers

import (
	"carpeek-backend/database"
	"carpeek-backend/models"
	"carpeek-backend/utils"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// GetFullImageURL constructs the full URL for a given image path
func GetFullImageURL(imagePath string) string {
	if imagePath == "" {
		return ""
	}

	// If it's already a full URL, return it
	if strings.HasPrefix(imagePath, "http://") || strings.HasPrefix(imagePath, "https://") {
		return imagePath
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		// Fallback to relative path if BASE_URL is not set
		// But in this task we want full URLs
		return imagePath
	}

	// Remove leading slash from imagePath if present
	imagePath = strings.TrimPrefix(imagePath, "/")

	// Ensure baseURL ends with a slash
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	return baseURL + imagePath
}

const (
	UserCookieName = "carpeek_uid"
	UserHeaderName = "X-User-ID"
)

// GetUserFromRequest extracts user identity from cookie or header, or creates a new one
func GetUserFromRequest(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (*models.User, error) {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "carpeek-default-secret-change-me"
	}

	var anonymousID string
	var found bool

	// 1. Check cookie
	cookie, err := r.Cookie(UserCookieName)
	if err == nil {
		val, ok := utils.VerifySignature(cookie.Value, secret)
		if ok {
			anonymousID = val
			found = true
		}
	}

	// 2. Fallback to header (localStorage mirror)
	if !found {
		headerVal := r.Header.Get(UserHeaderName)
		if headerVal != "" {
			val, ok := utils.VerifySignature(headerVal, secret)
			if ok {
				anonymousID = val
				found = true
			}
		}
	}

	// 3. Create new if not found
	if !found {
		anonymousID = uuid.New().String()
	}

	// 4. Get or create in database
	user, err := database.GetOrCreateUser(db, anonymousID)
	if err != nil {
		return nil, err
	}

	// 5. Set/Refresh cookie
	signedID := utils.SignValue(anonymousID, secret)
	http.SetCookie(w, &http.Cookie{
		Name:     UserCookieName,
		Value:    signedID,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   365 * 24 * 60 * 60, // 1 year
	})

	// Also set a header for the frontend to store in localStorage if needed
	w.Header().Set(UserHeaderName, signedID)

	return user, nil
}

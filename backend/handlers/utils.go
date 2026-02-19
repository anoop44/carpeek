package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/middleware"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
	_ "time/tzdata"
)

// jsonError writes a JSON error response with the given message and HTTP status code.
// This ensures all API responses are consistently JSON-formatted.
func jsonError(w http.ResponseWriter, message string, statusCode int) {
	utils.JSONError(w, message, statusCode)
}

// jsonErrorf writes a formatted JSON error response.
func jsonErrorf(w http.ResponseWriter, statusCode int, format string, args ...interface{}) {
	utils.JSONErrorf(w, statusCode, format, args...)
}

// GetSecondsUntilNextChallenge calculates seconds until the next midnight in the given timezone
func GetSecondsUntilNextChallenge(timezone string) int64 {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.UTC
	}

	now := time.Now().In(loc)
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, loc)

	return int64(nextMidnight.Sub(now).Seconds())
}

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

// GetUserFromRequest extracts user identity from JWT (context), cookie or header, or creates a new one
func GetUserFromRequest(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (*models.User, error) {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "carpeek-default-secret-change-me"
	}

	var anonymousID string
	var found bool

	// 0. Check Context (set by JWT middleware)
	if ctxUserID, ok := r.Context().Value(middleware.UserIDKey).(string); ok && ctxUserID != "" {
		anonymousID = ctxUserID
		found = true
	}

	// 1. Check cookie if not found in context
	if !found {
		cookie, err := r.Cookie(UserCookieName)
		if err == nil {
			val, ok := utils.VerifySignature(cookie.Value, secret)
			if ok {
				anonymousID = val
				found = true
			}
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
		utils.LogEvent("USER", "Created new anonymous user", anonymousID)
	}

	// 4. Get or create in database
	user, err := database.GetOrCreateUser(db, anonymousID)
	if err != nil {
		utils.LogError("GetOrCreateUser", err)
		return nil, err
	}

	// 5. Set/Refresh cookie (only if we're not just using JWT, or maybe always to keep cookie alive)
	// If it came from JWT, we might not need to set cookie, but harmless to refresh.
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

package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/middleware"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"time"
	_ "time/tzdata"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

// GetModelImageURL constructs the full URL for a car model image
// Model images are stored in images/actual/ directory
func GetModelImageURL(imagePath string) string {
	if imagePath == "" {
		return ""
	}

	// Model image_url values are relative paths like "bmw/file.jpg"
	// They are stored in images/actual/ directory
	if !strings.HasPrefix(imagePath, "actual/") {
		imagePath = "actual/" + imagePath
	}

	return GetFullImageURL(imagePath)
}

// GetFullImageURL constructs the full URL for a given image path
func GetFullImageURL(imagePath string, challengeID ...int) string {
	if imagePath == "" {
		return ""
	}

	// If it's a challenge image and we have an ID, obfuscate the URL
	if len(challengeID) > 0 {
		// Frontend proxies /api/... to backend /api/v1/...
		return fmt.Sprintf("/api/challenge/image/%d", challengeID[0])
	}

	// If it's already a full URL, return it
	if strings.HasPrefix(imagePath, "http://") || strings.HasPrefix(imagePath, "https://") {
		return imagePath
	}

	baseURL := os.Getenv("BASE_URL")

	// Ensure baseURL ends with a slash if set
	if baseURL != "" && !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	// Remove leading slash from imagePath if present
	imagePath = strings.TrimPrefix(imagePath, "/")

	// Ensure the path starts with "images/" since that's our mount point
	if !strings.HasPrefix(imagePath, "images/") {
		imagePath = "images/" + imagePath
	}

	if baseURL == "" {
		return "/api/" + imagePath
	}

	return baseURL + imagePath
}

const (
	UserCookieName         = "carpeek_uid"
	UserHeaderName         = "X-User-ID"
	BrowserSignatureHeader = "X-Browser-Signature"
)

// GetUserFromRequest extracts user identity ONLY from JWT context (set by RequireAuth middleware)
func GetUserFromRequest(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (*models.User, error) {
	// Identity MUST come from JWT context set by RequireAuth middleware
	ctxUserID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || ctxUserID == "" {
		return nil, fmt.Errorf("user not authenticated in context")
	}

	// Get from database
	user, err := database.GetUserByAnonymousID(db, ctxUserID)
	if err != nil {
		utils.LogError("GetUserByAnonymousID", err)
		return nil, err
	}

	return user, nil
}

// GetUserBySignature identifies or creates a user based on the raw browser signature.
// This is ONLY used by the auth/session endpoint during handshake.
func GetUserBySignature(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (*models.User, error) {
	signature := r.Header.Get(BrowserSignatureHeader)
	if signature == "" {
		return nil, fmt.Errorf("browser signature missing")
	}

	// Validate signature format (should be UUID)
	if _, err := uuid.Parse(signature); err != nil {
		return nil, fmt.Errorf("invalid browser signature format")
	}

	// Get or create in database
	user, err := database.GetOrCreateUser(db, signature)
	if err != nil {
		utils.LogError("GetOrCreateUser", err)
		return nil, err
	}

	// Set a legacy header or cookie if needed, but primarily we rely on JWT now
	w.Header().Set(UserHeaderName, signature)

	return user, nil
}

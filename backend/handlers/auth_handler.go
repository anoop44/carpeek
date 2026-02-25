package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// GoogleTokenInfo represents the response from Google's tokeninfo endpoint
type GoogleTokenInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"` // returned as string "true"/"false" by some endpoints, check actual response
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Error         string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// SessionResponse represents the response for session creation
type SessionResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"` // seconds
	UserID    string `json:"user_id"`
}

// GetSessionHandler handles the request to get a short-lived session token
func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("GetSessionHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get or create user based on Browser Signature (handshake)
	user, err := GetUserBySignature(w, r, sqlxDB)
	if err != nil {
		utils.LogError("GetSessionHandler.Handshake", err)
		jsonError(w, "Session handshake failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate short-lived JWT
	token, err := utils.GenerateShortLivedToken(user.AnonymousID)
	if err != nil {
		utils.LogError("GetSessionHandler.GenerateToken", err)
		jsonError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := SessionResponse{
		Token:     token,
		ExpiresIn: 300, // 5 minutes
		UserID:    user.AnonymousID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GoogleLoginHandler handles the Google login request
func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("GoogleLoginHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Parse request body
	var req models.GoogleLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.LogError("DecodeGoogleLoginRequest", err)
		jsonError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log incoming request
	utils.LogDebug("AUTH", "Request body", map[string]string{"anonymousID": req.AnonymousID})

	if req.IDToken == "" {
		jsonError(w, "ID token is required", http.StatusBadRequest)
		return
	}

	// Verify ID token with Google
	tokenInfo, err := verifyGoogleToken(req.IDToken)
	if err != nil {
		utils.LogError("GoogleTokenVerification", err)
		jsonError(w, "Invalid Google token", http.StatusUnauthorized)
		return
	}

	// Connect to DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get current anonymous ID if provided, otherwise fail? 
	// The prompt says "convert the current anonymous user", so anonymous_id is expected.
	if req.AnonymousID == "" {
		jsonError(w, "Anonymous ID is required", http.StatusBadRequest)
		return
	}

	// Link account or login
	user, err := database.LinkGoogleAccount(sqlxDB, req.AnonymousID, tokenInfo.Sub, tokenInfo.Email, tokenInfo.Name, tokenInfo.Picture)
	if err != nil {
		utils.LogError("LinkGoogleAccount", err)
		jsonError(w, "Failed to login/link account", http.StatusInternalServerError)
		return
	}

	// Log the login event
	utils.LogEvent("AUTH", "User logged in with Google", map[string]interface{}{
		"userID": user.ID,
		"email":  tokenInfo.Email,
	})

	// Also generate a session token for immediate use
	token, _ := utils.GenerateShortLivedToken(user.AnonymousID)
	
	// Set auth token header
	w.Header().Set("X-Auth-Token", token)
	w.Header().Set("Content-Type", "application/json")
	
	// Create expanded response with token
	response := struct {
		*models.User
		AuthToken string `json:"auth_token"`
	}{
		User:      user,
		AuthToken: token,
	}
	
	json.NewEncoder(w).Encode(response)
}

// verifyGoogleToken verifies the ID token using Google's tokeninfo endpoint
func verifyGoogleToken(idToken string) (*GoogleTokenInfo, error) {
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken)
	if err != nil {
		return nil, fmt.Errorf("failed to call google api: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("google api returned status: %d", resp.StatusCode)
	}

	var tokenInfo GoogleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to decode token info: %v", err)
	}

	if tokenInfo.Error != "" {
		return nil, fmt.Errorf("token error: %s", tokenInfo.ErrorDescription)
	}

	if tokenInfo.Sub == "" {
		return nil, fmt.Errorf("token missing subject")
	}

	return &tokenInfo, nil
}

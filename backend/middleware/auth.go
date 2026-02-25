package middleware

import (
	"autocorrect-backend/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

// RequireAuth middleware strictly verifies the JWT token in the Authorization header.
// If valid, it adds the UserID to context.
// If invalid or missing, it returns a 401 Unauthorized error.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONError(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.JSONError(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.LogError("AuthMiddleware.TokenInvalid", err)
			utils.JSONError(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add UserID to context for handlers to use
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

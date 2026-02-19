package api

import (
	"autocorrect-backend/handlers"
	"autocorrect-backend/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// NewRouter creates a new HTTP router with all API routes
func NewRouter() http.Handler {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// Rate Limits Configuration
	// High Volume (GET public data): 10 req/sec, burst 20
	highRate := middleware.CustomRateLimit(10, 20)
	// Medium Volume (GET detailed data): 5 req/sec, burst 10
	mediumRate := middleware.CustomRateLimit(5, 10)
	// Low Volume (POST/PUT): 1 req/sec, burst 3
	lowRate := middleware.CustomRateLimit(1, 3)
	// Critical (Auth/Session): 0.5 req/sec (1 per 2s), burst 2
	criticalRate := middleware.CustomRateLimit(0.5, 2)

	// Auth Middleware wrapper
	auth := middleware.RequireAuth

	// --- Public Endpoints (Rate Limited) ---

	// Daily challenge
	api.Handle("/challenge/today", highRate(http.HandlerFunc(handlers.GetTodaysChallengeHandler))).Methods("GET")
	api.Handle("/challenge/stats", mediumRate(http.HandlerFunc(handlers.GetChallengeStatsHandler))).Methods("GET")
	api.Handle("/challenge/{id}", mediumRate(http.HandlerFunc(handlers.GetChallengeByIDHandler))).Methods("GET")
	api.Handle("/makes", highRate(http.HandlerFunc(handlers.GetMakesHandler))).Methods("GET")
	api.Handle("/models", highRate(http.HandlerFunc(handlers.GetModelsByMakeHandler))).Methods("GET")
	api.Handle("/leaderboard", mediumRate(http.HandlerFunc(handlers.GetLeaderboardHandler))).Methods("GET")
	
	// Bonus Status (GET)
	api.Handle("/challenge/bonus/status", mediumRate(http.HandlerFunc(handlers.GetBonusStatusHandler))).Methods("GET")

	// --- Protected Endpoints (Auth + Rate Limited) ---

	// Submissions
	// Note: We authenticate these to prevent abuse
	api.Handle("/challenge/submit", lowRate(auth(http.HandlerFunc(handlers.SubmitChallengeHandler)))).Methods("POST")
	api.Handle("/challenge/bonus/submit", lowRate(auth(http.HandlerFunc(handlers.SubmitBonusHandler)))).Methods("POST")
	
	// --- Auth Endpoints (Strict Rate Limit) ---

	// Get Session (JWT) - This is the handshake
	api.Handle("/auth/session", criticalRate(http.HandlerFunc(handlers.GetSessionHandler))).Methods("GET")
	
	// Google Login
	api.Handle("/auth/google", criticalRate(http.HandlerFunc(handlers.GoogleLoginHandler))).Methods("POST")

	// Static files (images)
	// Apply global rate limit to images? Maybe not needed for static, but good practice if serving directly
	// Let's rely on Nginx/CDN in prod, but here standard handler
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust this for production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-User-ID", "X-Timezone"},
		ExposedHeaders:   []string{"X-User-ID"},
		AllowCredentials: true,
	})

	// Wrap the router with logging middleware
	handler := LoggingMiddleware(router)

	return c.Handler(handler)
}
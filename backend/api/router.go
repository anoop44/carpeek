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

	// --- Public Handlers (Handshake) ---
	api.Handle("/auth/session", criticalRate(http.HandlerFunc(handlers.GetSessionHandler))).Methods("GET")
	api.Handle("/auth/google", criticalRate(http.HandlerFunc(handlers.GoogleLoginHandler))).Methods("POST")

	// --- Protected Endpoints (Auth + Rate Limited) ---
	
	// Challenge & Game Data
	api.Handle("/challenge/today", highRate(auth(http.HandlerFunc(handlers.GetTodaysChallengeHandler)))).Methods("GET")
	api.Handle("/challenge/stats", mediumRate(auth(http.HandlerFunc(handlers.GetChallengeStatsHandler)))).Methods("GET")
	api.Handle("/challenge/{id}", mediumRate(auth(http.HandlerFunc(handlers.GetChallengeByIDHandler)))).Methods("GET")
	api.Handle("/makes", highRate(auth(http.HandlerFunc(handlers.GetMakesHandler)))).Methods("GET")
	api.Handle("/models", highRate(auth(http.HandlerFunc(handlers.GetModelsByMakeHandler)))).Methods("GET")
	api.Handle("/leaderboard", mediumRate(auth(http.HandlerFunc(handlers.GetLeaderboardHandler)))).Methods("GET")
	api.Handle("/challenge/bonus/status", mediumRate(auth(http.HandlerFunc(handlers.GetBonusStatusHandler)))).Methods("GET")
	api.Handle("/challenge/image/{id}", highRate(http.HandlerFunc(handlers.GetChallengeImageHandler))).Methods("GET")

	// Submissions
	api.Handle("/challenge/submit", lowRate(auth(http.HandlerFunc(handlers.SubmitChallengeHandler)))).Methods("POST")
	api.Handle("/challenge/bonus/submit", lowRate(auth(http.HandlerFunc(handlers.SubmitBonusHandler)))).Methods("POST")

	// Subscription
	api.Handle("/subscription/status", mediumRate(auth(http.HandlerFunc(handlers.GetSubscriptionStatusHandler)))).Methods("GET")
	api.Handle("/subscription/webhook", http.HandlerFunc(handlers.RevenueCatWebhookHandler)).Methods("POST")

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
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-User-ID", "X-Timezone", "X-Browser-Signature"},
		ExposedHeaders:   []string{"X-User-ID", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the router with logging middleware
	handler := LoggingMiddleware(router)

	return c.Handler(handler)
}
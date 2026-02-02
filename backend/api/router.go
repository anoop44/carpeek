package api

import (
	"carpeek-backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// NewRouter creates a new HTTP router with all API routes
func NewRouter() http.Handler {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// Daily challenge endpoints
	api.HandleFunc("/challenge/today", handlers.GetTodaysChallengeHandler).Methods("GET")
	api.HandleFunc("/challenge/{id}", handlers.GetChallengeByIDHandler).Methods("GET")
	api.HandleFunc("/makes", handlers.GetMakesHandler).Methods("GET")
	api.HandleFunc("/models", handlers.GetModelsByMakeHandler).Methods("GET")
	
	// Submission endpoints
	api.HandleFunc("/challenge/submit", handlers.SubmitChallengeHandler).Methods("POST")

	// Static files (images)
	// In production, this will be cached by Firebase Hosting/CDN
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

	return c.Handler(router)
}
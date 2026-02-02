package handlers

import (
	"carpeek-backend/database"
	"carpeek-backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

// Cache for challenges
var (
	challengeCache = make(map[string]*models.DetailedChallengeResponse)
	cacheDates     []string // To maintain order for eviction
	cacheMutex     sync.RWMutex
)

const maxCacheSize = 3

// getLocalizedDate returns the date string for the given timezone
func getLocalizedDate(timezone string) string {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("Invalid timezone %s, defaulting to UTC: %v", timezone, err)
		loc = time.UTC
	}
	return time.Now().In(loc).Format("2006-01-02")
}

// getChallengeFromCacheOrDB fetches the challenge from cache or DB
func getChallengeFromCacheOrDB(sqlxDB *sqlx.DB, date string) (*models.DetailedChallengeResponse, error) {
	cacheMutex.RLock()
	if challenge, ok := challengeCache[date]; ok {
		cacheMutex.RUnlock()
		log.Printf("Cache hit for date: %s", date)
		return challenge, nil
	}
	cacheMutex.RUnlock()

	log.Printf("Cache miss for date: %s, fetching from DB", date)
	challenge, err := database.GetDetailedTodaysChallenge(sqlxDB, date)
	if err != nil {
		return nil, err
	}

	// Add to cache
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Check again in case another goroutine added it
	if _, ok := challengeCache[date]; ok {
		return challengeCache[date], nil
	}

	// Evict if necessary (FIFO)
	if len(cacheDates) >= maxCacheSize {
		oldestDate := cacheDates[0]
		delete(challengeCache, oldestDate)
		cacheDates = cacheDates[1:]
		log.Printf("Evicted %s from cache", oldestDate)
	}

	challengeCache[date] = challenge
	cacheDates = append(cacheDates, date)
	log.Printf("Added %s to cache", date)

	return challenge, nil
}

// GetTodaysChallengeHandler handles the request for today's challenge
func GetTodaysChallengeHandler(w http.ResponseWriter, r *http.Request) {
	timezone := r.Header.Get("X-Timezone")
	if timezone == "" {
		timezone = "UTC"
	}
	date := getLocalizedDate(timezone)

	log.Printf("GET /api/v1/challenge/today: starting fetch for date %s (TZ: %s)", date, timezone)
	db := database.DB
	if db == nil {
		log.Printf("GET /api/v1/challenge/today: database connection not available")
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get challenge from cache or DB
	cachedChallenge, err := getChallengeFromCacheOrDB(sqlxDB, date)
	if err != nil {
		log.Printf("GET /api/v1/challenge/today: challenge not found for date %s: %v", date, err)
		// If no challenge exists for date, return a default response
		defaultChallenge := models.ChallengeResponse{
			ID:       0,
			ImageURL: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(defaultChallenge)
		return
	}

	// Format a minimal response
	response := models.ChallengeResponse{
		ID:       cachedChallenge.ID,
		ImageURL: GetFullImageURL(cachedChallenge.ImageURL),
	}

	log.Printf("GET /api/v1/challenge/today: returning minimal response for challenge ID %d", response.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("GET /api/v1/challenge/today: failed to encode response: %v", err)
	}
}



// GetMakesHandler handles the request for all car makes
func GetMakesHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	sqlxDB := sqlx.NewDb(db, "postgres")
	makes, err := database.GetAllMakes(sqlxDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(makes)
}

// GetModelsByMakeHandler handles the request for models of a specific make
func GetModelsByMakeHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	makeIDStr := r.URL.Query().Get("make_id")
	if makeIDStr == "" {
		http.Error(w, "make_id query parameter is required", http.StatusBadRequest)
		return
	}

	var makeID int
	_, err := fmt.Sscanf(makeIDStr, "%d", &makeID)
	if err != nil {
		http.Error(w, "invalid make_id", http.StatusBadRequest)
		return
	}

	sqlxDB := sqlx.NewDb(db, "postgres")
	models, err := database.GetModelsByMakeID(sqlxDB, makeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models)
}
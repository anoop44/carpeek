package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
		utils.LogEvent("CACHE", "Invalid timezone, defaulting to UTC", timezone)
		loc = time.UTC
	}
	return time.Now().In(loc).Format("2006-01-02")
}

// getChallengeFromCacheOrDB fetches the challenge from cache or DB
func getChallengeFromCacheOrDB(sqlxDB *sqlx.DB, date string) (*models.DetailedChallengeResponse, error) {
	cacheMutex.RLock()
	if challenge, ok := challengeCache[date]; ok {
		cacheMutex.RUnlock()
		utils.LogEvent("CACHE", "Hit for date", date)
		return challenge, nil
	}
	cacheMutex.RUnlock()

	utils.LogEvent("CACHE", "Miss for date, fetching from DB", date)
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
		utils.LogEvent("CACHE", "Evicted from cache", oldestDate)
	}

	challengeCache[date] = challenge
	cacheDates = append(cacheDates, date)
	utils.LogEvent("CACHE", "Added to cache", date)

	return challenge, nil
}

// GetTodaysChallengeHandler handles the request for today's challenge
func GetTodaysChallengeHandler(w http.ResponseWriter, r *http.Request) {
	timezone := r.Header.Get("X-Timezone")
	if timezone == "" {
		timezone = "UTC"
	}
	date := getLocalizedDate(timezone)

	db := database.DB
	if db == nil {
		utils.LogError("GetTodaysChallengeHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get challenge from cache or DB
	cachedChallenge, err := getChallengeFromCacheOrDB(sqlxDB, date)
	if err != nil {
		utils.LogError("GetTodaysChallenge", err)
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

	// Get user identity
	user, err := GetUserFromRequest(w, r, sqlxDB)
	var streakStats *models.UserActivityStats
	var userStatus *models.UserChallengeStatus

	if err == nil && user != nil {
		// Update participation bitmap
		now := time.Now().In(time.UTC)
		if timezone != "UTC" {
			loc, err := time.LoadLocation(timezone)
			if err == nil {
				now = time.Now().In(loc)
			}
		}

		database.UpdateUserActivity(sqlxDB, user.ID, now, "participation")

		// Get streak stats
		streakStats, _ = database.GetUserActivityStats(sqlxDB, user.ID, now)

		// Get user challenge status
		status, err := database.GetUserChallengeStatus(sqlxDB, user.ID, cachedChallenge.ID)
		if err == nil {
			userStatus = status
		}
	}

	// Format a detailed response
	response := models.DetailedChallengeResponse{
		ID:                   cachedChallenge.ID,
		Date:                 cachedChallenge.Date,
		ImageURL:             GetFullImageURL(cachedChallenge.ImageURL),
		NextChallengeSeconds: GetSecondsUntilNextChallenge(timezone),
		StreakStats:          streakStats,
		UserStatus:           userStatus,
	}

	// If the user has completed the challenge, expose the solution
	if userStatus != nil && userStatus.IsCompleted {
		response.Make = cachedChallenge.Make
		response.Model = cachedChallenge.Model
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.LogError("EncodeChallengeResponse", err)
	}
}



// GetMakesHandler handles the request for all car makes
func GetMakesHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("GetMakesHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	sqlxDB := sqlx.NewDb(db, "postgres")
	makes, err := database.GetAllMakes(sqlxDB)
	if err != nil {
		utils.LogError("GetAllMakes", err)
		jsonError(w, err.Error(), http.StatusInternalServerError)
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
		utils.LogError("GetModelsByMakeHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	makeIDStr := r.URL.Query().Get("make_id")
	if makeIDStr == "" {
		jsonError(w, "make_id query parameter is required", http.StatusBadRequest)
		return
	}

	var makeID int
	_, err := fmt.Sscanf(makeIDStr, "%d", &makeID)
	if err != nil {
		jsonError(w, "invalid make_id", http.StatusBadRequest)
		return
	}

	sqlxDB := sqlx.NewDb(db, "postgres")
	models, err := database.GetModelsByMakeID(sqlxDB, makeID)
	if err != nil {
		utils.LogError("GetModelsByMakeID", err)
		jsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models)
}

// GetChallengeStatsHandler returns global stats for a specific challenge
func GetChallengeStatsHandler(w http.ResponseWriter, r *http.Request) {
	challengeIDStr := r.URL.Query().Get("challenge_id")
	if challengeIDStr == "" {
		jsonError(w, "challenge_id is required", http.StatusBadRequest)
		return
	}

	challengeID, err := strconv.Atoi(challengeIDStr)
	if err != nil {
		jsonError(w, "invalid challenge_id", http.StatusBadRequest)
		return
	}

	db := database.DB
	if db == nil {
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	stats, err := database.GetChallengeStats(sqlxDB, challengeID)
	if err != nil {
		utils.LogError("GetChallengeStats", err)
		jsonError(w, "Failed to get challenge stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
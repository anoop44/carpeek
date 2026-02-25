package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
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

		if err := database.UpdateUserActivity(sqlxDB, user.ID, now, "participation"); err != nil {
			utils.LogError("UpdateParticipationActivity", err)
		}

		// Get streak stats
		streakStats, _ = database.GetUserActivityStats(sqlxDB, user.ID, now)

		// Get user challenge status
		status, err := database.GetUserChallengeStatus(sqlxDB, user.ID, cachedChallenge.ID)
		if err == nil {
			userStatus = status
		}
	}

	var pointsEarned float64
	var bonusRound *models.BonusRoundInfo
	if userStatus != nil && userStatus.IsCompleted && user != nil {
		score, _ := database.GetUserChallengeScore(sqlxDB, user.ID, cachedChallenge.ID)
		if score != nil {
			pointsEarned = score.TotalPoints
		}
		bonusInfo, _ := database.GetBonusRoundInfo(sqlxDB, cachedChallenge.ID, user.ID)
		bonusRound = bonusInfo
	}

	// Format a detailed response
	response := models.DetailedChallengeResponse{
		ID:                   cachedChallenge.ID,
		Date:                 cachedChallenge.Date,
		ImageURL:             GetFullImageURL(cachedChallenge.ImageURL, cachedChallenge.ID),
		NextChallengeSeconds: GetSecondsUntilNextChallenge(timezone),
		StreakStats:          streakStats,
		UserStatus:           userStatus,
		PointsEarned:         pointsEarned,
		BonusRound:           bonusRound,
	}

	// If the user has completed the challenge, expose the solution
	if userStatus != nil && userStatus.IsCompleted && cachedChallenge.Model != nil {
		response.Make = cachedChallenge.Make
		// Copy the model to avoid modifying the cached version
		modelCopy := *cachedChallenge.Model
		if modelCopy.ImageURL != nil {
			fullURL := GetModelImageURL(*modelCopy.ImageURL)
			modelCopy.ImageURL = &fullURL
		}
		response.Model = &modelCopy
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

// GetChallengeImageHandler serves the actual image file for a challenge
// but masks the filename by using the challenge ID in the URL.
func GetChallengeImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}

	db := database.DB
	if db == nil {
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get challenge details (we need the actual image path)
	var challenge struct {
		ImageURL string    `db:"image_url"`
		Date     time.Time `db:"date"`
	}
	err = sqlxDB.Get(&challenge, "SELECT image_url, date FROM challenges WHERE id = $1", id)
	if err != nil {
		utils.LogError("GetChallengeImage.Fetch", err)
		jsonError(w, "Challenge not found", http.StatusNotFound)
		return
	}

	// // Security Check: Don't serve images for future challenges
	// timezone := r.Header.Get("X-Timezone")
	// if timezone == "" {
	// 	timezone = "UTC"
	// }
	// loc, _ := time.LoadLocation(timezone)
	// if loc == nil {
	// 	loc = time.UTC
	// }
	// today := time.Now().In(loc).Truncate(24 * time.Hour)
	// challengeDate := challenge.Date.In(loc).Truncate(24 * time.Hour)

	// if challengeDate.After(today) {
	// 	jsonError(w, "Access denied", http.StatusForbidden)
	// 	return
	// }

	// Clean path and ensure it's relative to images dir
	imagePath := challenge.ImageURL
	imagePath = strings.TrimPrefix(imagePath, "/")

	// Assuming images are in ./images directory
	fullPath := imagePath
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		utils.LogError("GetChallengeImage.Stat", fmt.Errorf("file not found: %s", fullPath))
		jsonError(w, "Image file not found", http.StatusNotFound)
		return
	}

	// Serve the file
	http.ServeFile(w, r, fullPath)
}

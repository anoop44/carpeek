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

// GetLeaderboardHandler handles the leaderboard request
func GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("GetLeaderboardHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Determine leaderboard type
	lbType := r.URL.Query().Get("type")
	if lbType == "" {
		lbType = "daily"
	}

	var leaderboard []models.LeaderboardEntry
	var err error

	if lbType == "daily" {
		// 1. Determine today's challenge ID
		timezone := r.Header.Get("X-Timezone")
		if timezone == "" {
			timezone = "UTC"
		}
		date := getLocalizedDate(timezone)
		
		challenge, err := database.GetTodaysChallenge(sqlxDB, date)
		if err != nil {
			// If no challenge for today yet, return empty list
			utils.LogError("GetTodaysChallengeLeaderboard", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]models.LeaderboardEntry{})
			return
		}

		// 2. Get leaderboard for that challenge
		leaderboard, err = database.GetDailyLeaderboard(sqlxDB, challenge.ID)
	} else if lbType == "alltime" {
		leaderboard, err = database.GetAllTimeLeaderboard(sqlxDB)
	} else {
		jsonError(w, "Invalid leaderboard type", http.StatusBadRequest)
		return
	}

	if err != nil {
		utils.LogError("FetchLeaderboard", err)
		jsonError(w, "Failed to fetch leaderboard", http.StatusInternalServerError)
		return
	}

	utils.LogEvent("LEADERBOARD", "Leaderboard fetched", map[string]interface{}{
		"type": lbType,
	})

	// 3. Mark the current user (if authenticated/cookies present)
	user, err := GetUserFromRequest(w, r, sqlxDB)
	if err == nil && user != nil {
		for i := range leaderboard {
			if leaderboard[i].UserID == user.AnonymousID {
				leaderboard[i].IsCurrentUser = true
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leaderboard)
}

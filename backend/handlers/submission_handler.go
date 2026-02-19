package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// SubmitChallengeHandler handles the submission of a challenge guess
func SubmitChallengeHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("SubmitChallengeHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Parse the request body
	var req models.SubmissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.LogError("DecodeSubmissionRequest", err)
		jsonError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log the incoming request for debugging
	utils.LogDebug("SUBMISSION", "Request body", req)

	// Validate required fields
	if req.MakeID <= 0 || req.ModelID <= 0 {
		jsonError(w, "Both make_id and model_id are required", http.StatusBadRequest)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get user identity
	user, err := GetUserFromRequest(w, r, sqlxDB)
	if err != nil {
		utils.LogError("GetUserFromRequestInSubmission", err)
		jsonError(w, "Failed to identify user", http.StatusInternalServerError)
		return
	}

	// Get the challenge
	challenge, err := database.GetChallengeByID(sqlxDB, req.ChallengeID)
	if err != nil {
		utils.LogError("GetChallengeByID", err)
		jsonError(w, "Challenge not found", http.StatusNotFound)
		return
	}

	// Check current status
	status, err := database.GetUserChallengeStatus(sqlxDB, user.ID, challenge.ID)
	if err != nil {
		utils.LogError("GetUserChallengeStatusSubmission", err)
		jsonError(w, "Failed to get user status", http.StatusInternalServerError)
		return
	}

	if status.IsCompleted {
		jsonError(w, "Challenge already completed or max attempts reached", http.StatusForbidden)
		return
	}

	// Validate the submission
	result, err := database.ValidateSubmission(sqlxDB, challenge.ID, req.MakeID, req.ModelID)
	if err != nil {
		utils.LogError("ValidateSubmission", err)
		jsonError(w, "Failed to validate submission", http.StatusInternalServerError)
		return
	}

	points, err := database.RecordSubmission(sqlxDB, user.ID, challenge.ID, req.MakeID, req.ModelID, result.IsMakeCorrect, result.IsModelCorrect, result.Correct)
	if err != nil {
		utils.LogError("RecordSubmission", err)
		// Continue even if recording fails, but we've logged it
	}
	result.PointsEarned = points

	// Update user status for the response
	updatedStatus, err := database.GetUserChallengeStatus(sqlxDB, user.ID, challenge.ID)
	if err == nil {
		result.UserStatus = updatedStatus
	}

	// If the challenge was solved correctly, include bonus round info
	if result.Correct {
		bonusInfo, err := database.GetBonusRoundInfo(sqlxDB, challenge.ID, user.ID)
		if err == nil {
			// Only include if at least one bonus type is enabled
			if bonusInfo.YearRangeEnabled || bonusInfo.GenerationEnabled || bonusInfo.CodenameEnabled {
				result.BonusRound = bonusInfo
			}
		}
	}

	// Calculate seconds until next challenge
	timezone := r.Header.Get("X-Timezone")
	if timezone == "" {
		timezone = "UTC"
	}
	// Record activity
	now := time.Now().In(time.UTC)
	if timezone != "UTC" {
		loc, err := time.LoadLocation(timezone)
		if err == nil {
			now = time.Now().In(loc)
		}
	}
	database.UpdateUserActivity(sqlxDB, user.ID, now, "submission")
	result.StreakStats, _ = database.GetUserActivityStats(sqlxDB, user.ID, now)

	result.NextChallengeSeconds = GetSecondsUntilNextChallenge(timezone)

	// Log the submission event
	utils.LogEvent("SUBMISSION", "User submitted guess", map[string]interface{}{
		"userID":      user.ID,
		"challengeID": challenge.ID,
		"correct":      result.Correct,
		"points":       result.PointsEarned,
		"nextIn":       result.NextChallengeSeconds,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetChallengeByIDHandler handles the request for a specific challenge by ID
func GetChallengeByIDHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Get the challenge ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get the challenge by ID
	challenge, err := database.GetChallengeByID(sqlxDB, id)
	if err != nil {
		jsonError(w, "Challenge not found", http.StatusNotFound)
		return
	}

	// Format the response
	response := models.ChallengeResponse{
		ID:       challenge.ID,
		ImageURL: GetFullImageURL(challenge.ImageURL),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
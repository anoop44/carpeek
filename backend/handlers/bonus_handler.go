package handlers

import (
	"autocorrect-backend/database"
	"autocorrect-backend/models"
	"autocorrect-backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// SubmitBonusHandler handles bonus round submissions
// Each bonus type (year_range, generation, codename) can only be attempted once
// Correct answers give 1 point each
func SubmitBonusHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("SubmitBonusHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Parse the request body
	var req models.BonusSubmissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.LogError("DecodeBonusSubmissionRequest", err)
		jsonError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log incoming request
	utils.LogDebug("BONUS", "Request body", req)

	// Validate required fields
	if req.ChallengeID <= 0 || req.BonusType == "" || req.Value == "" {
		jsonError(w, "challenge_id, bonus_type, and value are required", http.StatusBadRequest)
		return
	}

	// Validate bonus_type
	if req.BonusType != "year_range" && req.BonusType != "generation" && req.BonusType != "codename" {
		jsonError(w, "bonus_type must be 'year_range', 'generation', or 'codename'", http.StatusBadRequest)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get user identity
	user, err := GetUserFromRequest(w, r, sqlxDB)
	if err != nil {
		utils.LogError("GetUserFromRequestInBonus", err)
		jsonError(w, "Failed to identify user", http.StatusInternalServerError)
		return
	}

	// Check if user has already completed the main challenge
	status, err := database.GetUserChallengeStatus(sqlxDB, user.ID, req.ChallengeID)
	if err != nil {
		utils.LogError("GetUserChallengeStatusBonus", err)
		jsonError(w, "Failed to get user status", http.StatusInternalServerError)
		return
	}

	// Bonus rounds are only available after completing the main challenge correctly
	if !status.IsCorrect {
		jsonError(w, "Bonus rounds are only available after correctly solving the main challenge", http.StatusForbidden)
		return
	}

	// Check if this bonus type is eligible and not already attempted
	eligible, err := database.CheckBonusEligibility(sqlxDB, req.ChallengeID, user.ID, req.BonusType)
	if err != nil {
		utils.LogError("CheckBonusEligibility", err)
		jsonError(w, "Failed to check bonus eligibility", http.StatusInternalServerError)
		return
	}

	if !eligible {
		jsonError(w, "This bonus round is not available or has already been attempted", http.StatusForbidden)
		return
	}

	// Record the bonus submission
	result, err := database.RecordBonusSubmission(sqlxDB, user.ID, req.ChallengeID, req.BonusType, req.Value)
	if err != nil {
		utils.LogError("RecordBonusSubmission", err)
		jsonError(w, "Failed to record bonus submission", http.StatusInternalServerError)
		return
	}

	// Log the bonus submission event
	utils.LogEvent("BONUS", "User submitted bonus", map[string]interface{}{
		"userID":      user.ID,
		"challengeID": req.ChallengeID,
		"bonusType":   req.BonusType,
		"correct":     result.Correct,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetBonusStatusHandler returns the bonus round status for a challenge
func GetBonusStatusHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		utils.LogError("GetBonusStatusHandler", fmt.Errorf("database connection not available"))
		jsonError(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Get challenge_id from query params
	challengeIDStr := r.URL.Query().Get("challenge_id")
	if challengeIDStr == "" {
		jsonError(w, "challenge_id is required", http.StatusBadRequest)
		return
	}

	challengeID, err := strconv.Atoi(challengeIDStr)
	if err != nil {
		jsonError(w, "Invalid challenge_id", http.StatusBadRequest)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get user identity
	user, err := GetUserFromRequest(w, r, sqlxDB)
	if err != nil {
		utils.LogError("GetUserFromRequestInBonusStatus", err)
		jsonError(w, "Failed to identify user", http.StatusInternalServerError)
		return
	}

	// Get bonus round info
	bonusInfo, err := database.GetBonusRoundInfo(sqlxDB, challengeID, user.ID)
	if err != nil {
		utils.LogError("GetBonusRoundInfo", err)
		jsonError(w, "Failed to get bonus round info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bonusInfo)
}

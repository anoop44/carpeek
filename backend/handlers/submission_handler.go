package handlers

import (
	"carpeek-backend/database"
	"carpeek-backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// SubmitChallengeHandler handles the submission of a challenge guess
func SubmitChallengeHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Parse the request body
	var req models.SubmissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.MakeID <= 0 || req.ModelID <= 0 {
		http.Error(w, "Both make_id and model_id are required", http.StatusBadRequest)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get user identity
	user, err := GetUserFromRequest(w, r, sqlxDB)
	if err != nil {
		http.Error(w, "Failed to identify user", http.StatusInternalServerError)
		return
	}

	// Get the challenge
	challenge, err := database.GetChallengeByID(sqlxDB, req.ChallengeID)
	if err != nil {
		http.Error(w, "Challenge not found", http.StatusNotFound)
		return
	}

	// Check current status
	status, err := database.GetUserChallengeStatus(sqlxDB, user.ID, challenge.ID)
	if err != nil {
		http.Error(w, "Failed to get user status", http.StatusInternalServerError)
		return
	}

	if status.IsCompleted {
		http.Error(w, "Challenge already completed or max attempts reached", http.StatusForbidden)
		return
	}

	// Validate the submission
	result, err := database.ValidateSubmission(sqlxDB, challenge.ID, req.MakeID, req.ModelID)
	if err != nil {
		http.Error(w, "Failed to validate submission", http.StatusInternalServerError)
		return
	}

	// Record the submission
	err = database.RecordSubmission(sqlxDB, user.ID, challenge.ID, req.MakeID, req.ModelID, result.Correct)
	if err != nil {
		// Log error but continue
	}

	// Update user status for the response
	updatedStatus, err := database.GetUserChallengeStatus(sqlxDB, user.ID, challenge.ID)
	if err == nil {
		result.UserStatus = updatedStatus
	}

	// Ensure the image URL is a full URL
	result.ImageURL = GetFullImageURL(result.ImageURL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetChallengeByIDHandler handles the request for a specific challenge by ID
func GetChallengeByIDHandler(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	if db == nil {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	// Get the challenge ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid challenge ID", http.StatusBadRequest)
		return
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Get the challenge by ID
	challenge, err := database.GetChallengeByID(sqlxDB, id)
	if err != nil {
		http.Error(w, "Challenge not found", http.StatusNotFound)
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
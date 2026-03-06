package database

import (
	"autocorrect-backend/models"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// GetTodaysChallenge returns the challenge for a specific date
func GetTodaysChallenge(db *sqlx.DB, date string) (*models.Challenge, error) {
	query := `
		SELECT id, date, image_url, solution_make_id, solution_model_id, created_at, updated_at
		FROM challenges
		WHERE date = $1
	`

	// No longer hardcoded today
	today := date

	var result struct {
		ID              int           `db:"id"`
		Date            time.Time     `db:"date"`
		ImageURL        string        `db:"image_url"`
		SolutionMakeID  sql.NullInt64 `db:"solution_make_id"`
		SolutionModelID sql.NullInt64 `db:"solution_model_id"`
		CreatedAt       time.Time     `db:"created_at"`
		UpdatedAt       time.Time     `db:"updated_at"`
	}

	err := db.Get(&result, query, today)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's challenge: %v", err)
	}

	challenge := &models.Challenge{
		ID:              result.ID,
		Date:            result.Date,
		ImageURL:        result.ImageURL,
		SolutionMakeID:  result.SolutionMakeID,
		SolutionModelID: result.SolutionModelID,
		CreatedAt:       result.CreatedAt,
		UpdatedAt:       result.UpdatedAt,
	}

	return challenge, nil
}

// GetDetailedTodaysChallenge returns the detailed challenge for a specific date
func GetDetailedTodaysChallenge(db *sqlx.DB, date string) (*models.DetailedChallengeResponse, error) {
	query := `
		SELECT
			c.id,
			c.date,
			c.image_url,
			m.id as make_id,
			m.name as make_name,
			mo.id as model_id,
			mo.name as model_name,
			mo.year_range as model_year_range,
			mo.generation as model_generation,
			mo.location as model_location,
			mo.codename as model_codename,
			mo.image_url as model_image_url,
			mo.known_for as model_known_for
		FROM challenges c
		LEFT JOIN makes m ON c.solution_make_id = m.id
		LEFT JOIN models mo ON c.solution_model_id = mo.id
		WHERE c.date = $1
	`

	// No longer hardcoded today
	today := date
	var result struct {
		ID              int       `db:"id"`
		Date            time.Time `db:"date"`
		ImageURL        string    `db:"image_url"`
		MakeID          *int      `db:"make_id"`          // Use pointer to handle possible NULL
		MakeName        *string   `db:"make_name"`        // Use pointer to handle possible NULL
		ModelID         *int      `db:"model_id"`         // Use pointer to handle possible NULL
		ModelName       *string   `db:"model_name"`       // Use pointer to handle possible NULL
		ModelYearRange  *string   `db:"model_year_range"` // Use pointer to handle possible NULL
		ModelGeneration *string   `db:"model_generation"` // Use pointer to handle possible NULL
		ModelLocation   *string   `db:"model_location"`   // Use pointer to handle possible NULL
		ModelCodename   *string   `db:"model_codename"`   // Use pointer to handle possible NULL
		ModelImageURL   *string   `db:"model_image_url"`  // Use pointer to handle possible NULL
		ModelKnownFor   *string   `db:"model_known_for"`  // Use pointer to handle possible NULL
	}

	err := db.Get(&result, query, today)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's challenge: %v", err)
	}

	detailedChallenge := &models.DetailedChallengeResponse{
		ID:       result.ID,
		Date:     result.Date.Format("2006-01-02"),
		ImageURL: result.ImageURL,
	}

	// Add make info if available
	if result.MakeID != nil && result.MakeName != nil {
		detailedChallenge.Make = &models.Make{
			ID:   *result.MakeID,
			Name: *result.MakeName,
		}
	}

	// Add model info if available
	if result.ModelID != nil && result.ModelName != nil {
		model := &models.Model{
			ID:         *result.ModelID,
			Name:       *result.ModelName,
			YearRange:  result.ModelYearRange,
			Generation: result.ModelGeneration,
			Location:   result.ModelLocation,
			Codename:   result.ModelCodename,
			ImageURL:   result.ModelImageURL,
			KnownFor:   result.ModelKnownFor,
		}

		// Set make_id from the challenge record
		if result.MakeID != nil {
			model.MakeID = *result.MakeID
		}

		detailedChallenge.Model = model
	}

	return detailedChallenge, nil
}

// ValidateSubmission validates a submission against a challenge
func ValidateSubmission(db *sqlx.DB, challengeID int, submittedMakeID int, submittedModelIDs []int) (*models.SubmissionResult, int, error) {
	query := `
		SELECT
			c.id,
			c.date,
			c.image_url,
			c.solution_make_id,
			c.solution_model_id,
			m.name as make_name,
			mo.name as model_name,
			mo.year_range as model_year_range,
			mo.generation as model_generation,
			mo.location as model_location,
			mo.codename as model_codename,
			mo.known_for as model_known_for,
			mo.image_url as model_image_url
		FROM challenges c
		LEFT JOIN makes m ON c.solution_make_id = m.id
		LEFT JOIN models mo ON c.solution_model_id = mo.id
		WHERE c.id = $1
	`

	var result struct {
		ID                      int       `db:"id"`
		Date                    time.Time `db:"date"`
		ImageURL                string    `db:"image_url"`
		SolutionMakeID          *int      `db:"solution_make_id"`
		SolutionModelID         *int      `db:"solution_model_id"`
		SolutionMakeName        *string   `db:"make_name"`
		SolutionModelName       *string   `db:"model_name"`
		SolutionModelYearRange  *string   `db:"model_year_range"`
		SolutionModelGeneration *string   `db:"model_generation"`
		SolutionModelLocation   *string   `db:"model_location"`
		SolutionModelCodename   *string   `db:"model_codename"`
		SolutionModelKnownFor   *string   `db:"model_known_for"`
		SolutionModelImageURL  *string   `db:"model_image_url"`
	}

	err := db.Get(&result, query, challengeID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get challenge for validation: %v", err)
	}

	submissionResult := &models.SubmissionResult{
		ID:       result.ID,
		Date:     result.Date.Format("2006-01-02"),
		ImageURL: result.ImageURL,
		Solution: &models.SolutionInfo{
			MakeName:  "",
			ModelName: "",
		},
	}
	if result.SolutionMakeName != nil {
		submissionResult.Solution.MakeName = *result.SolutionMakeName
	}
	if result.SolutionModelName != nil {
		submissionResult.Solution.ModelName = *result.SolutionModelName
	}
	submissionResult.Solution.YearRange = result.SolutionModelYearRange
	submissionResult.Solution.Generation = result.SolutionModelGeneration
	submissionResult.Solution.Codename = result.SolutionModelCodename
	submissionResult.Solution.KnownFor = result.SolutionModelKnownFor
	submissionResult.Solution.ImageURL = result.SolutionModelImageURL

	solMakeID := 0
	if result.SolutionMakeID != nil {
		solMakeID = *result.SolutionMakeID
	}

	// Evaluate if the correct solution model ID is among the submitted model IDs
	isModelCorrect := false
	matchedModelID := 0
	if len(submittedModelIDs) > 0 {
		matchedModelID = submittedModelIDs[0] // Default fallback for DB recording
	}
	
	if result.SolutionModelID != nil {
		solModelID := *result.SolutionModelID
		for _, id := range submittedModelIDs {
			if id == solModelID {
				isModelCorrect = true
				matchedModelID = id // Use the exact matching ID for DB recording
				break
			}
		}
	}

	// Even if not perfectly correct, if they guessed the right model name (from any ID in group), 
	// we should treat the model as correct conceptually, but wait, the generation might be wrong.
	// Actually, the user guesses Make + Model Name from the dropdown. The exact Generation is NOT guessed initially!
	// So simply matching the name is what we do!
	// The problem described: "When such models are returned in api, the response should have all the ids of the model in it and when user selects, all the ids associated should be sent back for verification and challenge answer should be compared against every id in submission"

	isMakeCorrect := submittedMakeID == solMakeID

	// Set the correctness flags
	submissionResult.IsMakeCorrect = isMakeCorrect
	submissionResult.IsModelCorrect = isModelCorrect

	// Check if both make and model are correct
	if isMakeCorrect && isModelCorrect {
		submissionResult.Correct = true
		submissionResult.Partial = false
		submissionResult.Message = "Perfect! You correctly identified the car."
	} else if isMakeCorrect || isModelCorrect {
		// Partially correct - either make or model is correct
		submissionResult.Correct = false
		submissionResult.Partial = true
		if isMakeCorrect {
			submissionResult.Message = "Close! You got the make right but the model is incorrect."
		} else {
			submissionResult.Message = "Close! You got the model right but the make is incorrect."
		}
	} else {
		// Both are incorrect
		submissionResult.Correct = false
		submissionResult.Partial = false
		submissionResult.Message = "Incorrect. Neither the make nor the model is correct."
	}

	return submissionResult, matchedModelID, nil
}

// GetMakeByID retrieves a make by its ID
func GetMakeByID(db *sqlx.DB, id int) (*models.Make, error) {
	query := `SELECT id, name, created_at, updated_at FROM makes WHERE id = $1`

	var make models.Make
	err := db.Get(&make, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get make by ID: %v", err)
	}

	return &make, nil
}

// GetChallengeByID returns a challenge by its ID
func GetChallengeByID(db *sqlx.DB, id int) (*models.Challenge, error) {
	query := `
		SELECT id, date, image_url, solution_make_id, solution_model_id, created_at, updated_at
		FROM challenges
		WHERE id = $1
	`

	var challenge models.Challenge
	err := db.Get(&challenge, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge by ID: %v", err)
	}

	return &challenge, nil
}

// GetModelByID retrieves a model by its ID
func GetModelByID(db *sqlx.DB, id int) (*models.Model, error) {
	query := `SELECT id, name, make_id, year_range, generation, location, codename, created_at, updated_at FROM models WHERE id = $1`

	var model models.Model
	err := db.Get(&model, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get model by ID: %v", err)
	}

	return &model, nil
}

// GetAllMakes retrieves all makes from the database
func GetAllMakes(db *sqlx.DB) ([]models.Make, error) {
	query := `SELECT id, name, created_at, updated_at FROM makes ORDER BY name ASC`

	var makes []models.Make
	err := db.Select(&makes, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all makes: %v", err)
	}

	return makes, nil
}

// GetModelsByMakeID retrieves unique models for a specific make, with all associated IDs
func GetModelsByMakeID(db *sqlx.DB, makeID int) ([]models.GroupedModel, error) {
	query := `SELECT id, name FROM models WHERE make_id = $1 ORDER BY name ASC`

	var rawModels []struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}
	err := db.Select(&rawModels, query, makeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get models by make ID: %v", err)
	}

	modelMap := make(map[string][]int)
	var names []string
	for _, m := range rawModels {
		if _, exists := modelMap[m.Name]; !exists {
			names = append(names, m.Name)
		}
		modelMap[m.Name] = append(modelMap[m.Name], m.ID)
	}

	result := make([]models.GroupedModel, 0, len(names))
	for _, name := range names {
		result = append(result, models.GroupedModel{
			Name: name,
			ID:   modelMap[name],
		})
	}

	return result, nil
}

// GetOrCreateUser retrieves a user by anonymous ID or creates a new one
func GetOrCreateUser(db *sqlx.DB, anonymousID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, anonymous_id, google_id, email, display_name, profile_picture_url, is_linked, created_at, updated_at FROM users WHERE anonymous_id = $1`
	err := db.Get(&user, query, anonymousID)
	if err == nil {
		return &user, nil
	}

	// Create new user if not found
	insertQuery := `INSERT INTO users (anonymous_id) VALUES ($1) RETURNING id, anonymous_id, google_id, email, display_name, profile_picture_url, is_linked, created_at, updated_at`
	err = db.QueryRowx(insertQuery, anonymousID).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &user, nil
}

// GetUserByAnonymousID retrieves a user by anonymous ID without creating one
func GetUserByAnonymousID(db *sqlx.DB, anonymousID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, anonymous_id, google_id, email, display_name, profile_picture_url, is_linked, created_at, updated_at FROM users WHERE anonymous_id = $1`
	err := db.Get(&user, query, anonymousID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByGoogleID retrieves a user by their Google ID
func GetUserByGoogleID(db *sqlx.DB, googleID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, anonymous_id, google_id, email, display_name, profile_picture_url, is_linked, created_at, updated_at FROM users WHERE google_id = $1`
	err := db.Get(&user, query, googleID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// LinkGoogleAccount links a Google account to the current anonymous user record.
// If the Google account is already linked to another user, that other user record is removed.
func LinkGoogleAccount(db *sqlx.DB, anonymousID string, googleID string, email string, name string, picture string) (*models.User, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// 1. Get current user record
	var currentUserID int
	err = tx.Get(&currentUserID, "SELECT id FROM users WHERE anonymous_id = $1", anonymousID)
	if err != nil {
		// If current session user record doesn't exist, something is wrong
		return nil, fmt.Errorf("current user record not found: %v", err)
	}

	// 2. Check if this Google account is linked to a DIFFERENT user record
	var otherUserID int
	err = tx.Get(&otherUserID, "SELECT id FROM users WHERE google_id = $1 AND id != $2", googleID, currentUserID)
	if err == nil {
		// It's linked elsewhere. The user said "no need to move submission and other details",
		// so we remove the other user record to prevent duplication on the leaderboard.
		_, err = tx.Exec("DELETE FROM users WHERE id = $1", otherUserID)
		if err != nil {
			return nil, fmt.Errorf("failed to remove existing link: %v", err)
		}
	}

	// 3. Update the current user record with Google info
	updateQuery := `
		UPDATE users 
		SET google_id = $2, 
			email = $3, 
			display_name = $4, 
			profile_picture_url = $5, 
			is_linked = TRUE, 
			updated_at = CURRENT_TIMESTAMP 
		WHERE id = $1 
		RETURNING id, anonymous_id, google_id, email, display_name, profile_picture_url, is_linked, created_at, updated_at
	`
	var user models.User
	err = tx.QueryRowx(updateQuery, currentUserID, googleID, email, name, picture).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to update current user with google details: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &user, nil
}

// RecordSubmission saves a user's guess for a challenge and updates their score
// Points system:
// - 5 points: Both correct on 1st attempt
// - 3 points: Both correct on 2nd attempt
// - 1 point: Both correct on 3rd attempt
// - 0.5 points: Make correct in latest submission (only if consistently correct across all attempts)
func RecordSubmission(db *sqlx.DB, userID int, challengeID int, makeID int, modelID int, isMakeCorrect bool, isModelCorrect bool, isFullyCorrect bool) (float64, error) {
	// Begin transaction
	tx, err := db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Get current attempt count
	var currentAttempts int
	err = tx.Get(&currentAttempts, "SELECT COUNT(*) FROM submissions WHERE user_id = $1 AND challenge_id = $2", userID, challengeID)
	if err != nil {
		return 0, fmt.Errorf("failed to get attempt count: %v", err)
	}
	attemptNumber := currentAttempts + 1

	// Insert the submission with detailed correctness info
	insertQuery := `
		INSERT INTO submissions (user_id, challenge_id, make_id, model_id, is_correct, is_make_correct, is_model_correct, attempt_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = tx.Exec(insertQuery, userID, challengeID, makeID, modelID, isFullyCorrect, isMakeCorrect, isModelCorrect, attemptNumber)
	if err != nil {
		return 0, fmt.Errorf("failed to record submission: %v", err)
	}

	// Now update or insert into user_challenge_scores
	pointsEarned, err := updateUserChallengeScore(tx, userID, challengeID, attemptNumber, isMakeCorrect, isModelCorrect, isFullyCorrect)
	if err != nil {
		return 0, fmt.Errorf("failed to update score: %v", err)
	}

	return pointsEarned, tx.Commit()
}

// updateUserChallengeScore calculates and stores the user's score for a challenge
func updateUserChallengeScore(tx *sqlx.Tx, userID int, challengeID int, attemptNumber int, isMakeCorrect bool, isModelCorrect bool, isFullyCorrect bool) (float64, error) {
	// Check if there's an existing score record
	var existingScore models.UserChallengeScore
	err := tx.Get(&existingScore, "SELECT * FROM user_challenge_scores WHERE user_id = $1 AND challenge_id = $2", userID, challengeID)
	scoreExists := err == nil

	// If already fully solved, don't update (they already got their points)
	if scoreExists && existingScore.IsFullySolved {
		return 0, nil
	}

	// Calculate make_ever_wrong - check if make was ever wrong in any submission
	// This handles the edge case: if make was correct before but wrong now, no make bonus
	makeEverWrong := false
	if scoreExists {
		makeEverWrong = existingScore.MakeEverWrong
	}
	if !isMakeCorrect {
		makeEverWrong = true
	}

	// Calculate points
	var fullSolvePoints float64 = 0
	var makeBonusPoints float64 = 0

	if isFullyCorrect {
		// Full solve points based on attempt number
		switch attemptNumber {
		case 1:
			fullSolvePoints = 5.0
		case 2:
			fullSolvePoints = 3.0
		case 3:
			fullSolvePoints = 1.0
		}
	}

	// Make bonus: 0.5 points if make is correct in latest submission AND was never wrong
	// Edge case: if earlier had correct make but now wrong, no bonus
	if isMakeCorrect && !makeEverWrong && !isFullyCorrect {
		makeBonusPoints = 0.5
	}

	totalPoints := fullSolvePoints + makeBonusPoints

	if scoreExists {
		// Update existing record
		updateQuery := `
			UPDATE user_challenge_scores 
			SET attempt_number = $3,
				full_solve_points = $4,
				make_bonus_points = $5,
				total_points = $6,
				is_fully_solved = $7,
				make_ever_wrong = $8,
				updated_at = CURRENT_TIMESTAMP
			WHERE user_id = $1 AND challenge_id = $2
		`
		_, err = tx.Exec(updateQuery, userID, challengeID, attemptNumber, fullSolvePoints, makeBonusPoints, totalPoints, isFullyCorrect, makeEverWrong)
	} else {
		// Insert new record
		insertQuery := `
			INSERT INTO user_challenge_scores 
			(user_id, challenge_id, attempt_number, full_solve_points, make_bonus_points, total_points, is_fully_solved, make_ever_wrong)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`
		_, err = tx.Exec(insertQuery, userID, challengeID, attemptNumber, fullSolvePoints, makeBonusPoints, totalPoints, isFullyCorrect, makeEverWrong)
	}

	return totalPoints, err
}

// GetUserChallengeStatus returns the current status of a user for a specific challenge
func GetUserChallengeStatus(db *sqlx.DB, userID int, challengeID int) (*models.UserChallengeStatus, error) {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM submissions WHERE user_id = $1 AND challenge_id = $2", userID, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission count: %v", err)
	}

	var isCorrect bool
	err = db.Get(&isCorrect, "SELECT EXISTS(SELECT 1 FROM submissions WHERE user_id = $1 AND challenge_id = $2 AND is_correct = true)", userID, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if completed correctly: %v", err)
	}

	maxAttempts := 3 // 3 attempts: 1st = 5pts, 2nd = 3pts, 3rd = 1pt

	status := &models.UserChallengeStatus{
		Attempts:    count,
		MaxAttempts: maxAttempts,
		IsCompleted: isCorrect || count >= maxAttempts,
		IsCorrect:   isCorrect,
	}

	return status, nil
}

// GetUserChallengeScore returns the user's score details for a specific challenge
func GetUserChallengeScore(db *sqlx.DB, userID int, challengeID int) (*models.UserChallengeScore, error) {
	var score models.UserChallengeScore
	query := `SELECT id, user_id, challenge_id, attempt_number, full_solve_points, make_bonus_points, bonus_round_points, total_points, is_fully_solved, make_ever_wrong, created_at, updated_at FROM user_challenge_scores WHERE user_id = $1 AND challenge_id = $2`
	err := db.Get(&score, query, userID, challengeID)
	if err != nil {
		return nil, err
	}
	return &score, nil
}

// GetBonusRoundInfo determines which bonus rounds are available for a challenge
// A bonus type is enabled if:
// - The model has data for that field (not null)
// - If codename is enabled only if codename values differ from generation
// - If both generation and codename exist and are the same, only generation is enabled
func GetBonusRoundInfo(db *sqlx.DB, challengeID int, userID int) (*models.BonusRoundInfo, error) {
	// Get the challenge's solution model details
	query := `
		SELECT 
			mo.year_range,
			mo.generation,
			mo.codename
		FROM challenges c
		JOIN models mo ON c.solution_model_id = mo.id
		WHERE c.id = $1
	`

	var result struct {
		YearRange  *string `db:"year_range"`
		Generation *string `db:"generation"`
		Codename   *string `db:"codename"`
	}

	err := db.Get(&result, query, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge model info: %v", err)
	}

	info := &models.BonusRoundInfo{}

	// Year range is enabled if it exists
	info.YearRangeEnabled = result.YearRange != nil && *result.YearRange != ""

	// Generation is enabled if it exists
	info.GenerationEnabled = result.Generation != nil && *result.Generation != ""

	// Codename is enabled if:
	// 1. It exists
	// 2. It's different from generation (if both exist)
	if result.Codename != nil && *result.Codename != "" {
		if result.Generation != nil && *result.Generation != "" {
			// Both exist - only enable codename if different from generation
			info.CodenameEnabled = *result.Codename != *result.Generation
		} else {
			// Only codename exists
			info.CodenameEnabled = true
		}
	}

	// Check which bonuses the user has already attempted and their results
	attemptedQuery := `
		SELECT bonus_type, is_correct FROM bonus_submissions
		WHERE user_id = $1 AND challenge_id = $2
	`
	var attempts []struct {
		BonusType string `db:"bonus_type"`
		IsCorrect bool   `db:"is_correct"`
	}
	err = db.Select(&attempts, attemptedQuery, userID, challengeID)
	if err != nil && err != sql.ErrNoRows {
		// Ignore "no rows" error, it just means no attempts yet
	}

	for _, a := range attempts {
		switch a.BonusType {
		case "year_range":
			info.YearRangeAttempted = true
			info.YearRangeCorrect = a.IsCorrect
		case "generation":
			info.GenerationAttempted = true
			info.GenerationCorrect = a.IsCorrect
		case "codename":
			info.CodenameAttempted = true
			info.CodenameCorrect = a.IsCorrect
		}
	}

	// Set point values for display
	info.YearRangePoints = 1.0
	info.GenerationPoints = 1.0
	info.CodenamePoints = 1.0

	return info, nil
}

// CheckBonusEligibility verifies if a user can attempt a specific bonus type
func CheckBonusEligibility(db *sqlx.DB, challengeID int, userID int, bonusType string) (bool, error) {
	// Validate bonus type
	if bonusType != "year_range" && bonusType != "generation" && bonusType != "codename" {
		return false, fmt.Errorf("invalid bonus type: %s", bonusType)
	}

	// Check if already attempted
	var count int
	err := db.Get(&count,
		"SELECT COUNT(*) FROM bonus_submissions WHERE user_id = $1 AND challenge_id = $2 AND bonus_type = $3",
		userID, challengeID, bonusType)
	if err != nil {
		return false, fmt.Errorf("failed to check bonus attempts: %v", err)
	}

	if count > 0 {
		return false, nil // Already attempted
	}

	// Check if this bonus type is eligible for this challenge
	info, err := GetBonusRoundInfo(db, challengeID, userID)
	if err != nil {
		return false, err
	}

	switch bonusType {
	case "year_range":
		return info.YearRangeEnabled, nil
	case "generation":
		return info.GenerationEnabled, nil
	case "codename":
		return info.CodenameEnabled, nil
	}

	return false, nil
}

// RecordBonusSubmission records a bonus round attempt and updates points
func RecordBonusSubmission(db *sqlx.DB, userID int, challengeID int, bonusType string, submittedValue string) (*models.BonusSubmissionResult, error) {
	// Get the correct value
	query := `
		SELECT 
			mo.year_range,
			mo.generation,
			mo.codename
		FROM challenges c
		JOIN models mo ON c.solution_model_id = mo.id
		WHERE c.id = $1
	`

	var result struct {
		YearRange  *string `db:"year_range"`
		Generation *string `db:"generation"`
		Codename   *string `db:"codename"`
	}

	err := db.Get(&result, query, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge model info: %v", err)
	}

	var isCorrect bool
	var correctValue string
	submittedValue = strings.TrimSpace(submittedValue)

	switch bonusType {
	case "year_range":
		if result.YearRange != nil {
			correctValue = *result.YearRange
		}
		isCorrect = checkYearRange(submittedValue, correctValue)

	case "generation":
		// Build correctValue for display: prefer generation, fallback codename
		if result.Generation != nil {
			correctValue = *result.Generation
		}
		// Check submitted value against both generation AND codename
		genMatch := result.Generation != nil && strings.EqualFold(submittedValue, strings.TrimSpace(*result.Generation))
		codeMatch := result.Codename != nil && strings.EqualFold(submittedValue, strings.TrimSpace(*result.Codename))
		isCorrect = genMatch || codeMatch
		// Show both values if they differ
		if result.Generation != nil && result.Codename != nil && !strings.EqualFold(*result.Generation, *result.Codename) {
			correctValue = fmt.Sprintf("%s / %s", *result.Generation, *result.Codename)
		} else if result.Codename != nil && result.Generation == nil {
			correctValue = *result.Codename
		}

	case "codename":
		if result.Codename != nil {
			correctValue = *result.Codename
		}
		// Also accept generation as match for codename bonus
		codeMatch := result.Codename != nil && strings.EqualFold(submittedValue, strings.TrimSpace(*result.Codename))
		genMatch := result.Generation != nil && strings.EqualFold(submittedValue, strings.TrimSpace(*result.Generation))
		isCorrect = codeMatch || genMatch
		if result.Codename != nil && result.Generation != nil && !strings.EqualFold(*result.Codename, *result.Generation) {
			correctValue = fmt.Sprintf("%s / %s", *result.Codename, *result.Generation)
		} else if result.Generation != nil && result.Codename == nil {
			correctValue = *result.Generation
		}

	default:
		return nil, fmt.Errorf("unknown bonus type: %s", bonusType)
	}

	// Begin transaction
	tx, err := db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Insert the bonus submission
	insertQuery := `
		INSERT INTO bonus_submissions (user_id, challenge_id, bonus_type, submitted_value, is_correct)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = tx.Exec(insertQuery, userID, challengeID, bonusType, submittedValue, isCorrect)
	if err != nil {
		return nil, fmt.Errorf("failed to record bonus submission: %v", err)
	}

	var pointsEarned float64 = 0
	if isCorrect {
		pointsEarned = 1.0
		// Update user_challenge_scores
		err = updateBonusPoints(tx, userID, challengeID, 1.0)
		if err != nil {
			return nil, fmt.Errorf("failed to update bonus points: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	response := &models.BonusSubmissionResult{
		Correct:      isCorrect,
		PointsEarned: pointsEarned,
		CorrectValue: correctValue, // Always include for frontend display
	}

	if isCorrect {
		response.Message = fmt.Sprintf("Correct! You earned %.1f bonus point.", pointsEarned)
	} else {
		response.Message = "Incorrect. Better luck next time!"
	}

	return response, nil
}

// checkYearRange implements smart year-range comparison
// Rules:
// - "present" in stored range upper bound is replaced with current year
// - Single user year vs stored range: passes if year falls within range
// - User range vs stored single year: always fails
// - Exact match always passes
func checkYearRange(submitted string, stored string) bool {
	if stored == "" {
		return false
	}

	currentYear := time.Now().Year()

	// Normalize "present" in stored value
	normalizedStored := strings.ToLower(strings.TrimSpace(stored))
	normalizedStored = strings.ReplaceAll(normalizedStored, "present", fmt.Sprintf("%d", currentYear))

	storedParts := splitYearRange(normalizedStored)
	submittedParts := splitYearRange(strings.TrimSpace(submitted))

	if len(submittedParts) == 0 || len(storedParts) == 0 {
		return false
	}

	// Case: user submitted a range, but stored is a single year → fail
	if len(submittedParts) == 2 && len(storedParts) == 1 {
		return false
	}

	// Case: user submitted a single year
	if len(submittedParts) == 1 {
		userYear := submittedParts[0]
		if len(storedParts) == 1 {
			// Both single years → exact match
			return userYear == storedParts[0]
		}
		// Stored is a range → check if user year falls within
		return userYear >= storedParts[0] && userYear <= storedParts[1]
	}

	// Case: both are ranges → exact comparison
	if len(submittedParts) == 2 && len(storedParts) == 2 {
		return submittedParts[0] == storedParts[0] && submittedParts[1] == storedParts[1]
	}

	return false
}

// splitYearRange splits "2018-2022" or "2018–2022" into [2018, 2022], "2020" into [2020]
func splitYearRange(s string) []int {
	// Replace en-dash, em-dash with hyphen
	s = strings.ReplaceAll(s, "–", "-")
	s = strings.ReplaceAll(s, "—", "-")

	parts := strings.Split(s, "-")
	var years []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		year := 0
		_, err := fmt.Sscanf(p, "%d", &year)
		if err != nil {
			continue
		}
		years = append(years, year)
	}
	return years
}


// updateBonusPoints adds bonus points to user_challenge_scores
func updateBonusPoints(tx *sqlx.Tx, userID int, challengeID int, points float64) error {
	// Check if score record exists
	var exists bool
	err := tx.Get(&exists, "SELECT EXISTS(SELECT 1 FROM user_challenge_scores WHERE user_id = $1 AND challenge_id = $2)", userID, challengeID)
	if err != nil {
		return err
	}

	if exists {
		// Update existing record
		updateQuery := `
			UPDATE user_challenge_scores 
			SET bonus_round_points = bonus_round_points + $3,
				total_points = total_points + $3,
				updated_at = CURRENT_TIMESTAMP
			WHERE user_id = $1 AND challenge_id = $2
		`
		_, err = tx.Exec(updateQuery, userID, challengeID, points)
	} else {
		// Create new record with just bonus points
		insertQuery := `
			INSERT INTO user_challenge_scores 
			(user_id, challenge_id, attempt_number, bonus_round_points, total_points, is_fully_solved, make_ever_wrong)
			VALUES ($1, $2, 0, $3, $3, FALSE, FALSE)
		`
		_, err = tx.Exec(insertQuery, userID, challengeID, points)
	}

	return err
}

// UpdateUserActivity updates the participation or submission bitmap for a user
func UpdateUserActivity(db *sqlx.DB, userID int, date time.Time, activityType string) error {
	// activityType should be 'participation' or 'submission'
	day := date.Day()
	// Always use UTC for the month bucket to avoid timezone shifting in DB DATE columns
	monthDate := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)

	column := "participation_bitmap"
	if activityType == "submission" {
		column = "submission_bitmap"
	}

	query := fmt.Sprintf(`
		INSERT INTO user_activity_bitmaps (user_id, month_date, %s)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, month_date)
		DO UPDATE SET %s = user_activity_bitmaps.%s | $3, updated_at = CURRENT_TIMESTAMP
	`, column, column, column)

	bitmask := 1 << (day - 1)
	_, err := db.Exec(query, userID, monthDate, bitmask)
	return err
}

// GetUserActivityStats calculates the streak and participation stats for a user
func GetUserActivityStats(db *sqlx.DB, userID int, now time.Time) (*models.UserActivityStats, error) {
	// Fetch all bitmaps for the user, ordered by date descending
	query := `
		SELECT month_date, participation_bitmap, submission_bitmap
		FROM user_activity_bitmaps
		WHERE user_id = $1
		ORDER BY month_date DESC
	`
	var rows []struct {
		MonthDate            time.Time `db:"month_date"`
		ParticipationBitmap int       `db:"participation_bitmap"`
		SubmissionBitmap    int       `db:"submission_bitmap"`
	}

	err := db.Select(&rows, query, userID)
	if err != nil {
		return nil, err
	}

	stats := &models.UserActivityStats{}
	
	if len(rows) == 0 {
		return stats, nil
	}

	// Calculate Current Streaks
	stats.AttendanceStreak = calculateCurrentStreak(rows, now, "participation")
	stats.SubmissionStreak = calculateCurrentStreak(rows, now, "submission")

	// Calculate Max Streaks and Totals
	maxAttendance := 0
	currentAttendance := 0
	maxSubmission := 0
	currentSubmission := 0
	totalParticipation := 0
	totalSubmission := 0

	// We need to iterate from oldest to newest to find max streaks correctly
	for i := len(rows) - 1; i >= 0; i-- {
		pMap := rows[i].ParticipationBitmap
		sMap := rows[i].SubmissionBitmap
		
		daysInMonth := getDaysInMonth(rows[i].MonthDate)
		
		for day := 1; day <= daysInMonth; day++ {
			// Participation
			if (pMap & (1 << (day - 1))) != 0 {
				currentAttendance++
				totalParticipation++
				if currentAttendance > maxAttendance {
					maxAttendance = currentAttendance
				}
			} else {
				currentAttendance = 0
			}

			// Submission
			if (sMap & (1 << (day - 1))) != 0 {
				currentSubmission++
				totalSubmission++
				if currentSubmission > maxSubmission {
					maxSubmission = currentSubmission
				}
			} else {
				currentSubmission = 0
			}
		}
	}

	stats.MaxAttendanceStreak = maxAttendance
	stats.MaxSubmissionStreak = maxSubmission
	stats.TotalDaysParticipated = totalParticipation
	stats.TotalDaysSubmitted = totalSubmission

	return stats, nil
}

func calculateCurrentStreak(rows []struct {
	MonthDate            time.Time `db:"month_date"`
	ParticipationBitmap int       `db:"participation_bitmap"`
	SubmissionBitmap    int       `db:"submission_bitmap"`
}, now time.Time, activityType string) int {
	streak := 0
	currentDate := now
	
	// Check if today or yesterday is the start
	// If it's 10 AM today and they haven't opened yet, the streak still includes yesterday.
	
	for {
		found := false
		// Find the row for the current month
		// We use Year/Month comparison to avoid timezone mismatch issues since DB stores DATE (no zone)
		// and Go retrieves it usually as UTC, while currentDate is formatted to User's Local Time.
		targetYear, targetMonth, _ := currentDate.Date()
		
		var bitmap int
		hasBitmap := false
		for _, row := range rows {
			// Compare year and month in UTC to match our storage convention
			rowUTC := row.MonthDate.UTC()
			if rowUTC.Year() == targetYear && rowUTC.Month() == targetMonth {
				if activityType == "participation" {
					bitmap = row.ParticipationBitmap
				} else {
					bitmap = row.SubmissionBitmap
				}
				hasBitmap = true
				break
			}
		}
		
		if !hasBitmap {
			break
		}
		
		dayBit := 1 << (currentDate.Day() - 1)
		if (bitmap & dayBit) != 0 {
			streak++
			currentDate = currentDate.AddDate(0, 0, -1)
			found = true
		} else {
			// If we didn't find today's bit, check if we are looking at 'now'
			// If so, maybe they just haven't played today yet, so check yesterday
			if streak == 0 && currentDate.Year() == now.Year() && currentDate.Month() == now.Month() && currentDate.Day() == now.Day() {
				currentDate = currentDate.AddDate(0, 0, -1)
				continue
			}
			break
		}
		
		if !found {
			break
		}
	}
	
	return streak
}

func getDaysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}

// GetUserBonusAttempts returns the user's bonus attempts for a challenge
func GetUserBonusAttempts(db *sqlx.DB, userID int, challengeID int) ([]models.BonusSubmission, error) {
	query := `
		SELECT id, user_id, challenge_id, bonus_type, submitted_value, is_correct, created_at
		FROM bonus_submissions
		WHERE user_id = $1 AND challenge_id = $2
	`

	var attempts []models.BonusSubmission
	err := db.Select(&attempts, query, userID, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bonus attempts: %v", err)
	}

	return attempts, nil
}

// GetChallengeStats calculates global stats for a specific challenge and lifetime stats
func GetChallengeStats(db *sqlx.DB, challengeID int) (*models.ChallengeStats, error) {
	stats := &models.ChallengeStats{ChallengeID: challengeID}

	// Players today: Count unique users in submissions for THIS challenge
	err := db.Get(&stats.PlayersToday, "SELECT COUNT(DISTINCT user_id) FROM user_challenge_scores WHERE challenge_id = $1", challengeID)
	if err != nil {
		return nil, err
	}

	// Total players: Count unique users in user_challenge_scores across ALL challenges
	err = db.Get(&stats.TotalPlayers, "SELECT COUNT(DISTINCT user_id) FROM user_challenge_scores")
	if err != nil {
		return nil, err
	}

	// Average accuracy for today: use the same per-user formula as the leaderboard
	// accuracy = 110 - (attempt_number * 10) for solved, 0 for unsolved, averaged across users
	err = db.Get(&stats.AverageAccuracy, `
		SELECT COALESCE(
			AVG(CASE 
				WHEN is_fully_solved THEN GREATEST(0, LEAST(100, 110.0 - attempt_number * 10.0))
				ELSE 0 
			END),
			0.0
		) FROM user_challenge_scores WHERE challenge_id = $1`, challengeID)
	if err != nil {
		return nil, err
	}

	// Global Average accuracy: same formula across ALL challenges
	err = db.Get(&stats.GlobalAverageAccuracy, `
		SELECT COALESCE(
			AVG(CASE 
				WHEN is_fully_solved THEN GREATEST(0, LEAST(100, 110.0 - attempt_number * 10.0))
				ELSE 0 
			END),
			0.0
		) FROM user_challenge_scores`)
	if err != nil {
		return nil, err
	}

	// Total bonus points for THIS challenge
	err = db.Get(&stats.TotalBonusPoints, "SELECT COALESCE(SUM(bonus_round_points), 0.0) FROM user_challenge_scores WHERE challenge_id = $1", challengeID)
	if err != nil {
		return nil, err
	}

	return stats, nil
}


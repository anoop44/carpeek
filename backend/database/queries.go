package database

import (
	"carpeek-backend/models"
	"database/sql"
	"fmt"
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
			mo.codename as model_codename
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
		MakeID          *int    `db:"make_id"`          // Use pointer to handle possible NULL
		MakeName        *string `db:"make_name"`        // Use pointer to handle possible NULL
		ModelID         *int    `db:"model_id"`         // Use pointer to handle possible NULL
		ModelName       *string `db:"model_name"`       // Use pointer to handle possible NULL
		ModelYearRange  *string `db:"model_year_range"` // Use pointer to handle possible NULL
		ModelGeneration *string `db:"model_generation"` // Use pointer to handle possible NULL
		ModelLocation   *string `db:"model_location"`   // Use pointer to handle possible NULL
		ModelCodename   *string `db:"model_codename"`   // Use pointer to handle possible NULL
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
func ValidateSubmission(db *sqlx.DB, challengeID int, submittedMakeID int, submittedModelID int) (*models.SubmissionResult, error) {
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
			mo.codename as model_codename
		FROM challenges c
		LEFT JOIN makes m ON c.solution_make_id = m.id
		LEFT JOIN models mo ON c.solution_model_id = mo.id
		WHERE c.id = $1
	`

	var result struct {
		ID                      int       `db:"id"`
		Date                    time.Time `db:"date"`
		ImageURL                string    `db:"image_url"`
		SolutionMakeID          *int    `db:"solution_make_id"`
		SolutionModelID         *int    `db:"solution_model_id"`
		SolutionMakeName        *string `db:"make_name"`
		SolutionModelName       *string `db:"model_name"`
		SolutionModelYearRange  *string `db:"model_year_range"`
		SolutionModelGeneration *string `db:"model_generation"`
		SolutionModelLocation   *string `db:"model_location"`
		SolutionModelCodename   *string `db:"model_codename"`
	}

	err := db.Get(&result, query, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge for validation: %v", err)
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

	solMakeID := 0
	if result.SolutionMakeID != nil {
		solMakeID = *result.SolutionMakeID
	}

	// Get the submitted model name for comparison
	var submittedModelName string
	err = db.Get(&submittedModelName, "SELECT name FROM models WHERE id = $1", submittedModelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submitted model name: %v", err)
	}

	// Check if both make and model are correct
	if submittedMakeID == solMakeID && result.SolutionModelName != nil && submittedModelName == *result.SolutionModelName {
		submissionResult.Correct = true
		submissionResult.Partial = false
		submissionResult.Message = "Perfect! You correctly identified the car."
	} else if submittedMakeID == solMakeID || (result.SolutionModelName != nil && submittedModelName == *result.SolutionModelName) {
		// Partially correct - either make or model is correct
		submissionResult.Correct = false
		submissionResult.Partial = true
		if submittedMakeID == solMakeID {
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

	return submissionResult, nil
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
	query := `SELECT id, anonymous_id, created_at, updated_at FROM users WHERE anonymous_id = $1`
	err := db.Get(&user, query, anonymousID)
	if err == nil {
		return &user, nil
	}

	// Create new user if not found
	insertQuery := `INSERT INTO users (anonymous_id) VALUES ($1) RETURNING id, anonymous_id, created_at, updated_at`
	err = db.QueryRowx(insertQuery, anonymousID).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &user, nil
}

// RecordSubmission saves a user's guess for a challenge
func RecordSubmission(db *sqlx.DB, userID int, challengeID int, makeID int, modelID int, isCorrect bool) error {
	query := `
		INSERT INTO submissions (user_id, challenge_id, make_id, model_id, is_correct)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(query, userID, challengeID, makeID, modelID, isCorrect)
	if err != nil {
		return fmt.Errorf("failed to record submission: %v", err)
	}
	return nil
}

// GetUserChallengeStatus returns the current status of a user for a specific challenge
func GetUserChallengeStatus(db *sqlx.DB, userID int, challengeID int) (*models.UserChallengeStatus, error) {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM submissions WHERE user_id = $1 AND challenge_id = $2", userID, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get submission count: %v", err)
	}

	var isCorrect bool
	err = db.Get(&isCorrect, "SELECT COALESCE(bool_or(is_correct), false) FROM submissions WHERE user_id = $1 AND challenge_id = $2", userID, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if completed correctly: %v", err)
	}

	maxAttempts := 5 // This could be dynamic in the future

	status := &models.UserChallengeStatus{
		Attempts:    count,
		MaxAttempts: maxAttempts,
		IsCompleted: isCorrect || count >= maxAttempts,
		IsCorrect:   isCorrect,
	}

	return status, nil
}

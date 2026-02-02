package models

import (
	"database/sql"
	"time"
)

// Make represents a car manufacturer
type Make struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Model represents a car model
type Model struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	MakeID     int       `json:"make_id" db:"make_id"`
	YearRange  *string   `json:"year_range" db:"year_range"`
	Generation *string   `json:"generation" db:"generation"`
	Location   *string   `json:"location" db:"location"`
	Codename   *string   `json:"codename" db:"codename"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// User represents an anonymous user
type User struct {
	ID          int       `json:"id" db:"id"`
	AnonymousID string    `json:"anonymous_id" db:"anonymous_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// GroupedModel represents a car model with multiple possible IDs
type GroupedModel struct {
	ID   []int  `json:"id"`
	Name string `json:"name"`
}

// Challenge represents a daily car challenge
type Challenge struct {
	ID              int       `json:"id" db:"id"`
	Date            time.Time `json:"date" db:"date"`
	ImageURL        string    `json:"image_url" db:"image_url"`
	SolutionMakeID  sql.NullInt64 `json:"solution_make_id" db:"solution_make_id"`
	SolutionModelID sql.NullInt64 `json:"solution_model_id" db:"solution_model_id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// ChallengeResponse represents the response for daily challenge API
type ChallengeResponse struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
}

// DetailedChallengeResponse represents a detailed response for daily challenge API
type DetailedChallengeResponse struct {
	ID         int    `json:"id"`
	Date       string `json:"date"`
	ImageURL   string `json:"image_url"`
	Make       *Make  `json:"make,omitempty"`
	Model      *Model `json:"model,omitempty"`
	UserStatus *UserChallengeStatus `json:"user_status,omitempty"`
}

// UserChallengeStatus represents the user's progress on a challenge
type UserChallengeStatus struct {
	Attempts      int  `json:"attempts"`
	MaxAttempts   int  `json:"max_attempts"`
	IsCompleted   bool `json:"is_completed"`
	IsCorrect     bool `json:"is_correct"`
}

// SubmissionRequest represents the request for challenge validation
type SubmissionRequest struct {
	ChallengeID int `json:"challenge_id" validate:"required"`
	MakeID      int `json:"make_id" validate:"required"`
	ModelID     int `json:"model_id" validate:"required"`
}

// SubmissionResult represents the response for challenge validation
type SubmissionResult struct {
	ID         int    `json:"id"`
	Date       string `json:"date"`
	Correct    bool   `json:"correct"`
	Partial    bool   `json:"partial"` // True if make is correct but model is wrong, or vice versa
	Message    string `json:"message"`
	ImageURL   string `json:"image_url"`
	Solution   *SolutionInfo `json:"solution,omitempty"`
	UserStatus *UserChallengeStatus `json:"user_status,omitempty"`
}

// SolutionInfo contains information about the correct solution
type SolutionInfo struct {
	MakeName   string  `json:"make_name"`
	ModelName  string  `json:"model_name"`
	YearRange  *string `json:"year_range,omitempty"`
	Generation *string `json:"generation,omitempty"`
	Codename   *string `json:"codename,omitempty"`
}
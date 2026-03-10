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
	ImageURL   *string   `json:"image_url" db:"image_url"`
	KnownFor   *string   `json:"known_for" db:"known_for"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// User represents a user (anonymous or authenticated)
type User struct {
	ID                    int        `json:"id" db:"id"`
	AnonymousID           string     `json:"anonymous_id" db:"anonymous_id"`
	GoogleID              *string    `json:"google_id,omitempty" db:"google_id"`
	Email                 *string    `json:"email,omitempty" db:"email"`
	DisplayName           *string    `json:"display_name,omitempty" db:"display_name"`
	ProfilePictureURL     *string    `json:"profile_picture_url,omitempty" db:"profile_picture_url"`
	IsLinked              bool       `json:"is_linked" db:"is_linked"`
	IsSubscriber          bool       `json:"is_subscriber" db:"is_subscriber"`
	SubscriptionStatus    *string    `json:"subscription_status,omitempty" db:"subscription_status"`
	SubscriptionProductID *string    `json:"subscription_product_id,omitempty" db:"subscription_product_id"`
	SubscriptionProvider  *string    `json:"subscription_provider,omitempty" db:"subscription_provider"`
	SubscriptionExpiresAt *time.Time `json:"subscription_expires_at,omitempty" db:"subscription_expires_at"`
	CreatedAt             time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at" db:"updated_at"`
}

// GroupedModel represents a car model with multiple possible IDs
type GroupedModel struct {
	ID   []int  `json:"id"`
	Name string `json:"name"`
}

// Challenge represents a daily car challenge
type Challenge struct {
	ID              int           `json:"id" db:"id"`
	Date            time.Time     `json:"date" db:"date"`
	ImageURL        string        `json:"image_url" db:"image_url"`
	SolutionMakeID  sql.NullInt64 `json:"solution_make_id" db:"solution_make_id"`
	SolutionModelID sql.NullInt64 `json:"solution_model_id" db:"solution_model_id"`
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" db:"updated_at"`
}

// ChallengeResponse represents the response for daily challenge API
type ChallengeResponse struct {
	ID                   int                `json:"id"`
	ImageURL             string             `json:"image_url"`
	NextChallengeSeconds int64              `json:"next_challenge_seconds"`
	StreakStats          *UserActivityStats `json:"streak_stats,omitempty"`
}

// DetailedChallengeResponse represents a detailed response for daily challenge API
type DetailedChallengeResponse struct {
	ID                   int                  `json:"id"`
	Date                 string               `json:"date"`
	ImageURL             string               `json:"image_url"`
	Make                 *Make                `json:"make,omitempty"`
	Model                *Model               `json:"model,omitempty"`
	UserStatus           *UserChallengeStatus `json:"user_status,omitempty"`
	NextChallengeSeconds int64                `json:"next_challenge_seconds"`
	StreakStats          *UserActivityStats   `json:"streak_stats,omitempty"`
	PointsEarned         float64              `json:"points_earned"`
	BonusRound           *BonusRoundInfo      `json:"bonus_round,omitempty"`
	AttemptHistory       []Submission         `json:"attempt_history,omitempty"`
}

// UserChallengeStatus represents the user's progress on a challenge
type UserChallengeStatus struct {
	Attempts    int  `json:"attempts"`
	MaxAttempts int  `json:"max_attempts"`
	IsCompleted bool `json:"is_completed"`
	IsCorrect   bool `json:"is_correct"`
}

type SubmissionRequest struct {
	ChallengeID int   `json:"challenge_id" validate:"required"`
	MakeID      int   `json:"make_id" validate:"required"`
	ModelIDs    []int `json:"model_ids" validate:"required"`
}

// SubmissionResult represents the response for challenge validation
type SubmissionResult struct {
	ID             int                  `json:"id"`
	Date           string               `json:"date"`
	Correct        bool                 `json:"correct"`
	Partial        bool                 `json:"partial"` // True if make is correct but model is wrong, or vice versa
	IsMakeCorrect  bool                 `json:"is_make_correct"`
	IsModelCorrect bool                 `json:"is_model_correct"`
	Message        string               `json:"message"`
	ImageURL       string               `json:"image_url"`
	Solution       *SolutionInfo        `json:"solution,omitempty"`
	UserStatus     *UserChallengeStatus `json:"user_status,omitempty"`
	BonusRound           *BonusRoundInfo      `json:"bonus_round,omitempty"` // Present if bonus rounds are available
	PointsEarned         float64              `json:"points_earned"`
	NextChallengeSeconds int64                `json:"next_challenge_seconds"`
	StreakStats          *UserActivityStats   `json:"streak_stats,omitempty"`
	AttemptHistory       []Submission         `json:"attempt_history,omitempty"`
}

// GoogleLoginRequest represents the request for Google login
type GoogleLoginRequest struct {
	IDToken     string `json:"id_token"`
	AnonymousID string `json:"anonymous_id"`
}

// SolutionInfo contains information about the correct solution
type SolutionInfo struct {
	MakeName   string  `json:"make_name"`
	ModelName  string  `json:"model_name"`
	YearRange  *string `json:"year_range,omitempty"`
	Generation *string `json:"generation,omitempty"`
	Codename   *string `json:"codename,omitempty"`
	KnownFor   *string `json:"known_for,omitempty"`
	ImageURL   *string `json:"image_url,omitempty"`
}

// LeaderboardEntry represents a user's standing on the leaderboard
type LeaderboardEntry struct {
	Rank              int     `json:"rank"`
	UserID            string  `json:"user_id"` // Anonymous ID or generated name
	PilotName         string  `json:"pilot_name"`
	ProfilePictureURL string  `json:"profile_picture_url,omitempty"`
	Level             int     `json:"level"`
	LevelTitle        string  `json:"level_title"`
	Score             float64 `json:"score"` // For daily: score in that challenge. For all-time: total score
	MainScore         float64 `json:"main_score"`
	BonusScore        float64 `json:"bonus_score"`
	Accuracy          float64 `json:"accuracy"` // Percentage
	Attempts          int     `json:"attempts"` // For daily: attempts used
	Time              string  `json:"time"`     // Formatting string like "1.2s" or date
	IsCurrentUser     bool    `json:"is_current_user"`
}

// Submission represents a user's guess for a challenge
type Submission struct {
	ID             int       `json:"id" db:"id"`
	UserID         int       `json:"user_id" db:"user_id"`
	ChallengeID    int       `json:"challenge_id" db:"challenge_id"`
	MakeID         int       `json:"make_id" db:"make_id"`
	ModelID        int       `json:"model_id" db:"model_id"`
	IsCorrect      bool      `json:"is_correct" db:"is_correct"`
	IsMakeCorrect  bool      `json:"is_make_correct" db:"is_make_correct"`
	IsModelCorrect bool      `json:"is_model_correct" db:"is_model_correct"`
	AttemptNumber  int       `json:"attempt_number" db:"attempt_number"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// UserChallengeScore tracks the calculated score for a user on a challenge
// Points system:
// - 5 points: Both correct on 1st attempt
// - 3 points: Both correct on 2nd attempt
// - 1 point: Both correct on 3rd attempt
// - 0.5 points: Make correct in latest submission (only if consistently correct across all attempts)
// - 1 point each: Bonus rounds (year_range, generation, codename)
type UserChallengeScore struct {
	ID               int       `json:"id" db:"id"`
	UserID           int       `json:"user_id" db:"user_id"`
	ChallengeID      int       `json:"challenge_id" db:"challenge_id"`
	AttemptNumber    int       `json:"attempt_number" db:"attempt_number"`
	FullSolvePoints  float64   `json:"full_solve_points" db:"full_solve_points"`
	MakeBonusPoints  float64   `json:"make_bonus_points" db:"make_bonus_points"`
	BonusRoundPoints float64   `json:"bonus_round_points" db:"bonus_round_points"`
	TotalPoints      float64   `json:"total_points" db:"total_points"`
	IsFullySolved    bool      `json:"is_fully_solved" db:"is_fully_solved"`
	MakeEverWrong    bool      `json:"make_ever_wrong" db:"make_ever_wrong"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// BonusRoundInfo indicates which bonus rounds are available for a challenge
// Based on model data: if year_range, generation, or codename have multiple distinct values
type BonusRoundInfo struct {
	YearRangeEnabled    bool    `json:"year_range_enabled"`
	GenerationEnabled   bool    `json:"generation_enabled"`
	CodenameEnabled     bool    `json:"codename_enabled"`
	YearRangeAttempted  bool    `json:"year_range_attempted"` // true if user already attempted
	GenerationAttempted bool    `json:"generation_attempted"`
	CodenameAttempted   bool    `json:"codename_attempted"`
	YearRangePoints     float64 `json:"year_range_points"`
	GenerationPoints    float64 `json:"generation_points"`
	CodenamePoints      float64 `json:"codename_points"`
	YearRangeCorrect    bool    `json:"year_range_correct"`
	GenerationCorrect   bool    `json:"generation_correct"`
	CodenameCorrect     bool    `json:"codename_correct"`
}

// BonusSubmission represents a bonus round attempt
type BonusSubmission struct {
	ID             int       `json:"id" db:"id"`
	UserID         int       `json:"user_id" db:"user_id"`
	ChallengeID    int       `json:"challenge_id" db:"challenge_id"`
	BonusType      string    `json:"bonus_type" db:"bonus_type"` // year_range, generation, codename
	SubmittedValue string    `json:"submitted_value" db:"submitted_value"`
	IsCorrect      bool      `json:"is_correct" db:"is_correct"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// BonusSubmissionRequest represents the request for bonus submission
type BonusSubmissionRequest struct {
	ChallengeID int    `json:"challenge_id" validate:"required"`
	BonusType   string `json:"bonus_type" validate:"required"` // year_range, generation, codename
	Value       string `json:"value" validate:"required"`
}

// BonusSubmissionResult represents the response for bonus submission
type BonusSubmissionResult struct {
	Correct      bool    `json:"correct"`
	Message      string  `json:"message"`
	CorrectValue string  `json:"correct_value"` // Always included for feedback
	PointsEarned float64 `json:"points_earned"`
}

// UserActivityStats represents the streak and participation data
type UserActivityStats struct {
	AttendanceStreak     int `json:"attendance_streak"`      // Continuous days opened
	SubmissionStreak     int `json:"submission_streak"`      // Continuous days guessed
	MaxAttendanceStreak  int `json:"max_attendance_streak"`  // All-time best attendance
	MaxSubmissionStreak  int `json:"max_submission_streak"`  // All-time best submission
	TotalDaysParticipated int `json:"total_days_participated"` // Total days opened life-time
	TotalDaysSubmitted    int `json:"total_days_submitted"`    // Total days guessed life-time
}

// ChallengeStats represents aggregated stats for a specific challenge
type ChallengeStats struct {
	ChallengeID           int     `json:"challenge_id"`
	PlayersToday          int     `json:"players_today"`
	TotalPlayers          int     `json:"total_players"`
	AverageAccuracy       float64 `json:"average_accuracy"`
	GlobalAverageAccuracy float64 `json:"global_average_accuracy"`
	TotalBonusPoints      float64 `json:"total_bonus_points"`
}

// SubscriptionStatusResponse represents the subscription status for a user
type SubscriptionStatusResponse struct {
	IsSubscriber         bool    `json:"is_subscriber"`
	SubscriptionStatus   *string `json:"subscription_status,omitempty"`
	SubscriptionProvider *string `json:"subscription_provider,omitempty"`
	ExpiresAt            *string `json:"expires_at,omitempty"`
}

// CreateRazorpaySubscriptionResponse represents the response when creating a Razorpay subscription
type CreateRazorpaySubscriptionResponse struct {
	SubscriptionID string `json:"subscription_id"`
	KeyID          string `json:"key_id"` // Frontend needs this to open checkout
}

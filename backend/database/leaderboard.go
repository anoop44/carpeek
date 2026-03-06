package database

import (
	"autocorrect-backend/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// GetDailyLeaderboard returns the leaderboard for a specific challenge
// Now uses the user_challenge_scores table for accurate point-based scoring
func GetDailyLeaderboard(db *sqlx.DB, challengeID int) ([]models.LeaderboardEntry, error) {
	// Query to get users who attempted the challenge, ordered by points
	// Also fetches lifetime points to determine level, and Google profile info
	query := `
		SELECT 
			u.anonymous_id,
			COALESCE(u.display_name, '') as display_name,
			COALESCE(u.profile_picture_url, '') as profile_picture_url,
			ucs.attempt_number,
			ucs.total_points,
			ucs.is_fully_solved,
			ucs.full_solve_points,
			ucs.make_bonus_points,
			ucs.bonus_round_points,
			(SELECT COALESCE(SUM(total_points), 0) FROM user_challenge_scores WHERE user_id = u.id) as lifetime_points
		FROM user_challenge_scores ucs
		JOIN users u ON ucs.user_id = u.id
		WHERE ucs.challenge_id = $1
		ORDER BY ucs.total_points DESC, ucs.attempt_number ASC
		LIMIT 100
	`

	var rows []struct {
		AnonymousID       string  `db:"anonymous_id"`
		DisplayName       string  `db:"display_name"`
		ProfilePictureURL string  `db:"profile_picture_url"`
		AttemptNumber     int     `db:"attempt_number"`
		TotalPoints       float64 `db:"total_points"`
		IsFullySolved     bool    `db:"is_fully_solved"`
		FullSolvePoints   float64 `db:"full_solve_points"`
		MakeBonusPoints   float64 `db:"make_bonus_points"`
		BonusRoundPoints  float64 `db:"bonus_round_points"`
		LifetimePoints    float64 `db:"lifetime_points"`
	}

	err := db.Select(&rows, query, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily leaderboard: %v", err)
	}

	leaderboard := make([]models.LeaderboardEntry, len(rows))
	for i, row := range rows {
		// Use Google display name if available, otherwise generate a pilot name
		pilotName := generatePilotName(row.AnonymousID)
		if row.DisplayName != "" {
			pilotName = row.DisplayName
		}
		
		// Simulate a "Velocity" or time based on attempts and random variation per user
		baseTime := float64(row.AttemptNumber) * 0.8
		jitter := float64(hashStringToInt(row.AnonymousID)%200) / 100.0
		timeSeconds := baseTime + jitter

		level, title := calculateLevel(row.LifetimePoints)

		leaderboard[i] = models.LeaderboardEntry{
			Rank:              i + 1,
			UserID:            row.AnonymousID,
			PilotName:         pilotName,
			ProfilePictureURL: row.ProfilePictureURL,
			Level:             level,
			LevelTitle:        title,
			Score:             row.TotalPoints,
			MainScore:         row.FullSolvePoints + row.MakeBonusPoints,
			BonusScore:        row.BonusRoundPoints,
			Accuracy:          calculateAccuracy(row.AttemptNumber, row.IsFullySolved),
			Attempts:          row.AttemptNumber,
			Time:              fmt.Sprintf("%.3fs", timeSeconds),
		}
	}

	return leaderboard, nil
}

// GetAllTimeLeaderboard returns the top users by total points across all challenges
func GetAllTimeLeaderboard(db *sqlx.DB) ([]models.LeaderboardEntry, error) {
	query := `
		SELECT 
			u.anonymous_id,
			COALESCE(u.display_name, '') as display_name,
			COALESCE(u.profile_picture_url, '') as profile_picture_url,
			SUM(ucs.total_points) as total_points,
			COUNT(ucs.id) as challenges_attempted,
			SUM(CASE 
				WHEN ucs.is_fully_solved THEN (110.0 - ucs.attempt_number * 10.0) 
				ELSE 0 
			END) as total_accuracy_points
		FROM user_challenge_scores ucs
		JOIN users u ON ucs.user_id = u.id
		GROUP BY u.anonymous_id, u.display_name, u.profile_picture_url
		ORDER BY total_points DESC, total_accuracy_points DESC
		LIMIT 100
	`

	var rows []struct {
		AnonymousID         string  `db:"anonymous_id"`
		DisplayName         string  `db:"display_name"`
		ProfilePictureURL   string  `db:"profile_picture_url"`
		TotalPoints         float64 `db:"total_points"`
		ChallengesAttempted int     `db:"challenges_attempted"`
		TotalAccuracyPoints float64 `db:"total_accuracy_points"`
	}

	err := db.Select(&rows, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all-time leaderboard: %v", err)
	}

	leaderboard := make([]models.LeaderboardEntry, len(rows))
	for i, row := range rows {
		// Use Google display name if available, otherwise generate a pilot name
		pilotName := generatePilotName(row.AnonymousID)
		if row.DisplayName != "" {
			pilotName = row.DisplayName
		}
		
		// Calculate accuracy as average of individual accuracies
		var accuracy float64 = 0
		if row.ChallengesAttempted > 0 {
			accuracy = row.TotalAccuracyPoints / float64(row.ChallengesAttempted)
		}

		level, title := calculateLevel(row.TotalPoints)
		
		leaderboard[i] = models.LeaderboardEntry{
			Rank:              i + 1,
			UserID:            row.AnonymousID,
			PilotName:         pilotName,
			ProfilePictureURL: row.ProfilePictureURL,
			Level:             level,
			LevelTitle:        title,
			Score:             row.TotalPoints,
			Accuracy:          accuracy,
			Attempts:          row.ChallengesAttempted,
			Time:              "-",
		}
	}

	return leaderboard, nil
}

// Helpers

func generatePilotName(uuid string) string {
	// Deterministic name generation
	adjectives := []string{"CYBER", "NEON", "VECTOR", "TURBO", "HYPER", "MEGA", "GIGA", "NIGHT", "SOLAR", "LUNAR", "APEX", "DELTA", "OMEGA", "ALPHA", "PRIME"}
	nouns := []string{"DRIFT", "RACER", "PILOT", "HAWK", "EAGLE", "WOLF", "GHOST", "PHANTOM", "SHADOW", "BLADE", "RUNNER", "STORM", "FURY", "VORTEX", "NOVA"}
	
	val := hashStringToInt(uuid)
	adj := adjectives[val % len(adjectives)]
	noun := nouns[(val/len(adjectives)) % len(nouns)]
	suffix := val % 99
	
	return fmt.Sprintf("%s_%s_%02d", adj, noun, suffix)
}

func hashStringToInt(s string) int {
	h := 0
	for _, c := range s {
		h = 31*h + int(c)
	}
	if h < 0 {
		h = -h
	}
	return h
}

func calculateAccuracy(attempts int, isSolved bool) float64 {
	// If not solved, accuracy is 0
	if !isSolved {
		return 0
	}
	
	// 1 attempt = 100%
	// 2 attempts = 90%
	// 3 attempts = 80%
	if attempts <= 0 {
		return 0
	}
	acc := 110.0 - (float64(attempts) * 10.0)
	if acc > 100 {
		return 100
	}
	if acc < 0 {
		return 0
	}
	return acc
}

func calculateLevel(points float64) (int, string) {
	if points >= 1100 {
		return 7, "Showroom Savant"
	} else if points >= 775 {
		return 6, "Garage Insider"
	} else if points >= 525 {
		return 5, "Trim Detective"
	} else if points >= 325 {
		return 4, "Sharp Eyed"
	} else if points >= 175 {
		return 3, "Car Spotter"
	} else if points >= 75 {
		return 2, "Street Watcher"
	} else if points >= 25 {
		return 1, "Casual Observer"
	} else {
		return 0, "Missed the Spot"
	}
}

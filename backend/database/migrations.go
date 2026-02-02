package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// MigrationRecord represents a record in schema_migrations table
type MigrationRecord struct {
	Version   string
	AppliedAt string
}

// getMigrationsPath returns the path to migrations folder from environment variable or fallbacks
func getMigrationsPath() string {
	path := os.Getenv("MIGRATIONS_PATH")

	// 1. Check if the environment variable is set and valid
	if path != "" {
		if _, err := os.Stat(path); err == nil {
			return path
		}
		log.Printf("│ ⚠ MIGRATIONS_PATH is set to %s but directory not found", path)
	}

	// 2. Fallback 1: Docker/Backend folder context
	if _, err := os.Stat("./migrations"); err == nil {
		log.Println("│ ℹ Using fallback migration path: ./migrations")
		return "./migrations"
	}

	// 3. Fallback 2: Project root context
	if _, err := os.Stat("./backend/migrations"); err == nil {
		log.Println("│ ℹ Using fallback migration path: ./backend/migrations")
		return "./backend/migrations"
	}

	// 4. Fallback 3: Hardcoded absolute path for Docker
	if _, err := os.Stat("/app/migrations"); err == nil {
		log.Println("│ ℹ Using fallback migration path: /app/migrations")
		return "/app/migrations"
	}

	if path == "" {
		log.Fatal("│ ✗ MIGRATIONS_PATH environment variable is not set and no fallbacks found")
	} else {
		log.Fatalf("│ ✗ Migrations directory not found at %s and no fallbacks found", path)
	}
	return path
}

// ensureMigrationsTable creates the schema_migrations table if it doesn't exist
func ensureMigrationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %v", err)
	}
	return nil
}

// getAppliedMigrations returns a map of already applied migrations
func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	applied := make(map[string]bool)

	rows, err := db.Query("SELECT version FROM schema_migrations ORDER BY version")
	if err != nil {
		return nil, fmt.Errorf("failed to query applied migrations: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("failed to scan migration version: %v", err)
		}
		applied[version] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating applied migrations: %v", err)
	}

	return applied, nil
}

// extractMigrationVersion extracts the version number from a migration filename
// e.g., "001_up.sql" -> "001", "002_create_users_up.sql" -> "002_create_users"
func extractMigrationVersion(filename string) string {
	// Remove .sql extension
	name := strings.TrimSuffix(filename, ".sql")

	// Remove _up or _down suffix
	if strings.HasSuffix(name, "_up") {
		return strings.TrimSuffix(name, "_up")
	}
	if strings.HasSuffix(name, "_down") {
		return strings.TrimSuffix(name, "_down")
	}

	// No suffix, treat the whole name as version
	return name
}

// isUpMigration checks if a file is an up migration
// Files ending with _down.sql are down migrations, everything else is up
func isUpMigration(filename string) bool {
	name := strings.TrimSuffix(filename, ".sql")
	return !strings.HasSuffix(name, "_down")
}

// isDownMigration checks if a file is a down migration
func isDownMigration(filename string) bool {
	name := strings.TrimSuffix(filename, ".sql")
	return strings.HasSuffix(name, "_down")
}

// getMigrationFiles returns a sorted list of migration files of a specific type
func getMigrationFiles(migrationsPath string, getUpMigrations bool) ([]string, error) {
	entries, err := os.ReadDir(migrationsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %v", err)
	}

	var migrations []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		if getUpMigrations && isUpMigration(entry.Name()) {
			migrations = append(migrations, entry.Name())
		} else if !getUpMigrations && isDownMigration(entry.Name()) {
			migrations = append(migrations, entry.Name())
		}
	}

	// Sort migrations by filename to ensure correct order
	sort.Strings(migrations)

	return migrations, nil
}

// applyMigration applies a single migration file
func applyMigration(db *sql.DB, migrationsPath, filename string, isUp bool) error {
	filePath := filepath.Join(migrationsPath, filename)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file %s: %v", filename, err)
	}

	sqlContent := strings.TrimSpace(string(content))
	if sqlContent == "" {
		log.Printf("│   ⚠ Migration %s: Empty file, skipping", filename)
		return nil
	}

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction for %s: %v", filename, err)
	}

	// Execute migration SQL
	_, err = tx.Exec(sqlContent)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute migration %s: %v", filename, err)
	}

	version := extractMigrationVersion(filename)

	if isUp {
		// Record migration as applied
		_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %v", filename, err)
		}
	} else {
		// Remove migration record
		_, err = tx.Exec("DELETE FROM schema_migrations WHERE version = $1", version)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to remove migration record %s: %v", filename, err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit migration %s: %v", filename, err)
	}

	return nil
}

// RunMigrations runs all pending up migrations
func RunMigrations(db *sql.DB) error {
	startTime := time.Now()

	log.Println("┌─────────────────────────────────────────────────────────────┐")
	log.Println("│ DATABASE MIGRATIONS                                         │")
	log.Println("├─────────────────────────────────────────────────────────────┤")

	// Ensure migrations table exists
	log.Println("│ ► Checking schema_migrations table...")
	if err := ensureMigrationsTable(db); err != nil {
		log.Printf("│ ✗ Failed to create migrations table: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return err
	}
	log.Println("│ ✓ schema_migrations table ready")

	// Get path to migrations
	migrationsPath := getMigrationsPath()
	log.Println("│")
	log.Printf("│ ► Migrations path: %s", migrationsPath)

	// Check if migrations directory exists
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		log.Printf("│ ✗ Migrations directory not found at %s", migrationsPath)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return fmt.Errorf("migrations directory not found at %s", migrationsPath)
	}
	log.Println("│ ✓ Migrations directory exists")

	// Get list of up migration files
	migrationFiles, err := getMigrationFiles(migrationsPath, true)
	if err != nil {
		log.Printf("│ ✗ Failed to read migration files: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return err
	}

	log.Println("│")
	if len(migrationFiles) == 0 {
		log.Println("│ ⚠ No migration files found")
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return nil
	}

	log.Printf("│ ► Found %d up migration file(s):", len(migrationFiles))
	for _, f := range migrationFiles {
		log.Printf("│   • %s", f)
	}

	// Get already applied migrations
	applied, err := getAppliedMigrations(db)
	if err != nil {
		log.Printf("│ ✗ Failed to get applied migrations: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return err
	}

	log.Println("│")
	log.Printf("│ ► Previously applied migrations: %d", len(applied))
	if len(applied) > 0 {
		for version := range applied {
			log.Printf("│   • %s", version)
		}
	}

	// Apply pending migrations
	log.Println("│")
	pendingCount := 0
	appliedInThisRun := []string{}

	for _, filename := range migrationFiles {
		version := extractMigrationVersion(filename)
		if applied[version] {
			log.Printf("│ ─ %s [already applied]", filename)
			continue
		}

		log.Printf("│ ► Applying: %s...", filename)
		migrationStart := time.Now()

		if err := applyMigration(db, migrationsPath, filename, true); err != nil {
			log.Printf("│ ✗ Migration failed: %v", err)
			log.Println("└─────────────────────────────────────────────────────────────┘")
			return fmt.Errorf("migration failed: %v", err)
		}

		migrationElapsed := time.Since(migrationStart)
		log.Printf("│ ✓ %s applied successfully (%v)", filename, migrationElapsed.Round(time.Millisecond))
		appliedInThisRun = append(appliedInThisRun, filename)
		pendingCount++
	}

	// Summary
	elapsed := time.Since(startTime)
	log.Println("│")
	log.Println("├─────────────────────────────────────────────────────────────┤")
	log.Println("│ MIGRATION SUMMARY                                           │")
	log.Println("├─────────────────────────────────────────────────────────────┤")

	if pendingCount == 0 {
		log.Println("│ ✓ Database is up to date - no migrations needed")
	} else {
		log.Printf("│ ✓ Successfully applied %d new migration(s):", pendingCount)
		for _, f := range appliedInThisRun {
			log.Printf("│   • %s", f)
		}
	}

	log.Printf("│ ► Total migrations in database: %d", len(applied)+pendingCount)
	log.Printf("│ ► Migration process took: %v", elapsed.Round(time.Millisecond))
	log.Println("└─────────────────────────────────────────────────────────────┘")

	return nil
}

// RunDownMigrations runs down migrations to rollback
// If targetVersion is empty, rolls back all migrations
// If targetVersion is specified, rolls back to that version (exclusive)
func RunDownMigrations(db *sql.DB, targetVersion string) error {
	startTime := time.Now()

	log.Println("┌─────────────────────────────────────────────────────────────┐")
	log.Println("│ DATABASE ROLLBACK                                           │")
	log.Println("├─────────────────────────────────────────────────────────────┤")

	if targetVersion != "" {
		log.Printf("│ ► Target version: %s", targetVersion)
	} else {
		log.Println("│ ⚠ No target version specified - rolling back ALL migrations")
	}

	// Ensure migrations table exists
	if err := ensureMigrationsTable(db); err != nil {
		log.Printf("│ ✗ Failed to ensure migrations table: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return err
	}

	// Get path to migrations
	migrationsPath := getMigrationsPath()
	log.Printf("│ ► Migrations path: %s", migrationsPath)

	// Check if migrations directory exists
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		log.Printf("│ ✗ Migrations directory not found at %s", migrationsPath)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return fmt.Errorf("migrations directory not found at %s", migrationsPath)
	}

	// Get list of down migration files (sorted in ascending order)
	migrationFiles, err := getMigrationFiles(migrationsPath, false)
	if err != nil {
		log.Printf("│ ✗ Failed to read migration files: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return err
	}

	if len(migrationFiles) == 0 {
		log.Println("│ ⚠ No down migration files found")
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return nil
	}

	// Reverse the order for rollback (latest first)
	for i, j := 0, len(migrationFiles)-1; i < j; i, j = i+1, j-1 {
		migrationFiles[i], migrationFiles[j] = migrationFiles[j], migrationFiles[i]
	}

	log.Println("│")
	log.Printf("│ ► Found %d down migration file(s):", len(migrationFiles))
	for _, f := range migrationFiles {
		log.Printf("│   • %s", f)
	}

	// Get applied migrations
	applied, err := getAppliedMigrations(db)
	if err != nil {
		log.Printf("│ ✗ Failed to get applied migrations: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return err
	}

	log.Println("│")
	log.Printf("│ ► Currently applied migrations: %d", len(applied))

	// Rollback applied migrations
	log.Println("│")
	rollbackCount := 0
	rolledBack := []string{}

	for _, filename := range migrationFiles {
		version := extractMigrationVersion(filename)

		// Stop if we've reached the target version
		if targetVersion != "" && version == targetVersion {
			log.Printf("│ ► Reached target version %s, stopping rollback", targetVersion)
			break
		}

		if !applied[version] {
			log.Printf("│ ─ %s [not applied, skipping]", filename)
			continue
		}

		log.Printf("│ ► Rolling back: %s...", filename)
		rollbackStart := time.Now()

		if err := applyMigration(db, migrationsPath, filename, false); err != nil {
			log.Printf("│ ✗ Rollback failed: %v", err)
			log.Println("└─────────────────────────────────────────────────────────────┘")
			return fmt.Errorf("rollback failed: %v", err)
		}

		rollbackElapsed := time.Since(rollbackStart)
		log.Printf("│ ✓ %s rolled back successfully (%v)", filename, rollbackElapsed.Round(time.Millisecond))
		rolledBack = append(rolledBack, filename)
		rollbackCount++
	}

	// Summary
	elapsed := time.Since(startTime)
	log.Println("│")
	log.Println("├─────────────────────────────────────────────────────────────┤")
	log.Println("│ ROLLBACK SUMMARY                                            │")
	log.Println("├─────────────────────────────────────────────────────────────┤")

	if rollbackCount == 0 {
		log.Println("│ ⚠ No migrations to rollback")
	} else {
		log.Printf("│ ✓ Successfully rolled back %d migration(s):", rollbackCount)
		for _, f := range rolledBack {
			log.Printf("│   • %s", f)
		}
	}

	log.Printf("│ ► Remaining migrations: %d", len(applied)-rollbackCount)
	log.Printf("│ ► Rollback process took: %v", elapsed.Round(time.Millisecond))
	log.Println("└─────────────────────────────────────────────────────────────┘")

	return nil
}

// GetMigrationStatus returns the current migration status
func GetMigrationStatus(db *sql.DB) ([]MigrationRecord, error) {
	if err := ensureMigrationsTable(db); err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT version, applied_at FROM schema_migrations ORDER BY version")
	if err != nil {
		return nil, fmt.Errorf("failed to query migration status: %v", err)
	}
	defer rows.Close()

	var records []MigrationRecord
	for rows.Next() {
		var record MigrationRecord
		if err := rows.Scan(&record.Version, &record.AppliedAt); err != nil {
			return nil, fmt.Errorf("failed to scan migration record: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}

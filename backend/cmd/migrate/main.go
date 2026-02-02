package main

import (
	"flag"
	"fmt"
	"log"

	"carpeek-backend/database"

	"github.com/joho/godotenv"
)

func main() {
	// Define flags
	down := flag.Bool("down", false, "Run down migrations (rollback)")
	target := flag.String("target", "", "Target version to rollback to (used with -down)")
	status := flag.Bool("status", false, "Show migration status")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if *status {
		showStatus(db)
		return
	}

	if *down {
		runDownMigrations(db, *target)
		return
	}

	// Default: run up migrations
	runUpMigrations(db)
}

func printHelp() {
	fmt.Println("CarPeek Database Migration Tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  migrate [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -down          Run down migrations (rollback)")
	fmt.Println("  -target=XXX    Target version to rollback to (used with -down)")
	fmt.Println("  -status        Show current migration status")
	fmt.Println("  -help          Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  migrate                    # Run all pending up migrations")
	fmt.Println("  migrate -status            # Show applied migrations")
	fmt.Println("  migrate -down              # Rollback all migrations")
	fmt.Println("  migrate -down -target=001  # Rollback to version 001")
	fmt.Println()
	fmt.Println("Environment Variables:")
	fmt.Println("  MIGRATIONS_PATH  Path to migrations folder (required)")
	fmt.Println("  DATABASE_URL     PostgreSQL connection string")
}

func runUpMigrations(db interface{ Close() error }) {
	sqlDB := database.DB
	if err := database.RunMigrations(sqlDB); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}

func runDownMigrations(db interface{ Close() error }, target string) {
	sqlDB := database.DB
	if err := database.RunDownMigrations(sqlDB, target); err != nil {
		log.Fatal("Failed to run down migrations:", err)
	}
}

func showStatus(db interface{ Close() error }) {
	sqlDB := database.DB
	records, err := database.GetMigrationStatus(sqlDB)
	if err != nil {
		log.Fatal("Failed to get migration status:", err)
	}

	if len(records) == 0 {
		fmt.Println("No migrations have been applied yet.")
		return
	}

	fmt.Println("Applied Migrations:")
	fmt.Println("-------------------")
	for _, record := range records {
		fmt.Printf("  %s (applied at: %s)\n", record.Version, record.AppliedAt)
	}
	fmt.Printf("\nTotal: %d migration(s) applied\n", len(records))
}

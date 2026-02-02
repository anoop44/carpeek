package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"carpeek-backend/api"
	"carpeek-backend/database"

	"github.com/joho/godotenv"
)

const (
	appName    = "CarPeek Backend"
	appVersion = "1.0.0"
)

func main() {
	startTime := time.Now()

	// Print startup banner
	printBanner()

	log.Println("╔══════════════════════════════════════════════════════════════╗")
	log.Println("║                    APPLICATION STARTUP                       ║")
	log.Println("╚══════════════════════════════════════════════════════════════╝")

	// Print system information
	log.Println("┌─────────────────────────────────────────────────────────────┐")
	log.Println("│ SYSTEM INFORMATION                                          │")
	log.Println("├─────────────────────────────────────────────────────────────┤")
	log.Printf("│ ► Go Version:    %s", runtime.Version())
	log.Printf("│ ► OS/Arch:       %s/%s", runtime.GOOS, runtime.GOARCH)
	log.Printf("│ ► CPUs:          %d", runtime.NumCPU())
	log.Printf("│ ► Startup Time:  %s", startTime.Format("2006-01-02 15:04:05 MST"))
	log.Println("└─────────────────────────────────────────────────────────────┘")

	// Load environment variables
	log.Println("")
	log.Println("┌─────────────────────────────────────────────────────────────┐")
	log.Println("│ LOADING ENVIRONMENT                                         │")
	log.Println("├─────────────────────────────────────────────────────────────┤")

	if err := godotenv.Load(); err != nil {
		log.Println("│ ⚠ No .env file found, using system environment variables")
	} else {
		log.Println("│ ✓ Loaded environment variables from .env file")
	}

	// Log key environment variables (without sensitive values)
	logEnvironmentVariables()
	log.Println("└─────────────────────────────────────────────────────────────┘")

	// Connect to database
	log.Println("")
	db, err := database.ConnectDB()
	if err != nil {
		log.Println("╔══════════════════════════════════════════════════════════════╗")
		log.Println("║                    ✗ STARTUP FAILED                          ║")
		log.Println("╚══════════════════════════════════════════════════════════════╝")
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Run migrations
	log.Println("")
	if err := database.Migrate(db); err != nil {
		log.Println("╔══════════════════════════════════════════════════════════════╗")
		log.Println("║                    ✗ STARTUP FAILED                          ║")
		log.Println("╚══════════════════════════════════════════════════════════════╝")
		log.Fatal("Migration failed:", err)
	}

	// Create router
	router := api.NewRouter()

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Print server ready message
	elapsed := time.Since(startTime)
	log.Println("")
	log.Println("╔══════════════════════════════════════════════════════════════╗")
	log.Println("║                    ✓ SERVER READY                            ║")
	log.Println("╠══════════════════════════════════════════════════════════════╣")
	log.Printf("║ ► Listening on:  http://0.0.0.0:%s", port)
	log.Printf("║ ► Startup took:  %v", elapsed.Round(time.Millisecond))
	log.Println("╚══════════════════════════════════════════════════════════════╝")
	log.Println("")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func printBanner() {
	log.Println("")
	log.Println("  ╔═══════════════════════════════════════════════════════════╗")
	log.Println("  ║                                                           ║")
	log.Printf("  ║            %s v%s                       ║", appName, appVersion)
	log.Println("  ║                                                           ║")
	log.Println("  ║        Daily Car Identification Challenge Game            ║")
	log.Println("  ║                                                           ║")
	log.Println("  ╚═══════════════════════════════════════════════════════════╝")
	log.Println("")
}

func logEnvironmentVariables() {
	// Log environment variables (mask sensitive ones)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		log.Println("│ ✓ DATABASE_URL:      [SET]")
	} else {
		log.Println("│ ⚠ DATABASE_URL:      [NOT SET - using default]")
	}

	port := os.Getenv("PORT")
	if port != "" {
		log.Printf("│ ✓ PORT:              %s", port)
	} else {
		log.Println("│ ⚠ PORT:              [NOT SET - using default 8080]")
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath != "" {
		log.Printf("│ ✓ MIGRATIONS_PATH:   %s", migrationsPath)
	} else {
		log.Println("│ ✗ MIGRATIONS_PATH:   [NOT SET - REQUIRED]")
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL != "" {
		log.Printf("│ ✓ BASE_URL:          %s", baseURL)
	} else {
		log.Println("│ ⚠ BASE_URL:          [NOT SET - using relative paths]")
	}
}
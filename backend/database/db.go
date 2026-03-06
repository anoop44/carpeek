package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() (*sql.DB, error) {
	log.Println("┌─────────────────────────────────────────────────────────────┐")
	log.Println("│ DATABASE CONNECTION                                         │")
	log.Println("├─────────────────────────────────────────────────────────────┤")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("│ ⚠ DATABASE_URL not set, using default connection string")
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Parse and log connection details (without password)
	logConnectionDetails(dbURL)

	log.Println("│")
	log.Println("│ ► Attempting to connect to database...")

	startTime := time.Now()

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("│ ✗ Failed to open database connection: %v", err)
		log.Println("└─────────────────────────────────────────────────────────────┘")
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("│ ► Connection pool configured:")
	log.Println("│   • Max open connections: 25")
	log.Println("│   • Max idle connections: 5")
	log.Println("│   • Connection max lifetime: 5m")

	// Test the connection with retries
	maxRetries := 5
	retryDelay := 2 * time.Second

	for i := 1; i <= maxRetries; i++ {
		log.Printf("│ ► Ping attempt %d/%d...", i, maxRetries)

		err = db.Ping()
		if err == nil {
			elapsed := time.Since(startTime)
			log.Println("│")
			log.Println("│ ✓ Database connection established successfully!")
			log.Printf("│ ► Connection time: %v", elapsed.Round(time.Millisecond))

			// Get database stats
			stats := db.Stats()
			log.Println("│")
			log.Println("│ ► Initial connection pool stats:")
			log.Printf("│   • Open connections: %d", stats.OpenConnections)
			log.Printf("│   • In use: %d", stats.InUse)
			log.Printf("│   • Idle: %d", stats.Idle)
			log.Println("└─────────────────────────────────────────────────────────────┘")

			DB = db
			return DB, nil
		}

		log.Printf("│ ⚠ Ping failed: %v", err)

		if i < maxRetries {
			log.Printf("│ ► Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	log.Printf("│ ✗ Failed to connect after %d attempts", maxRetries)
	log.Println("└─────────────────────────────────────────────────────────────┘")
	return nil, fmt.Errorf("failed to ping database after %d attempts: %v", maxRetries, err)
}

// logConnectionDetails logs the database connection details (without password)
func logConnectionDetails(dbURL string) {
	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		log.Println("│ ⚠ Could not parse DATABASE_URL")
		return
	}

	host := parsedURL.Hostname()
	port := parsedURL.Port()
	if port == "" {
		port = "5432"
	}
	dbName := ""
	if len(parsedURL.Path) > 1 {
		dbName = parsedURL.Path[1:]
	}
	user := ""
	if parsedURL.User != nil {
		user = parsedURL.User.Username()
	}

	log.Println("│")
	log.Println("│ ► Connection details:")
	log.Printf("│   • Host:     %s", host)
	log.Printf("│   • Port:     %s", port)
	log.Printf("│   • Database: %s", dbName)
	log.Printf("│   • User:     %s", user)
	log.Println("│   • Password: [HIDDEN]")

	// Log SSL mode if present
	sslMode := parsedURL.Query().Get("sslmode")
	if sslMode != "" {
		log.Printf("│   • SSL Mode: %s", sslMode)
	}
}

// Migrate runs all pending database migrations
func Migrate(db *sql.DB) error {
	return RunMigrations(db)
}

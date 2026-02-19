package utils

import (
	"log"
	"os"
	"strings"
)

var (
	isProduction bool
)

func init() {
	env := strings.ToLower(os.Getenv("APP_ENV"))
	if env == "" {
		env = strings.ToLower(os.Getenv("ENV"))
	}
	isProduction = (env == "production")
}

// LogEvent logs a major event if not in production
func LogEvent(category string, message string, details ...interface{}) {
	if isProduction {
		return
	}

	if len(details) > 0 {
		log.Printf("🔹 [EVENT][%s] %s | Details: %+v\n", strings.ToUpper(category), message, details)
	} else {
		log.Printf("🔹 [EVENT][%s] %s\n", strings.ToUpper(category), message)
	}
}

// LogDebug logs debug information if not in production
func LogDebug(context string, message string, details ...interface{}) {
	if isProduction {
		return
	}

	if len(details) > 0 {
		log.Printf("🔍 [DEBUG][%s] %s | Details: %+v\n", strings.ToUpper(context), message, details)
	} else {
		log.Printf("🔍 [DEBUG][%s] %s\n", strings.ToUpper(context), message)
	}
}

// LogAPI logs API request details if not in production
func LogAPI(method, path string, status int, duration string) {
	if isProduction {
		return
	}

	icon := "✅"
	if status >= 400 && status < 500 {
		icon = "⚠️"
	} else if status >= 500 {
		icon = "❌"
	}

	log.Printf("%s [API] %s %s | Status: %d | Duration: %s\n", icon, method, path, status, duration)
}

// LogError logs an error if not in production
func LogError(context string, err error) {
	if isProduction {
		return
	}

	if err != nil {
		log.Printf("🚨 [ERROR][%s] ⚠️ %v\n", strings.ToUpper(context), err)
	}
}

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONError writes a JSON error response with the given message and HTTP status code.
// This ensures all API responses are consistently JSON-formatted.
func JSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// JSONErrorf writes a formatted JSON error response.
func JSONErrorf(w http.ResponseWriter, statusCode int, format string, args ...interface{}) {
	JSONError(w, fmt.Sprintf(format, args...), statusCode)
}

// JSONResponse writes a JSON response with the given data and HTTP status code.
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		LogError("JSONResponse", err)
	}
}

package api

import (
	"autocorrect-backend/utils"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a custom response writer to capture the status code
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:      http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		utils.LogAPI(r.Method, r.URL.Path, rw.statusCode, duration.String())
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

package middleware

import (
	"autocorrect-backend/utils"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// clientLimiter holds the rate limiter and last actvity for a client
type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimitConfig holds the configuration for a specific endpoint limit
type RateLimitConfig struct {
	RequestsPerSecond float64
	Burst             int
}

var (
	// limiters stores the rate limiters for each client IP
	// map[string]*clientLimiter
	limiters = sync.Map{}

	// cleanupTicker runs periodically to remove old limiters
	cleanupTicker *time.Ticker
	cleanupDone   chan bool
)

func init() {
	// Start cleanup routine
	cleanupTicker = time.NewTicker(time.Minute)
	cleanupDone = make(chan bool)

	go func() {
		for {
			select {
			case <-cleanupDone:
				return
			case <-cleanupTicker.C:
				cleanupLimiters()
			}
		}
	}()
}

func cleanupLimiters() {
	limiters.Range(func(key, value interface{}) bool {
		client, ok := value.(*clientLimiter)
		if !ok {
			limiters.Delete(key)
			return true
		}
		// Remove if inactive for 3 minutes
		if time.Since(client.lastSeen) > 3*time.Minute {
			limiters.Delete(key)
		}
		return true
	})
}

// RateLimit returns a middleware that limits requests based on IP address
func RateLimit(requestsPerSecond float64, burst int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Identify client by IP
			ip := utils.GetClientIP(r) // We need this utility
			
			// Get or create limiter
			limiter := getLimiter(ip, requestsPerSecond, burst)
			
			if !limiter.Allow() {
				utils.JSONError(w, "Too many requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getLimiter(ip string, r float64, b int) *rate.Limiter {
	v, exists := limiters.Load(ip)
	if !exists {
		limiter := rate.NewLimiter(rate.Limit(r), b)
		limiters.Store(ip, &clientLimiter{
			limiter:  limiter,
			lastSeen: time.Now(),
		})
		return limiter
	}

	client := v.(*clientLimiter)
	client.lastSeen = time.Now()
	// If rate config changed for existing user, we might want to update, but typically we assume consistent config per endpoint
	// For simplicity, we just return the existing one. However, if used across different endpoints with different limits, 
	// the key should probably include the endpoint path or limit config.
	// But standard practice is per-IP global limit or per-IP-per-endpoint.
	// Let's use IP + Request Path, or just IP if generic. 
	// Given "customizable for individual apis", we should key by "IP:r:b" or similar to allow different limits.
	return client.limiter
}

// CustomRateLimit allows per-endpoint customization by including the limit parameters in the key
// This ensures separate buckets for high-volume vs sensitive endpoints
func CustomRateLimit(requestsPerSecond float64, burst int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := utils.GetClientIP(r)
			// Create a unique key for this limit configuration
			key := fmt.Sprintf("%s:%f:%d", ip, requestsPerSecond, burst)
			
			v, exists := limiters.Load(key)
			var limiter *rate.Limiter
			if !exists {
				limiter = rate.NewLimiter(rate.Limit(requestsPerSecond), burst)
				limiters.Store(key, &clientLimiter{
					limiter:  limiter,
					lastSeen: time.Now(),
				})
			} else {
				client := v.(*clientLimiter)
				client.lastSeen = time.Now()
				limiter = client.limiter
			}
			
			if !limiter.Allow() {
				utils.JSONError(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}

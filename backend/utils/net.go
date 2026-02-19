package utils

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIP retrieves the real client IP address from request headers
func GetClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	// X-Forwarded-For can contain multiple IPs, the left-most is the original client
	if strings.Contains(ip, ",") {
		ip = strings.Split(ip, ",")[0]
	}
	return strings.TrimSpace(ip)
}

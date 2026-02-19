package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

// SignValue signs a string value using HMAC-SHA256
// Format: value.signature
func SignValue(value, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(value))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("%s.%s", value, signature)
}

// VerifySignature validates a signed value and returns the original value if valid
func VerifySignature(signedValue, secret string) (string, bool) {
	parts := strings.Split(signedValue, ".")
	if len(parts) != 2 {
		return "", false
	}

	value := parts[0]
	signature := parts[1]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(value))
	expectedSignature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return value, true
	}

	return "", false
}

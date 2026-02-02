package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

// SignValue signs a string value using HMAC-SHA256
func SignValue(value, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(value))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s.%s", value, signature)
}

// VerifySignature verifies a signed string value
func VerifySignature(signedValue, secret string) (string, bool) {
	parts := strings.Split(signedValue, ".")
	if len(parts) != 2 {
		return "", false
	}

	value := parts[0]
	expectedSignature := parts[1]

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(value))
	actualSignature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if hmac.Equal([]byte(actualSignature), []byte(expectedSignature)) {
		return value, true
	}

	return "", false
}

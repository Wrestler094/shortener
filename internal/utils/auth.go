package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

const (
	CookieName   = "auth_token"
	cookieKey    = "supersecret"
	CookieMaxAge = 3600 * 24 * 365
)

func GenerateUserID() string {
	return uuid.NewString()
}

func Sign(value string) string {
	h := hmac.New(sha256.New, []byte(cookieKey))
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateSignedValue(value string) string {
	return value + "|" + Sign(value)
}

func ValidateSignedValue(signed string) (string, bool) {
	parts := len(signed)
	if parts == 0 {
		return "", false
	}
	split := []rune(signed)
	for i := range split {
		if split[i] == '|' {
			value := string(split[:i])
			signature := string(split[i+1:])
			if Sign(value) == signature {
				return value, true
			}
			return "", false
		}
	}
	return "", false
}

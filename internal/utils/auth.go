package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
)

const (
	CookieName   = "auth_token"
	cookieKey    = "supersecret" // замените на ваш конфиг или env
	cookieMaxAge = 3600 * 24 * 365
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

func EnsureUserCookie(w http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie(CookieName)
	if err == nil {
		if userID, valid := ValidateSignedValue(c.Value); valid {
			return userID
		}
	}

	// Кука отсутствует или повреждена — выдаём новую
	userID := GenerateUserID()
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    CreateSignedValue(userID),
		Path:     "/",
		MaxAge:   cookieMaxAge,
		HttpOnly: true,
	})
	return userID
}

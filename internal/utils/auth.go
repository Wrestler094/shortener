package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

const (
	// CookieName - имя cookie для хранения токена аутентификации
	CookieName = "auth_token"
	// cookieKey - секретный ключ для подписи cookie
	cookieKey = "supersecret"
	// CookieMaxAge - время жизни cookie в секундах (1 год)
	CookieMaxAge = 3600 * 24 * 365
)

// GenerateUserID генерирует уникальный идентификатор пользователя
// Возвращает строку с UUID
func GenerateUserID() string {
	return uuid.NewString()
}

// Sign создает HMAC-SHA256 подпись для значения
// value - значение для подписи
// Возвращает подпись в виде hex-строки
func Sign(value string) string {
	h := hmac.New(sha256.New, []byte(cookieKey))
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

// CreateSignedValue создает подписанное значение в формате "value|signature"
// value - значение для подписи
// Возвращает строку в формате "value|signature"
func CreateSignedValue(value string) string {
	return value + "|" + Sign(value)
}

// ValidateSignedValue проверяет подпись значения
// signed - строка в формате "value|signature"
// Возвращает:
// - оригинальное значение
// - флаг валидности подписи
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

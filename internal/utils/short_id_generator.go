package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/Wrestler094/shortener/internal/configs"
)

// GenerateShortID генерирует случайный короткий идентификатор для URL
// Использует криптографически безопасный генератор случайных чисел
// Возвращает:
// - строку длиной ShortURLLen символов в base64url кодировке
// - ошибку, если не удалось сгенерировать случайные байты
func GenerateShortID() (string, error) {
	bytes := make([]byte, configs.ShortURLLen)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:configs.ShortURLLen], nil
}

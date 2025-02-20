package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Wrestler094/shortener/internal/configs"
)

func GenerateShortID() (string, error) {
	bytes := make([]byte, configs.ShortURLLen)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:configs.ShortURLLen], nil
}

package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Wrestler094/shortener/internal/storage"
	"io"
	"net/http"
	"strings"
)

var (
	ServerAddr  = "http://localhost:8080/"
	ShortURLLen = 8
)

func generateShortID() (string, error) {
	bytes := make([]byte, ShortURLLen)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:ShortURLLen], nil
}

// Temporary handler
func URLHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" && req.Method == http.MethodPost {
		SaveURL(res, req)
		return
	}

	if req.Method == http.MethodGet {
		urlParts := strings.Split(req.URL.Path, "/")
		urlParts = urlParts[1:]

		if len(urlParts) != 1 || urlParts[0] == "" {
			http.Error(res, "Request without id of shorten url", http.StatusBadRequest)
			return
		}

		GetURL(res, req, urlParts[0])
		return
	}

	http.Error(res, "Bad Request", http.StatusBadRequest)
}

func SaveURL(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}
	originalURL := strings.TrimSpace(string(body))

	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		http.Error(res, "Invalid URL format", http.StatusBadRequest)
		return
	}

	shortID, err := generateShortID()
	if err != nil {
		http.Error(res, "Failed to generate short URL", http.StatusBadRequest)
		return
	}

	storage.Save(shortID, originalURL)

	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(ServerAddr + shortID))
}

func GetURL(res http.ResponseWriter, _ *http.Request, id string) {
	url, ok := storage.Get(id)
	if !ok {
		http.Error(res, "Shorten URL not found", http.StatusBadRequest)
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

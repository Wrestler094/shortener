package handlers

import (
	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/utils"
	"io"
	"net/http"
	"strings"
)

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

	shortID, err := utils.GenerateShortID()
	if err != nil {
		http.Error(res, "Failed to generate short URL", http.StatusBadRequest)
		return
	}

	// TODO: Сделать првоерку на случай если id или URL уже существует
	storage.Save(shortID, originalURL)

	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(configs.ServerAddr + shortID))
}

func GetURL(res http.ResponseWriter, req *http.Request) {
	urlParts := strings.Split(req.URL.Path, "/")
	urlParts = urlParts[1:]

	if len(urlParts) != 1 || urlParts[0] == "" {
		http.Error(res, "Request without id of shorten url", http.StatusBadRequest)
		return
	}

	url, ok := storage.Get(urlParts[0])
	if !ok {
		http.Error(res, "Shorten URL not found", http.StatusBadRequest)
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

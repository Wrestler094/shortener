package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

func SaveJSONURL(res http.ResponseWriter, req *http.Request) {
	var shortenRequest ShortenRequest

	err := json.NewDecoder(req.Body).Decode(&shortenRequest)
	if err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	shortID, err := services.SaveURL(shortenRequest.URL)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(ShortenResponse{
		Result: configs.FlagBaseAddr + "/" + shortID,
	})
}

func SavePlainURL(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortID, err := services.SaveURL(string(body))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(configs.FlagBaseAddr + "/" + shortID))
}

func GetURL(res http.ResponseWriter, req *http.Request) {
	urlParts := strings.Split(req.URL.Path, "/")
	urlParts = urlParts[1:]

	if len(urlParts) != 1 || urlParts[0] == "" {
		http.Error(res, "request without id of shorten url", http.StatusBadRequest)
		return
	}

	url, ok := storage.Get(urlParts[0])
	if !ok {
		http.Error(res, "shorten URL not found", http.StatusBadRequest)
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

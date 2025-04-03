package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/mailru/easyjson"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage/memory"
)

//easyjson:json
type ShortenRequest struct {
	URL string `json:"url"`
}

//easyjson:json
type ShortenResponse struct {
	Result string `json:"result"`
}

func SaveJSONURL(res http.ResponseWriter, req *http.Request) {
	var shortenRequest ShortenRequest

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := easyjson.Unmarshal(body, &shortenRequest); err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}

	shortID, err := services.SaveURL(shortenRequest.URL)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	responseBody, err := easyjson.Marshal(ShortenResponse{
		Result: configs.FlagBaseAddr + "/" + shortID,
	})
	if err != nil {
		http.Error(res, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	res.Write(responseBody)
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

	url, ok := memory.Get(urlParts[0])
	if !ok {
		http.Error(res, "shorten URL not found", http.StatusBadRequest)
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

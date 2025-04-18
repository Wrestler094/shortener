package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/middlewares"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage/postgres"
	"github.com/Wrestler094/shortener/internal/utils"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

type URLHandler struct {
	service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) SaveJSONURL(res http.ResponseWriter, req *http.Request) {
	var shortenRequest ShortenRequest

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &shortenRequest); err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, _ := middlewares.GetUserIDFromContext(req.Context())
	shortID, err := h.service.SaveURL(shortenRequest.URL, userID)
	if err != nil {
		if errors.Is(err, postgres.ErrURLAlreadyExists) {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusConflict)
			responseBody, err := json.Marshal(ShortenResponse{
				Result: configs.FlagBaseAddr + "/" + shortID,
			})
			if err != nil {
				http.Error(res, "Failed to encode response", http.StatusInternalServerError)
				return
			}
			res.Write(responseBody)
			return
		}

		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(res, http.StatusCreated, ShortenResponse{
		Result: configs.FlagBaseAddr + "/" + shortID,
	})
}

func (h *URLHandler) SavePlainURL(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, _ := middlewares.GetUserIDFromContext(req.Context())
	shortID, err := h.service.SaveURL(string(body), userID)
	if err != nil {
		if errors.Is(err, postgres.ErrURLAlreadyExists) {
			res.Header().Set("Content-Type", "text/plain")
			res.WriteHeader(http.StatusConflict)
			res.Write([]byte(configs.FlagBaseAddr + "/" + shortID))
			return
		}
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(configs.FlagBaseAddr + "/" + shortID))
}

func (h *URLHandler) SaveBatchURLs(res http.ResponseWriter, req *http.Request) {
	var batch []dto.BatchRequestItem

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Failed to read request", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	if err := json.Unmarshal(body, &batch); err != nil || len(batch) == 0 {
		http.Error(res, "Invalid JSON or empty batch", http.StatusBadRequest)
		return
	}

	userID, _ := middlewares.GetUserIDFromContext(req.Context())
	result, err := h.service.SaveBatch(batch, userID)
	if err != nil {
		http.Error(res, "Failed to process batch", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(res, http.StatusCreated, result)
}

func (h *URLHandler) GetURL(res http.ResponseWriter, req *http.Request) {
	urlParts := strings.Split(req.URL.Path, "/")
	urlParts = urlParts[1:]

	if len(urlParts) != 1 || urlParts[0] == "" {
		http.Error(res, "request without id of shorten url", http.StatusBadRequest)
		return
	}

	url, ok := h.service.GetURLByID(urlParts[0])
	if !ok {
		http.Error(res, "shorten URL not found", http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *URLHandler) GetUserURLs(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(utils.CookieName)
	if err != nil {
		logger.Log.Error("Error of open file", zap.Error(err), zap.Any("cookie", cookie))
		http.Error(res, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := utils.ValidateSignedValue(cookie.Value)
	if !ok {
		logger.Log.Error("Error of open file", zap.String("userID", userID))
		http.Error(res, "unauthorized", http.StatusUnauthorized)
		return
	}

	userURLs, err := h.service.GetUserURLs(userID)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(userURLs) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	utils.WriteJSON(res, http.StatusOK, userURLs)
}

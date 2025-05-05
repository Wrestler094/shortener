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

// ShortenRequest представляет структуру запроса на сокращение URL
type ShortenRequest struct {
	URL string `json:"url"` // URL для сокращения
}

// ShortenResponse представляет структуру ответа с сокращенным URL
type ShortenResponse struct {
	Result string `json:"result"` // Сокращенный URL
}

// URLHandler обрабатывает HTTP-запросы, связанные с URL
type URLHandler struct {
	service *services.URLService // Сервис для работы с URL
}

// NewURLHandler создает новый экземпляр URLHandler
func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{service: service}
}

// SaveJSONURL обрабатывает POST-запрос для сохранения URL в формате JSON.
// Принимает JSON с полем "url" и возвращает сокращенный URL.
// Возможные коды ответа:
// - 201: URL успешно сокращен
// - 400: Неверный формат запроса
// - 409: URL уже существует
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
			utils.WriteJSON(res, http.StatusConflict, ShortenResponse{
				Result: configs.FlagBaseAddr + "/" + shortID,
			})
			return
		}

		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(res, http.StatusCreated, ShortenResponse{
		Result: configs.FlagBaseAddr + "/" + shortID,
	})
}

// SavePlainURL обрабатывает POST-запрос для сохранения URL в текстовом формате.
// Принимает URL в теле запроса и возвращает сокращенный URL.
// Возможные коды ответа:
// - 201: URL успешно сокращен
// - 400: Неверный формат запроса
// - 409: URL уже существует
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

// SaveBatchURLs обрабатывает POST-запрос для пакетного сохранения URL.
// Принимает массив URL в формате JSON и возвращает массив сокращенных URL.
// Возможные коды ответа:
// - 201: URL успешно сокращены
// - 400: Неверный формат запроса
// - 500: Внутренняя ошибка сервера
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

// GetURL обрабатывает GET-запрос для получения оригинального URL по сокращенному идентификатору.
// Возможные коды ответа:
// - 307: Редирект на оригинальный URL
// - 400: Неверный формат запроса
// - 410: URL был удален
func (h *URLHandler) GetURL(res http.ResponseWriter, req *http.Request) {
	urlParts := strings.Split(req.URL.Path, "/")
	urlParts = urlParts[1:]

	if len(urlParts) != 1 || urlParts[0] == "" {
		http.Error(res, "request without id of shorten url", http.StatusBadRequest)
		return
	}

	url, isDeleted, ok := h.service.GetURLByID(urlParts[0])
	if !ok {
		http.Error(res, "shorten URL not found", http.StatusBadRequest)
		return
	}

	if isDeleted {
		res.WriteHeader(http.StatusGone)
		return
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

// GetUserURLs обрабатывает GET-запрос для получения всех URL пользователя.
// Возможные коды ответа:
// - 200: Список URL успешно получен
// - 204: У пользователя нет сохраненных URL
// - 401: Пользователь не авторизован
// - 500: Внутренняя ошибка сервера
func (h *URLHandler) GetUserURLs(res http.ResponseWriter, req *http.Request) {
	userID, ok := middlewares.GetUserIDFromContext(req.Context())
	if !ok || userID == "" {
		http.Error(res, "unauthorized", http.StatusUnauthorized)
		return
	}

	userURLs, err := h.service.GetUserURLs(userID)
	if err != nil {
		logger.Log.Error("Failed to get user URLs", zap.Error(err))
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(userURLs) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	utils.WriteJSON(res, http.StatusOK, userURLs)
}

// DeleteUserURLs обрабатывает DELETE-запрос для удаления URL пользователя.
// Принимает массив идентификаторов URL для удаления.
// Возможные коды ответа:
// - 202: Запрос на удаление принят
// - 400: Неверный формат запроса
// - 500: Внутренняя ошибка сервера
func (h *URLHandler) DeleteUserURLs(res http.ResponseWriter, req *http.Request) {
	var urlForDelete []string

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&urlForDelete); err != nil {
		logger.Log.Error("Failed to decode request DeleteUserURLs", zap.Error(err))
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, _ := middlewares.GetUserIDFromContext(req.Context())
	err := h.service.DeleteUserURLs(userID, urlForDelete)
	if err != nil {
		logger.Log.Error("Failed to delete user URLs", zap.Error(err))
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusAccepted)
}

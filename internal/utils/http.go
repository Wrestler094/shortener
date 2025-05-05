package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON записывает JSON-ответ в http.ResponseWriter
// w - ResponseWriter для записи ответа
// status - HTTP-статус ответа
// data - данные для сериализации в JSON
// В случае ошибки сериализации возвращает 500 Internal Server Error
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

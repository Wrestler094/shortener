package dto

// BatchRequestItem представляет элемент запроса на пакетное сохранение URL
type BatchRequestItem struct {
	CorrelationID string `json:"correlation_id"` // Идентификатор корреляции для связи запроса и ответа
	OriginalURL   string `json:"original_url"`   // Оригинальный URL для сокращения
}

// BatchResponseItem представляет элемент ответа на пакетное сохранение URL
type BatchResponseItem struct {
	CorrelationID string `json:"correlation_id"` // Идентификатор корреляции из запроса
	ShortURL      string `json:"short_url"`      // Сокращенный URL
}

package dto

// UserURLItem представляет URL пользователя в ответе на запрос списка URL
type UserURLItem struct {
	ShortURL    string `json:"short_url"`    // Сокращенный URL
	OriginalURL string `json:"original_url"` // Оригинальный URL
}

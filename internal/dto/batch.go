package dto

//easyjson:json
type BatchRequestList []BatchRequestItem

//easyjson:json
type BatchRequestItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

//easyjson:json
type BatchResponseList []BatchResponseItem

//easyjson:json
type BatchResponseItem struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

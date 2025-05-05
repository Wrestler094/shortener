package handlers

// Handlers содержит все обработчики HTTP-запросов приложения
type Handlers struct {
	URLHandler  *URLHandler  // Обработчик запросов для работы с URL
	PingHandler *PingHandler // Обработчик запросов для проверки доступности хранилища
}

// NewHandlers создает новый экземпляр Handlers
// url - обработчик запросов для работы с URL
// ping - обработчик запросов для проверки доступности хранилища
func NewHandlers(url *URLHandler, ping *PingHandler) *Handlers {
	return &Handlers{
		URLHandler:  url,
		PingHandler: ping,
	}
}

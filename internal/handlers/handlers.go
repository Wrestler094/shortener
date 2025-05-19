package handlers

// Handlers содержит все обработчики HTTP-запросов приложения
type Handlers struct {
	URLHandler   *URLHandler   // Обработчик запросов для работы с URL
	StatsHandler *StatsHandler // Обработчик запросов для получения статистики сервиса
	PingHandler  *PingHandler  // Обработчик запросов для проверки доступности хранилища
}

// NewHandlers создает новый экземпляр Handlers
// url - обработчик запросов для работы с URL
// ping - обработчик запросов для проверки доступности хранилища
// stats - обработчик запросов для работы со статистикой
func NewHandlers(url *URLHandler, ping *PingHandler, stats *StatsHandler) *Handlers {
	return &Handlers{
		URLHandler:   url,
		PingHandler:  ping,
		StatsHandler: stats,
	}
}

package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	httphandlers "github.com/Wrestler094/shortener/internal/handlers/http"
	"github.com/Wrestler094/shortener/internal/middlewares"
)

// InitRouter инициализирует и настраивает HTTP-роутер приложения
// @param handlers - структура, содержащая все обработчики HTTP-запросов приложения
// @return *chi.Mux - настроенный роутер с определенными маршрутами и middleware
//
// Доступные маршруты:
// - POST / - сохранение URL в текстовом формате
// - GET /{id} - получение оригинального URL по сокращенному
// - GET /ping - проверка доступности хранилища
// - POST /api/shorten - сохранение URL в JSON формате
// - POST /api/shorten/batch - пакетное сохранение URL
// - GET /api/user/urls - получение списка URL пользователя
// - DELETE /api/user/urls - удаление URL пользователя
// - GET /api/internal/stats - получение статистики сервиса
// - GET /swagger/* - документация API в формате Swagger
//
// Применяемые middleware:
// - RequestLogger - логирование HTTP-запросов
// - Compressor - сжатие ответов
// - AuthCookieSetter - установка cookie для аутентификации
// - Recoverer - восстановление после паники
func InitRouter(httphandlers *httphandlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.Compressor)
	r.Use(middlewares.AuthCookieSetter)
	r.Use(middleware.Recoverer)

	r.Post("/", httphandlers.URLHandler.SavePlainURL)
	r.Get("/{id}", httphandlers.URLHandler.GetURL)
	r.Get("/ping", httphandlers.PingHandler.Ping)

	r.Post("/api/shorten", httphandlers.URLHandler.SaveJSONURL)
	r.Post("/api/shorten/batch", httphandlers.URLHandler.SaveBatchURLs)

	r.Get("/api/user/urls", httphandlers.URLHandler.GetUserURLs)
	r.Delete("/api/user/urls", httphandlers.URLHandler.DeleteUserURLs)

	r.Get("/api/internal/stats", httphandlers.StatsHandler.GetStats)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}

package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/middlewares"
)

// InitRouter инициализирует и настраивает HTTP-роутер приложения
// handlers - обработчики HTTP-запросов
// Возвращает настроенный роутер со следующими маршрутами:
// - POST / - сохранение URL в текстовом формате
// - GET /{id} - получение оригинального URL по сокращенному
// - GET /ping - проверка доступности хранилища
// - POST /api/shorten - сохранение URL в JSON формате
// - POST /api/shorten/batch - пакетное сохранение URL
// - GET /api/user/urls - получение списка URL пользователя
// - DELETE /api/user/urls - удаление URL пользователя
func InitRouter(handlers *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.Compressor)
	r.Use(middlewares.AuthCookieSetter)
	r.Use(middleware.Recoverer)

	r.Post("/", handlers.URLHandler.SavePlainURL)
	r.Get("/{id}", handlers.URLHandler.GetURL)
	r.Get("/ping", handlers.PingHandler.Ping)

	r.Post("/api/shorten", handlers.URLHandler.SaveJSONURL)
	r.Post("/api/shorten/batch", handlers.URLHandler.SaveBatchURLs)

	r.Get("/api/user/urls", handlers.URLHandler.GetUserURLs)
	r.Delete("/api/user/urls", handlers.URLHandler.DeleteUserURLs)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}

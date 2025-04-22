package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/middlewares"
)

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

	return r
}

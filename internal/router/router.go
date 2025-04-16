package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/middlewares"
)

func InitRouter(urlHandler *handlers.URLHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.Compressor)
	r.Use(middleware.Recoverer)

	r.Post("/", urlHandler.SavePlainURL)
	r.Post("/api/shorten", urlHandler.SaveJSONURL)
	r.Get("/{id}", urlHandler.GetURL)

	// todo
	r.Get("/ping", handlers.HandlePing)

	return r
}

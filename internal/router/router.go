package router

import (
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.Compressor)
	r.Use(middleware.Recoverer)

	r.Post("/", handlers.SavePlainURL)
	r.Get("/{id}", handlers.GetURL)
	r.Post("/api/shorten", handlers.SaveJSONURL)

	r.Get("/ping", handlers.HandlePing)

	return r
}

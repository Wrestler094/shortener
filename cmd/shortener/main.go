package main

import (
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", handlers.SaveURL)
	r.Get("/{id}", handlers.GetURL)

	log.Fatal(http.ListenAndServe(":8080", r))
}

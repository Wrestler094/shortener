package main

import (
	"github.com/Wrestler094/shortener/internal/middlewares"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/logger"
)

func registerRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.RequestLogger)
	r.Use(middleware.Recoverer)

	r.Post("/", handlers.SavePlainURL)
	r.Get("/{id}", handlers.GetURL)
	r.Post("/api/shorten", handlers.SaveJSONURL)

	return r
}

func main() {
	configs.ParseFlags()
	configs.ParseEnv()

	if err := logger.Initialize(configs.LoggerLevel); err != nil {
		log.Fatal(err)
	}

	router := registerRouter()

	logger.Log.Info("Running server", zap.String("address", configs.FlagRunAddr))
	logger.Log.Fatal("Server crashed", zap.Error(http.ListenAndServe(configs.FlagRunAddr, router)))
}

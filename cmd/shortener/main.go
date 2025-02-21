package main

import (
	"flag"
	"fmt"
	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func parseFlags() {
	flag.StringVar(&configs.FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&configs.FlagBaseAddr, "b", "http://localhost:8080/", "basic address and port of result url")
	flag.Parse()
}

func registerRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", handlers.SaveURL)
	r.Get("/{id}", handlers.GetURL)

	return r
}

func main() {
	parseFlags()
	router := registerRouter()

	fmt.Println("Running server on", configs.FlagRunAddr)
	log.Fatal(http.ListenAndServe(configs.FlagRunAddr, router))
}

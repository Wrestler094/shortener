package main

import (
	"github.com/Wrestler094/shortener/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc(`/`, handlers.SaveUrl)
	//mux.HandleFunc(`/`+`:id`, handlers.GetUrl)
	mux.HandleFunc(`/`, handlers.URLHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

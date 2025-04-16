package app

import (
	"net/http"

	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/router"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/storage/memory"
)

type App struct {
	Router http.Handler
}

func InitApp() *App {
	// Выбор хранилища (можно сюда добавить file/postgres)
	var store storage.IStorage = memory.NewMemoryStorage()

	// Инициализация сервисов
	urlService := services.NewURLService(store)

	// Инициализация хендлеров
	urlHandler := handlers.NewURLHandler(urlService)

	// Создание роутера
	r := router.InitRouter(urlHandler)

	return &App{
		Router: r,
	}
}

//func InitApp() *App {
//	store := memory.NewMemoryStorage()
//	service := services.NewURLService(store)
//	urlHandler := handlers.NewURLHandler(service)
//	router := router.InitRouter(urlHandler)
//
//	return &App{Router: router}
//}

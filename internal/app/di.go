package app

import (
	"log"
	"net/http"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/router"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/storage/file"
	"github.com/Wrestler094/shortener/internal/storage/memory"
	"github.com/Wrestler094/shortener/internal/storage/postgres"
)

type App struct {
	Router http.Handler
}

func InitApp() *App {
	// Выбор хранилища (можно сюда добавить file/postgres)
	var store storage.IStorage

	if configs.FlagDatabaseDSN != "" {
		// todo ask gpt
		postgresStore, err := postgres.NewPostgresStorage(configs.FlagDatabaseDSN)
		store = postgresStore
		if err != nil {
			log.Fatal(err)
		}

	} else {
		store = memory.NewMemoryStorage()
		// todo: может быть восстанавливать и в Postgresql??
		// todo: Передавать строку до файла
		file.RecoverURLs(store)
	}

	// Инициализация сервисов
	urlService := services.NewURLService(store)

	// Инициализация хендлеров
	urlHandler := handlers.NewURLHandler(urlService)
	pingHandler := handlers.NewPingHandler(store)

	hs := &handlers.Handlers{
		URLHandler:  urlHandler,
		PingHandler: pingHandler,
	}

	// Создание роутера
	r := router.InitRouter(hs)

	return &App{Router: r}
}

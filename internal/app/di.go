package app

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/router"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/storage/memory"
	"github.com/Wrestler094/shortener/internal/storage/postgres"
)

type App struct {
	Router http.Handler
}

func InitApp() *App {
	// Выбор хранилища (можно сюда добавить file/postgres)
	var fileStorage = persistence.NewFileStorage(configs.FlagFileStoragePath)
	var store storage.IStorage

	if configs.FlagDatabaseDSN != "" {
		postgresStore, err := postgres.NewPostgresStorage(configs.FlagDatabaseDSN)
		if err != nil {
			logger.Log.Fatal("Failed to initialize postgres storage", zap.Error(err))
			log.Fatal(err)
		}
		store = postgresStore
	} else {
		recoveredUrls := fileStorage.RecoverURLs()
		store = memory.NewMemoryStorage(recoveredUrls)
	}

	// Инициализация сервисов
	urlService := services.NewURLService(store, fileStorage)

	// Инициализация хендлеров
	urlHandler := handlers.NewURLHandler(urlService)
	pingHandler := handlers.NewPingHandler(store)

	// Создание роутера
	hs := handlers.NewHandlers(urlHandler, pingHandler)
	r := router.InitRouter(hs)

	return &App{Router: r}
}

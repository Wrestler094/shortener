package app

import (
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/deleter"
	"github.com/Wrestler094/shortener/internal/handlers"
	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/router"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage"
	"github.com/Wrestler094/shortener/internal/storage/memory"
	"github.com/Wrestler094/shortener/internal/storage/postgres"
)

// App представляет основную структуру приложения
type App struct {
	Router  http.Handler     // HTTP-роутер приложения
	Storage storage.IStorage // Хранилище для работы с URL
	Deleter deleter.Deleter  // Компонент для асинхронного удаления URL
}

// InitApp инициализирует приложение, создавая все необходимые зависимости:
// 1. Выбирает и инициализирует хранилище (PostgreSQL или память)
// 2. Создает сервис для работы с URL
// 3. Создает обработчики HTTP-запросов
// 4. Инициализирует роутер
// Возвращает экземпляр приложения с настроенным роутером
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

	urlDeleter := deleter.NewURLDeleter(store, time.Second*10)
	urlDeleter.StartBackgroundFlusher()

	// Инициализация сервисов
	urlService := services.NewURLService(store, fileStorage, urlDeleter)
	statsService := services.NewStatsService(store)

	// Инициализация хендлеров
	urlHandler := handlers.NewURLHandler(urlService)
	pingHandler := handlers.NewPingHandler(store)
	statsHandler := handlers.NewStatsHandler(statsService)

	// Создание роутера
	hs := handlers.NewHandlers(urlHandler, pingHandler, statsHandler)
	r := router.InitRouter(hs)

	return &App{Router: r, Storage: store, Deleter: urlDeleter}
}

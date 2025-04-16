package app

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/logger"
)

/*
TODO
	1. Нужно ли добавлять логгер в DI?
	2. Сделать структуру для хэндлеров для DI
	3.
*/

//type Handlers struct {
//	URLHandler  *URLHandler
//	PingHandler *PingHandler
//	UserHandler *UserHandler
//	// и т.д.
//}

func Run() {
	configs.ParseFlags()
	configs.ParseEnv()

	// TODO: GRT норм что тут?
	if err := logger.Initialize(configs.LoggerLevel); err != nil {
		log.Fatal(err)
	}

	// TODO
	//	if err := postgres.InitPostgres(configs.FlagDatabaseDSN); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	storage.Storage = memory.NewMemoryStorage()
	//	initialisedRouter := router.InitRouter()
	//
	//	file.RecoverURLs()

	app := InitApp()

	logger.Log.Info("Running server", zap.String("address", configs.FlagRunAddr))
	logger.Log.Fatal("Server crashed", zap.Error(http.ListenAndServe(configs.FlagRunAddr, app.Router)))
}

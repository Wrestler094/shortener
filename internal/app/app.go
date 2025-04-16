package app

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/logger"
)

func Run() {
	configs.ParseFlags()
	configs.ParseEnv()

	if err := logger.Initialize(configs.LoggerLevel); err != nil {
		log.Fatal(err)
	}

	app := InitApp()

	logger.Log.Info("Running server", zap.String("address", configs.FlagRunAddr))
	logger.Log.Fatal("Server crashed", zap.Error(http.ListenAndServe(configs.FlagRunAddr, app.Router)))
}

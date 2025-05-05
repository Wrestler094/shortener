package app

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/logger"
)

// Run запускает приложение, инициализирует конфигурацию, логгер и HTTP-сервер.
// Функция выполняет следующие действия:
// 1. Парсит флаги командной строки
// 2. Парсит переменные окружения
// 3. Инициализирует логгер
// 4. Инициализирует приложение
// 5. Запускает HTTP-сервер на указанном адресе
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

package app

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/utils"
)

// Run запускает приложение, инициализирует конфигурацию, логгер и HTTP-сервер.
// Функция выполняет следующие действия:
// 1. Парсит флаги командной строки
// 2. Парсит переменные окружения
// 3. Инициализирует логгер
// 4. Инициализирует приложение
// 5. Запускает HTTP-сервер на указанном адресе
func Run() {
	configs.ParseJSON()
	configs.ParseFlags()
	configs.ParseEnv()

	if err := logger.Initialize(configs.LoggerLevel); err != nil {
		log.Fatal(err)
	}

	app := InitApp()

	if configs.FlagEnableHTTPS == "true" {
		// Создаём self-signed TLS сертификат
		cert, err := utils.GenerateSelfSignedCert()
		if err != nil {
			logger.Log.Fatal("Failed to generate cert", zap.Error(err))
		}

		port := fmt.Sprintf(":%d", configs.HTTPSPort)
		server := &http.Server{
			Addr:    port,
			Handler: app.Router,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}

		logger.Log.Info("Running server (HTTPS mode enabled (self-signed cert))", zap.String("address", port))
		logger.Log.Fatal("Server crashed", zap.Error(server.ListenAndServeTLS("", "")))
	} else {
		logger.Log.Info("Running server (HTTP mode enabled)", zap.String("address", configs.FlagRunAddr))
		logger.Log.Fatal("Server crashed", zap.Error(http.ListenAndServe(configs.FlagRunAddr, app.Router)))
	}
}

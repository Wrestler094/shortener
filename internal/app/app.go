package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/logger"
	"github.com/Wrestler094/shortener/internal/storage"
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
	configs.InitConfig()

	if err := logger.Initialize(configs.LoggerLevel); err != nil {
		log.Fatal(err)
	}

	app := InitApp()
	server := &http.Server{
		Addr:    configs.FlagRunAddr,
		Handler: app.Router,
	}

	// HTTPS (self-signed TLS)
	if configs.FlagEnableHTTPS {
		cert, err := utils.GenerateSelfSignedCert()
		if err != nil {
			logger.Log.Fatal("Failed to generate TLS cert", zap.Error(err))
		}
		server.Addr = fmt.Sprintf(":%d", configs.HTTPSPort)
		server.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	}

	// Сигнальный канал
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Запуск сервера в горутине
	go func() {
		logger.Log.Info("Server running", zap.String("addr", server.Addr))
		var err error
		if configs.FlagEnableHTTPS {
			err = server.ListenAndServeTLS("", "")
		} else {
			err = server.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server crashed", zap.Error(err))
		}
	}()

	// Ожидаем сигнал завершения
	<-stop
	logger.Log.Info("Shutdown signal received")

	// Завершаем deleter
	app.Deleter.Stop()

	// Завершаем HTTP-сервер
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error("Graceful shutdown failed", zap.Error(err))
	}

	// Закрываем БД (если есть метод Close)
	if closer, ok := app.Storage.(storage.IClosableStorage); ok {
		if err := closer.Close(); err != nil {
			logger.Log.Error("Failed to close storage", zap.Error(err))
		}
	}

	logger.Log.Info("Shutdown complete")
}

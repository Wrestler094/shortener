// Middleware для детального логирования HTTP-запросов.
// Основные возможности:
// - Логирование всех входящих HTTP-запросов с использованием структурированного логгера zap
// - Сбор метрик производительности (время выполнения запроса)
// - Отслеживание размера ответа и HTTP-статуса
// - Перехват и модификация ответа без изменения его содержимого
//
// Middleware использует кастомный ResponseWriter для перехвата информации
// об ответе, сохраняя при этом все оригинальные возможности стандартного
// http.ResponseWriter.
package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/logger"
)

// responseData хранит информацию об ответе HTTP-сервера
type responseData struct {
	status int // HTTP-статус ответа
	size   int // Размер тела ответа в байтах
}

// loggingResponseWriter реализует интерфейс http.ResponseWriter
// и позволяет перехватывать информацию об ответе
type loggingResponseWriter struct {
	http.ResponseWriter               // Встраиваем оригинальный http.ResponseWriter
	responseData        *responseData // Данные об ответе
}

// Write записывает данные в ответ и обновляет размер ответа
// b - данные для записи
// Возвращает количество записанных байт и ошибку
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

// WriteHeader устанавливает HTTP-статус ответа
// statusCode - код HTTP-статуса
func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

// RequestLogger - middleware для логирования входящих HTTP-запросов
// Логирует метод, путь, длительность обработки, статус и размер ответа
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: w, // встраиваем оригинальный http.ResponseWriter
			responseData:   responseData,
		}

		start := time.Now()
		next.ServeHTTP(&lw, r)
		duration := time.Since(start)

		logger.Log.Info("got incoming HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.RequestURI),
			zap.String("duration", strconv.FormatInt(int64(duration), 10)),
			zap.String("status", strconv.FormatInt(int64(responseData.status), 10)),
			zap.String("size", strconv.FormatInt(int64(responseData.size), 10)),
		)
	})
}

package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/Wrestler094/shortener/internal/logger"
)

// берём структуру для хранения сведений об ответе
type responseData struct {
	status int
	size   int
}

// добавляем реализацию http.ResponseWriter
type loggingResponseWriter struct {
	http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
	responseData        *responseData
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}

// RequestLogger — middleware-логер для входящих HTTP-запросов.
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

package middlewares

import (
	"context"
	"time"

	"github.com/Wrestler094/shortener/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// NewLoggingInterceptor создаёт unary-интерцептор для логирования запросов
func NewLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()

		resp, err = handler(ctx, req) // Выполняем сам хендлер

		code := status.Code(err)
		duration := time.Since(start)

		logger.Log.Info("gRPC request",
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
			zap.Duration("duration", duration),
			zap.String("status", code.String()),
			zap.Error(err),
		)

		return resp, err
	}
}

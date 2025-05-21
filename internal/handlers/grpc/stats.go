package grpchandlers

import (
	"context"

	"github.com/Wrestler094/shortener/internal/grpc/pb"
	"github.com/Wrestler094/shortener/internal/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

// StatsHandler реализует gRPC StatsService
type StatsHandler struct {
	pb.UnimplementedStatsServiceServer

	service *services.StatsService
}

// NewStatsHandler создает новый StatsHandler
func NewStatsHandler(s *services.StatsService) *StatsHandler {
	return &StatsHandler{service: s}
}

// GetStats реализует метод gRPC получения статистики
func (h *StatsHandler) GetStats(ctx context.Context, _ *emptypb.Empty) (*pb.StatsResponse, error) {
	urls, users, err := h.service.GetStats(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.StatsResponse{
		Urls:  int32(urls),
		Users: int32(users),
	}, nil
}

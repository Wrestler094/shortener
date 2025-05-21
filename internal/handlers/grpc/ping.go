package grpchandlers

import (
	"context"

	"github.com/Wrestler094/shortener/internal/grpc/pb"
	"github.com/Wrestler094/shortener/internal/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingHandler struct {
	pb.UnimplementedPingServiceServer
	service *services.PingService
}

func NewPingHandler(s *services.PingService) *PingHandler {
	return &PingHandler{service: s}
}

func (h *PingHandler) Ping(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := h.service.Ping(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

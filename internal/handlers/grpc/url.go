package grpchandlers

import (
	"context"
	"fmt"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/grpc/pb"
	"github.com/Wrestler094/shortener/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// URLHandler реализует gRPC-интерфейс pb.URLServiceServer
type URLHandler struct {
	pb.UnimplementedURLServiceServer

	service *services.URLService
}

// NewURLHandler создаёт новый gRPC-хендлер URLService
func NewURLHandler(s *services.URLService) *URLHandler {
	return &URLHandler{service: s}
}

// ShortenURL сохраняет оригинальный URL и возвращает сокращённый
func (h *URLHandler) ShortenURL(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	shortID, err := h.service.SaveURL(ctx, req.GetUrl(), req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save URL: %v", err)
	}

	return &pb.ShortenResponse{
		ShortUrl: fmt.Sprintf("%s/%s", configs.FlagBaseAddr, shortID),
	}, nil
}

// GetOriginalURL возвращает оригинальный URL по короткому идентификатору
func (h *URLHandler) GetOriginalURL(ctx context.Context, req *pb.URLRequest) (*pb.URLResponse, error) {
	original, isDeleted, found := h.service.GetURLByID(ctx, req.GetShortUrl())
	if !found {
		return nil, status.Error(codes.NotFound, "short URL not found")
	}

	return &pb.URLResponse{
		OriginalUrl: original,
		IsDeleted:   isDeleted,
	}, nil
}

// DeleteUserURLs удаляет список URL пользователя
func (h *URLHandler) DeleteUserURLs(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := h.service.DeleteUserURLs(ctx, req.GetUserId(), req.GetShortUrls())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete URLs: %v", err)
	}

	return &pb.DeleteResponse{Success: true}, nil
}

// GetUserURLs возвращает список URL пользователя
func (h *URLHandler) GetUserURLs(ctx context.Context, req *pb.UserRequest) (*pb.UserURLsResponse, error) {
	urls, err := h.service.GetUserURLs(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user URLs: %v", err)
	}

	resp := &pb.UserURLsResponse{}
	for _, u := range urls {
		resp.Urls = append(resp.Urls, &pb.UserURLItem{
			ShortUrl:    u.ShortURL,
			OriginalUrl: u.OriginalURL,
		})
	}
	return resp, nil
}

// SaveBatch сохраняет список URL
func (h *URLHandler) SaveBatch(ctx context.Context, req *pb.BatchRequest) (*pb.BatchResponse, error) {
	var input []dto.BatchRequestItem
	for _, item := range req.GetUrls() {
		input = append(input, dto.BatchRequestItem{
			CorrelationID: item.CorrelationId,
			OriginalURL:   item.OriginalUrl,
		})
	}

	results, err := h.service.SaveBatch(ctx, input, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save batch: %v", err)
	}

	var output []*pb.BatchResponseItem
	for _, r := range results {
		output = append(output, &pb.BatchResponseItem{
			CorrelationId: r.CorrelationID,
			ShortUrl:      r.ShortURL,
		})
	}

	return &pb.BatchResponse{Urls: output}, nil
}

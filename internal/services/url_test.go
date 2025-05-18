package services

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Wrestler094/shortener/internal/deleter"
	"github.com/Wrestler094/shortener/internal/dto"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/storage/memory"
)

func newTestService() *URLService {
	store := memory.NewMemoryStorage(make(map[string]string))
	fileStorage := persistence.NewFileStorage("")
	urlDeleter := deleter.NewURLDeleter(store, time.Second)
	return NewURLService(store, fileStorage, urlDeleter)
}

func TestURLService_SaveURL(t *testing.T) {
	service := newTestService()

	tests := []struct {
		name    string
		url     string
		userID  string
		wantErr bool
	}{
		{
			name:    "valid URL",
			url:     "https://example.com",
			userID:  "user1",
			wantErr: false,
		},
		{
			name:    "invalid URL",
			url:     "example.com",
			userID:  "user1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortID, err := service.SaveURL(context.Background(), tt.url, tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, shortID)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, shortID)
			}
		})
	}
}

func TestURLService_SaveBatch(t *testing.T) {
	service := newTestService()

	batch := []dto.BatchRequestItem{
		{CorrelationID: "1", OriginalURL: "https://example1.com"},
		{CorrelationID: "2", OriginalURL: "https://example2.com"},
		{CorrelationID: "3", OriginalURL: "https://example3.com"},
	}

	result, err := service.SaveBatch(context.Background(), batch, "user1")
	require.NoError(t, err)
	require.Len(t, result, 3)

	for i, item := range result {
		assert.Equal(t, batch[i].CorrelationID, item.CorrelationID)
		assert.Contains(t, item.ShortURL, "/")
	}
}

func TestURLService_GetURLByID(t *testing.T) {
	service := newTestService()

	// Сначала сохраняем URL
	shortID, err := service.SaveURL(context.Background(), "https://example.com", "user1")
	require.NoError(t, err)

	// Получаем URL
	url, isDeleted, ok := service.GetURLByID(context.Background(), shortID)
	assert.True(t, ok)
	assert.False(t, isDeleted)
	assert.Equal(t, "https://example.com", url)

	// Проверяем несуществующий URL
	url, isDeleted, ok = service.GetURLByID(context.Background(), "nonexistent")
	assert.False(t, ok)
	assert.False(t, isDeleted)
	assert.Empty(t, url)
}

package services_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Wrestler094/shortener/internal/deleter"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage/memory"
)

func BenchmarkURLService_SaveURL_Memory(b *testing.B) {
	store := memory.NewMemoryStorage(make(map[string]string))
	fileStorage := persistence.NewFileStorage("")
	urlDeleter := deleter.NewURLDeleter(store, time.Second)
	service := services.NewURLService(store, fileStorage, urlDeleter)

	for i := 0; i < b.N; i++ {
		original := fmt.Sprintf("https://site.com/%d", i)
		_, _ = service.SaveURL(original, "user1")
	}
}

func BenchmarkURLService_GetOriginalURL_Memory(b *testing.B) {
	store := memory.NewMemoryStorage(make(map[string]string))
	fileStorage := persistence.NewFileStorage("")
	urlDeleter := deleter.NewURLDeleter(store, time.Second)
	service := services.NewURLService(store, fileStorage, urlDeleter)
	_ = store.Save("abc", "https://yandex.ru", "user1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = service.GetURLByID("abc")
	}
}

func BenchmarkURLService_DeleteUserURLs_Memory(b *testing.B) {
	store := memory.NewMemoryStorage(make(map[string]string))
	fileStorage := persistence.NewFileStorage("")
	urlDeleter := deleter.NewURLDeleter(store, time.Second)
	service := services.NewURLService(store, fileStorage, urlDeleter)

	ids := make([]string, 1000)
	for i := range ids {
		short := fmt.Sprintf("short%d", i)
		_ = store.Save(short, "https://site.com", "user1")
		ids[i] = short
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.DeleteUserURLs("user1", ids)
	}
}

package httphandlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Wrestler094/shortener/internal/configs"
	"github.com/Wrestler094/shortener/internal/deleter"
	"github.com/Wrestler094/shortener/internal/persistence"
	"github.com/Wrestler094/shortener/internal/services"
	"github.com/Wrestler094/shortener/internal/storage/memory"
)

func newTestHandler() *URLHandler {
	fileStorage := persistence.NewFileStorage("")
	recoveredUrls := fileStorage.RecoverURLs()
	store := memory.NewMemoryStorage(recoveredUrls)
	urlDeleter := deleter.NewURLDeleter(store, time.Second)

	service := services.NewURLService(store, fileStorage, urlDeleter)
	return NewURLHandler(service)
}

func TestSavePlainURL(t *testing.T) {
	handler := newTestHandler()

	tests := []struct {
		name       string
		body       string
		wantCode   int
		wantPrefix string
	}{
		{
			name:       "valid plain URL",
			body:       "http://yandex.ru",
			wantCode:   http.StatusCreated,
			wantPrefix: configs.FlagBaseAddr + "/",
		},
		{
			name:       "invalid plain URL",
			body:       "yandex.ru",
			wantCode:   http.StatusBadRequest,
			wantPrefix: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			handler.SavePlainURL(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			respBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			assert.Equal(t, tt.wantCode, res.StatusCode)
			if tt.wantPrefix != "" {
				assert.True(t, strings.HasPrefix(string(respBody), tt.wantPrefix))
			}
		})
	}
}

func TestSaveJSONURL(t *testing.T) {
	handler := newTestHandler()

	tests := []struct {
		name       string
		body       string
		wantCode   int
		wantField  string
		wantPrefix string
	}{
		{
			name:       "valid JSON URL",
			body:       `{"url": "http://yandex.ru"}`,
			wantCode:   http.StatusCreated,
			wantField:  `"result":"`,
			wantPrefix: configs.FlagBaseAddr + "/",
		},
		{
			name:       "invalid JSON URL",
			body:       `{"url": "yandex.ru"}`,
			wantCode:   http.StatusBadRequest,
			wantField:  "",
			wantPrefix: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.SaveJSONURL(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			respBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			assert.Equal(t, tt.wantCode, res.StatusCode)

			if tt.wantField != "" {
				assert.Contains(t, string(respBody), tt.wantField)
				assert.Contains(t, string(respBody), tt.wantPrefix)
			}
		})
	}
}

func TestGetURL(t *testing.T) {
	handler := newTestHandler()

	t.Run("invalid method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/bad", nil)
		rec := httptest.NewRecorder()

		handler.GetURL(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

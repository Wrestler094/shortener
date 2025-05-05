// Middleware для управления аутентификацией пользователей через cookie.
// Основные функции:
// - Автоматическая генерация и установка cookie для новых пользователей
// - Проверка и валидация существующих cookie
// - Безопасное хранение идентификатора пользователя в cookie с использованием HMAC-подписи
// - Передача идентификатора пользователя через контекст запроса
package middlewares

import (
	"context"
	"net/http"

	"github.com/Wrestler094/shortener/internal/utils"
)

// contextKey представляет тип ключа для контекста
type contextKey string

// userIDContextKey - ключ для хранения идентификатора пользователя в контексте
const userIDContextKey = contextKey("user_id")

// GetUserIDFromContext извлекает идентификатор пользователя из контекста
// ctx - контекст запроса
// Возвращает:
// - идентификатор пользователя
// - флаг наличия идентификатора в контексте
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDContextKey).(string)
	return id, ok
}

// AuthCookieSetter - middleware для установки и проверки cookie аутентификации
// Если cookie отсутствует или невалидна, создает новую
// Добавляет идентификатор пользователя в контекст запроса
func AuthCookieSetter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID string

		c, err := r.Cookie(utils.CookieName)
		if err == nil {
			if id, valid := utils.ValidateSignedValue(c.Value); valid {
				userID = id
			}
		}

		if userID == "" {
			userID = utils.GenerateUserID()
			signed := utils.CreateSignedValue(userID)
			http.SetCookie(w, &http.Cookie{
				Name:     utils.CookieName,
				Value:    signed,
				Path:     "/",
				MaxAge:   utils.CookieMaxAge,
				HttpOnly: true,
			})
		}

		ctx := context.WithValue(r.Context(), userIDContextKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

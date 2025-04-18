package middlewares

import (
	"context"
	"net/http"

	"github.com/Wrestler094/shortener/internal/utils"
)

type contextKey string

const userIDContextKey = contextKey("user_id")
const cookieMaxAge = 3600 * 24 * 365

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDContextKey).(string)
	return id, ok
}

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
				MaxAge:   cookieMaxAge,
				HttpOnly: true,
			})
		}

		ctx := context.WithValue(r.Context(), userIDContextKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

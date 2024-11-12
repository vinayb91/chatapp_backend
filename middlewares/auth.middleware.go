package middlewares

import (
	"context"
	"net/http"

	"github.com/vinayb91/chatapp_backend/utils"
)

type contextKey string

const TokenContextKey = contextKey("token")

func ProtectRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.VerifyJWT(utils.GetCookie(r))
		if err != nil {
			http.Error(w, "Unauthorized - Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), TokenContextKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

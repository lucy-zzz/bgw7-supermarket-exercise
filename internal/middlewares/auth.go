package middlewares

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "token12345" {
			response.Error(w, http.StatusUnauthorized, "Not authorized.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

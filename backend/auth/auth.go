package auth

import (
	"context"
	"log/slog"
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/utils"
)

func TokenAuthMiddlewareFunc(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return TokenAuthMiddlewareHandler(http.HandlerFunc(handler))
}

func TokenAuthMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-KEY")
		println("Token: ", token)
		if !isValidToken(r.Context(), token) {
			utils.WriteHttpError(w, "You are not authorized to perform this action.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isValidToken(ctx context.Context, token string) bool {
	tokenRows, err := database.GetDB().QueryContext(context.Background(), "SELECT * FROM api_keys WHERE key = $1", token)
	if err != nil {
		slog.ErrorContext(ctx, "Error while validating token: ", err)
		return false
	}

	// If the token is not found, it is invalid
	if !tokenRows.Next() {
		println("Token not found: ", token, " ", tokenRows.Err())
		return false
	}

	return true
}

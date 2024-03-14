package auth

import (
	"context"
	"log/slog"
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/utils"
)

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-KEY")
		if !validToken(token) {
			utils.WriteHttpError(w, "You are not authorized to perform this action.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validToken(token string) bool {
	tokenRows, err := database.GetDB().QueryContext(context.Background(), "SELECT * FROM api_keys WHERE key = $1", token)
	if err != nil {
		slog.Error("Error while validating token: ", err)
		return false
	}

	// If the token is not found, it is invalid
	if !tokenRows.Next() {
		return false
	}

	return true
}

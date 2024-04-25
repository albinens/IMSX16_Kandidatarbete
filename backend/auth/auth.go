package auth

import (
	"context"
	"log/slog"
	"net/http"

	"example.com/m/v2/database"
	"example.com/m/v2/utils"
)

func CreateGatewayUser(ctx context.Context, username, password string) error {
	_, err := database.GetDB().ExecContext(ctx, "INSERT INTO gateway_users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		slog.ErrorContext(ctx, "Error while creating gateway user: ", err)
		return err
	}

	return nil
}

func DeleteGatewayUser(ctx context.Context, username string) error {
	_, err := database.GetDB().ExecContext(ctx, "DELETE FROM gateway_users WHERE username = $1", username)
	if err != nil {
		slog.ErrorContext(ctx, "Error while deleting gateway user: ", err)
		return err
	}

	return nil
}

func CreateApiKey(ctx context.Context, key string) error {
	_, err := database.GetDB().ExecContext(ctx, "INSERT INTO api_keys (key) VALUES ($1)", key)
	if err != nil {
		slog.ErrorContext(ctx, "Error while creating API key: ", err)
		return err
	}

	return nil
}

func DeleteApiKey(ctx context.Context, key string) error {
	_, err := database.GetDB().ExecContext(ctx, "DELETE FROM api_keys WHERE key = $1", key)
	if err != nil {
		slog.ErrorContext(ctx, "Error while deleting API key: ", err)
		return err
	}

	return nil
}

func VerifyGatewayUserMiddlewareFunc(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return VerifyGatewayUserMiddlewareHandler(http.HandlerFunc(handler))
}

func VerifyGatewayUserMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			utils.WriteHttpError(w, "You must provide a username and password.", http.StatusUnauthorized)
			return
		}

		if !verifyGatewayUser(r.Context(), username, password) {
			utils.WriteHttpError(w, "Invalid username or password.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ApiKeys() (keys []string, err error) {
	rows, err := database.GetDB().Query("SELECT * FROM api_keys")
	if err != nil {
		slog.Error("Error while getting API keys: ", err)
		return nil, err
	}

	for rows.Next() {
		var key string
		err = rows.Scan(&key)
		if err != nil {
			slog.Error("Error while getting API keys: ", err)
			return nil, err
		}

		keys = append(keys, key)
	}

	return keys, nil
}

func verifyGatewayUser(ctx context.Context, username, password string) bool {
	rows, err := database.GetDB().QueryxContext(ctx, "SELECT password FROM gateway_users WHERE username = $1", username)
	if err != nil {
		slog.ErrorContext(ctx, "Error while verifying gateway user: ", err)
		return false
	}

	if !rows.Next() {
		return false
	}

	var storedPassword string
	err = rows.Scan(&storedPassword)
	if err != nil {
		slog.ErrorContext(ctx, "Error while verifying gateway user: ", err)
		return false
	}

	if storedPassword != password {
		return false
	}

	return true
}

func TokenAuthMiddlewareFunc(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return TokenAuthMiddlewareHandler(http.HandlerFunc(handler))
}

func TokenAuthMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-KEY")
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
		return false
	}

	return true
}

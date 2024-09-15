package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/K-Kizuku/eisa-auth/pkg/jwt"
)

type Key string

const UserIDKey Key = "userID"

func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		id, err := jwt.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid or token", http.StatusUnauthorized)
			slog.Error("Invalid or expired ID token", "error", err.Error())
			return
		}

		ctx = context.WithValue(ctx, UserIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

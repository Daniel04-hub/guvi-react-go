package middleware

import (
	"context"
	"guvi-project/db"
	"net/http"
	"strings"
)

type key int

const UserKey key = 0

// Auth middleware validates the Bearer token against Redis
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS pre-flight handling handled by main CORS, but if checks fail here common method
		if r.Method == "OPTIONS" {
			next(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		
		// Check Redis
		email, err := db.RedisClient.Get(context.Background(), token).Result()
		if err != nil {
			http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Token is valid, email is in value. Pass it to context.
		ctx := context.WithValue(r.Context(), UserKey, email)
		next(w, r.WithContext(ctx))
	}
}

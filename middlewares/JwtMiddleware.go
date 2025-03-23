package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type contextKey string

const UserIDKey contextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("JWT_SECRET")

		cookie, err := r.Cookie("set-token")
		if err != nil {
			http.Error(w, "missing token cookie", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		userIDSTR, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "invalid token payload", http.StatusUnauthorized)
			return
		}
		userID, err := uuid.Parse(userIDSTR)
		if err != nil {
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) (uuid.UUID, bool) {
	UserID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	return UserID, ok
}

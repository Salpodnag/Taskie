package websockets

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var hub = NewHub()

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
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
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "invalid token payload", http.StatusUnauthorized)
		return
	}
	userID := int(userIDFloat)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка WebSocket:", err)
		return
	}

	hub.RegisterClient(userID, conn)
	defer hub.UnregisterClient(userID)

}

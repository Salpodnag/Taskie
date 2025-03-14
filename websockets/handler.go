package websockets

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(w http.ResponseWriter, r *http.Request, hub *Hub) {
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

		return
	}

	hub.RegisterClient(userID, conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	hub.UnregisterClient(userID)
}

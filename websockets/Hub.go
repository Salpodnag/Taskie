package websockets

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Hub struct {
	clients map[int]*websocket.Conn
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int]*websocket.Conn),
	}
}

func (hub *Hub) RegisterClient(UserID int, conn *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[UserID] = conn
}

func (hub *Hub) UnregisterClient(userID int) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	if conn, ok := hub.clients[userID]; ok {
		conn.Close()
		delete(hub.clients, userID)
	}
}

func (hub *Hub) Broadcast(message []byte) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	for userID, conn := range hub.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Ошибка отправки WebSocket-сообщения клиенту %d: %v", userID, err)
			conn.Close()
			delete(hub.clients, userID)
		}
	}
}

func (hub *Hub) SendToUser(userID int, message []byte) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	if conn, ok := hub.clients[userID]; ok {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Ошибка отправки сообщения пользователю %d: %v", userID, err)
			conn.Close()
			delete(hub.clients, userID)
		}
	}
}

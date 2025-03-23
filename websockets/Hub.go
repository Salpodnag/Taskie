package websockets

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[uuid.UUID]*websocket.Conn
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[uuid.UUID]*websocket.Conn),
	}
}

func (hub *Hub) RegisterClient(userID uuid.UUID, conn *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[userID] = conn
}

func (hub *Hub) UnregisterClient(userID uuid.UUID) {
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

		go func(c *websocket.Conn, id uuid.UUID) {
			if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Ошибка отправки пользователю %d: %v", id, err)
				hub.UnregisterClient(id)
			}
		}(conn, userID)
	}
}

func (hub *Hub) SendToUser(userID uuid.UUID, message []byte) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	if conn, ok := hub.clients[userID]; ok {
		go func() {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println("Ошибка отправки:", err)
				hub.UnregisterClient(userID)
			}
		}()
	}
}

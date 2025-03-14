package websockets

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
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

func (hub *Hub) RegisterClient(userID int, conn *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[userID] = conn
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

		go func(c *websocket.Conn, id int) {
			if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Ошибка отправки пользователю %d: %v", id, err)
				hub.UnregisterClient(id)
			}
		}(conn, userID)
	}
}

func (hub *Hub) SendToUser(userID int, message []byte) {
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

package websockets

import (
	"encoding/json"
	"fmt"
)

type WebSocketService struct {
	hub *Hub
}

func NewWebSocketService(hub *Hub) *WebSocketService {
	return &WebSocketService{hub: hub}
}

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (ws *WebSocketService) SendMessageBroadcast(messageType string, data interface{}) error {

	message := Message{
		Type: messageType,
		Data: data,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	ws.hub.Broadcast(messageJSON)
	return nil
}

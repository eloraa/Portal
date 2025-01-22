package server

import (
	"encoding/json"
	"log"
	"portal/internal/models"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	clients    map[*websocket.Conn]*models.Client
	rooms      map[string]*models.Room
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	broadcast  chan []byte
	mu         sync.RWMutex
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:    make(map[*websocket.Conn]*models.Client),
		rooms:      make(map[string]*models.Room),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		broadcast:  make(chan []byte),
	}
}

// generateRoomID creates a unique room identifier using UUID
func generateRoomID() string {
	return uuid.New().String()
}

func (wss *WebSocketServer) HandleConnection(conn *websocket.Conn) {
	client := &models.Client{
		Conn:    conn,
		RoomIDs: make([]string, 0),
	}

	wss.mu.Lock()
	wss.clients[conn] = client
	wss.mu.Unlock()

	defer func() {
		wss.mu.Lock()
		delete(wss.clients, conn)
		wss.mu.Unlock()
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msg models.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("error unmarshaling message: %v", err)
			continue
		}

		switch msg.Type {
		case "join_room":
			wss.handleJoinRoom(client, msg)
		case "create_room":
			wss.handleCreateRoom(client, msg)
		case "leave_room":
			wss.handleLeaveRoom(client, msg)
		case "signal":
			wss.handleSignal(client, msg)
		}
	}
}

func (wss *WebSocketServer) handleJoinRoom(client *models.Client, msg models.Message) {
	roomID := msg.RoomID

	wss.mu.Lock()
	defer wss.mu.Unlock()

	room, exists := wss.rooms[roomID]
	if !exists {
		// Send error message back to client
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Room does not exist",
		})
		return
	}

	if !room.IsPublic {
		password, ok := msg.Payload.(string)
		if !ok || password != room.Password {
			client.Conn.WriteJSON(models.Message{
				Type:    "error",
				Payload: "Invalid password",
			})
			return
		}
	}

	room.Members[client.Conn] = client.Username
	client.RoomIDs = append(client.RoomIDs, roomID)

	// Notify other members about the new user
	for conn := range room.Members {
		if conn != client.Conn {
			conn.WriteJSON(models.Message{
				Type:    "user_joined",
				RoomID:  roomID,
				UserID:  client.UserID,
				Payload: client.Username,
			})
		}
	}
}

func (wss *WebSocketServer) handleCreateRoom(client *models.Client, msg models.Message) {
	type createRoomPayload struct {
		IsPublic bool   `json:"isPublic"`
		Password string `json:"password,omitempty"`
	}

	payload, ok := msg.Payload.(map[string]interface{})
	if !ok {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Invalid payload",
		})
		return
	}

	roomID := generateRoomID() // Implement this helper function
	room := &models.Room{
		ID:       roomID,
		Creator:  client.UserID,
		IsPublic: payload["isPublic"].(bool),
		Password: payload["password"].(string),
		Members:  make(map[*websocket.Conn]string),
	}

	wss.mu.Lock()
	wss.rooms[roomID] = room
	wss.mu.Unlock()

	client.Conn.WriteJSON(models.Message{
		Type:    "room_created",
		RoomID:  roomID,
		Payload: room,
	})
}

func (wss *WebSocketServer) handleLeaveRoom(client *models.Client, msg models.Message) {
	roomID := msg.RoomID

	wss.mu.Lock()
	defer wss.mu.Unlock()

	room, exists := wss.rooms[roomID]
	if !exists {
		return
	}

	delete(room.Members, client.Conn)

	// Remove room from client's room list
	for i, id := range client.RoomIDs {
		if id == roomID {
			client.RoomIDs = append(client.RoomIDs[:i], client.RoomIDs[i+1:]...)
			break
		}
	}

	// Notify other members
	for conn := range room.Members {
		conn.WriteJSON(models.Message{
			Type:    "user_left",
			RoomID:  roomID,
			UserID:  client.UserID,
			Payload: client.Username,
		})
	}

	// If room is empty, delete it
	if len(room.Members) == 0 {
		delete(wss.rooms, roomID)
	}
}

func (wss *WebSocketServer) handleSignal(client *models.Client, msg models.Message) {
	roomID := msg.RoomID

	wss.mu.RLock()
	room, exists := wss.rooms[roomID]
	if !exists {
		wss.mu.RUnlock()
		return
	}

	// Forward the signal to all other members in the room
	for conn := range room.Members {
		if conn != client.Conn {
			conn.WriteJSON(models.Message{
				Type:    "signal",
				RoomID:  roomID,
				UserID:  client.UserID,
				Payload: msg.Payload,
			})
		}
	}
	wss.mu.RUnlock()
}

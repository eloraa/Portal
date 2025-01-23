package server

import (
	"encoding/json"
	"log"
	"portal/internal/models"
	"portal/internal/utils"
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
	// Check for existing connection with same userID
	for existingConn, existingClient := range wss.clients {
		if existingClient.UserID == client.UserID {
			// Close existing connection
			existingConn.Close()
			// Remove from all rooms
			wss.removeClientFromAllRooms(existingClient)
			delete(wss.clients, existingConn)
			break
		}
	}
	wss.clients[conn] = client
	wss.mu.Unlock()

	defer func() {
		wss.mu.Lock()
		// Remove client from all rooms and notify others
		wss.removeClientFromAllRooms(client)
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

func (wss *WebSocketServer) RoomExists(roomID string) bool {
	wss.mu.RLock()
	defer wss.mu.RUnlock()
	_, exists := wss.rooms[roomID]
	return exists
}

func (wss *WebSocketServer) handleJoinRoom(client *models.Client, msg models.Message) {
	roomID := msg.RoomID

	if !utils.ValidateRoomID(roomID) {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Invalid room ID format",
		})
		return
	}

	if msg.UserID == "" {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "UserID is required",
		})
		return
	}

	type joinPayload struct {
		Password string `json:"password,omitempty"`
		Username string `json:"username"`
		AvatarID string `json:"avatarId"`
	}

	payload, ok := msg.Payload.(map[string]interface{})
	if !ok {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Invalid payload",
		})
		return
	}

	// Validate required fields
	username, ok := payload["username"].(string)
	if !ok || username == "" {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Username is required",
		})
		return
	}

	avatarID, ok := payload["avatarId"].(string)
	if !ok || avatarID == "" {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "AvatarID is required",
		})
		return
	}

	// Validate avatar ID
	validAvatars := map[string]bool{
		"kazuha":  true,
		"diluc":   true,
		"ganyu":   true,
		"hutao":   true,
		"shotgun": true,
		"shenhe":  true,
	}

	if !validAvatars[avatarID] {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Invalid avatar ID",
		})
		return
	}

	// Update client info
	client.Username = username
	client.AvatarID = avatarID
	client.UserID = msg.UserID

	wss.mu.Lock()
	defer wss.mu.Unlock()

	room, exists := wss.rooms[roomID]
	if !exists {
		client.Conn.WriteJSON(models.Message{
			Type:   "room_not_found",
			RoomID: roomID,
			Payload: map[string]interface{}{
				"suggestedName": utils.GenerateRoomName(),
				"createRoom":    true,
			},
		})
		return
	}

	if !room.IsPublic {
		password, _ := payload["password"].(string)
		if password != room.Password {
			client.Conn.WriteJSON(models.Message{
				Type:    "error",
				Payload: "Invalid password",
			})
			return
		}
	}

	// Get current room members before adding the new user
	members := make([]map[string]string, 0)
	for conn := range room.Members {
		if memberClient, exists := wss.clients[conn]; exists {
			members = append(members, map[string]string{
				"userId":   memberClient.UserID,
				"username": memberClient.Username,
				"avatarId": memberClient.AvatarID,
			})
		}
	}

	// Add new member
	room.Members[client.Conn] = client.Username
	client.RoomIDs = append(client.RoomIDs, roomID)

	// Add the new member to the members list
	members = append(members, map[string]string{
		"userId":   client.UserID,
		"username": client.Username,
		"avatarId": client.AvatarID,
	})

	// Send current members to the new user
	client.Conn.WriteJSON(models.Message{
		Type:   "room_joined",
		RoomID: roomID,
		Payload: map[string]interface{}{
			"members":  members,
			"name":     room.Name,
			"isPublic": room.IsPublic,
		},
	})

	newMemberInfo := map[string]string{
		"userId":   client.UserID,
		"username": client.Username,
		"avatarId": client.AvatarID,
	}

	// Notify other members about the new user
	for conn := range room.Members {
		if conn != client.Conn {
			conn.WriteJSON(models.Message{
				Type:   "user_joined",
				RoomID: roomID,
				UserID: client.UserID,
				Payload: map[string]interface{}{
					"user": newMemberInfo,
					"name": room.Name,
				},
			})
		}
	}
}

func (wss *WebSocketServer) handleCreateRoom(client *models.Client, msg models.Message) {
	payload, ok := msg.Payload.(map[string]interface{})
	if !ok {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Invalid payload",
		})
		return
	}

	roomID := msg.RoomID // Use provided roomID if exists
	if roomID != "" {
		// Validate the provided room ID
		if !utils.ValidateRoomID(roomID) {
			client.Conn.WriteJSON(models.Message{
				Type:    "error",
				Payload: "Invalid room ID format. Must be 8 characters long and contain only letters and numbers",
			})
			return
		}
	} else {
		roomID = utils.GenerateShortID()
	}

	// Check if room ID already exists
	if wss.RoomExists(roomID) {
		client.Conn.WriteJSON(models.Message{
			Type:    "error",
			Payload: "Room ID already exists",
		})
		return
	}

	roomName, _ := payload["name"].(string)
	if roomName == "" {
		roomName = utils.GenerateRoomName()
	}

	isPublic, _ := payload["isPublic"].(bool)
	password, _ := payload["password"].(string)

	room := &models.Room{
		ID:       roomID,
		Name:     roomName,
		Creator:  client.UserID,
		IsPublic: isPublic,
		Password: password,
		Members:  make(map[*websocket.Conn]string),
	}

	wss.mu.Lock()
	wss.rooms[roomID] = room
	wss.mu.Unlock()

	client.Conn.WriteJSON(models.Message{
		Type:   "room_created",
		RoomID: roomID,
		Payload: map[string]interface{}{
			"name":     roomName,
			"isPublic": isPublic,
			"creator":  client.UserID,
		},
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

	// Notify other members with full user info
	leavingMemberInfo := map[string]string{
		"userId":   client.UserID,
		"username": client.Username,
		"avatarId": client.AvatarID,
	}

	for conn := range room.Members {
		conn.WriteJSON(models.Message{
			Type:    "user_left",
			RoomID:  roomID,
			UserID:  client.UserID,
			Payload: leavingMemberInfo,
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

// Add this method to clean up empty rooms periodically
func (wss *WebSocketServer) cleanupEmptyRooms() {
	wss.mu.Lock()
	defer wss.mu.Unlock()

	for roomID, room := range wss.rooms {
		if len(room.Members) == 0 {
			delete(wss.rooms, roomID)
		}
	}
}

// Add this new helper function
func (wss *WebSocketServer) removeClientFromAllRooms(client *models.Client) {
	for _, roomID := range client.RoomIDs {
		room, exists := wss.rooms[roomID]
		if !exists {
			continue
		}

		delete(room.Members, client.Conn)

		// Notify other members about the disconnection
		leavingMemberInfo := map[string]string{
			"userId":   client.UserID,
			"username": client.Username,
			"avatarId": client.AvatarID,
		}

		for conn := range room.Members {
			conn.WriteJSON(models.Message{
				Type:    "user_left",
				RoomID:  roomID,
				UserID:  client.UserID,
				Payload: leavingMemberInfo,
			})
		}

		// If room is empty, delete it
		if len(room.Members) == 0 {
			delete(wss.rooms, roomID)
		}
	}
}

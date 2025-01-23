package server

import (
	"net/http"
	"portal/internal/models"
	"portal/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (s *Server) SetupRoutes() {
	// API routes
	api := s.Router.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("", s.CreateUser)
		}

		// Room routes
		rooms := api.Group("/rooms")
		{
			rooms.GET("", s.GetRooms)
			rooms.GET("/:id", s.GetRoom)
			rooms.POST("", s.CreateRoom)
			rooms.POST("/:id", s.UpdateRoom)
		}
	}

	// WebSocket route
	s.Router.GET("/ws", s.handleWebSocket)

	// Health check
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}

// Optional room handlers
func (s *Server) GetRooms(c *gin.Context) {
	type roomInfo struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Members int    `json:"members"`
	}

	rooms := make([]roomInfo, 0)
	for _, room := range s.ws.rooms {
		if room.IsPublic {
			rooms = append(rooms, roomInfo{
				ID:      room.ID,
				Name:    room.Name,
				Members: len(room.Members),
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}

func (s *Server) GetRoom(c *gin.Context) {
	roomID := c.Param("id")
	room, exists := s.ws.rooms[roomID]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error":         "Room not found",
			"createRoom":    true,
			"suggestedName": utils.GenerateRoomName(),
			"roomId":        roomID,
		})
		return
	}

	if !room.IsPublic {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Room is private",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      room.ID,
		"name":    room.Name,
		"creator": room.Creator,
		"members": len(room.Members),
	})
}

func (s *Server) CreateRoom(c *gin.Context) {
	type createRoomRequest struct {
		IsPublic bool   `json:"isPublic"`
		Password string `json:"password,omitempty"`
		UserID   string `json:"userId" binding:"required"`
		Name     string `json:"name,omitempty"` // Optional, will generate if not provided
	}

	var req createRoomRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	roomName := req.Name
	if roomName == "" {
		roomName = utils.GenerateRoomName()
	}

	roomID := utils.GenerateShortID()
	room := &models.Room{
		ID:       roomID,
		Name:     roomName,
		Creator:  req.UserID,
		IsPublic: req.IsPublic,
		Password: req.Password,
		Members:  make(map[*websocket.Conn]string),
	}

	s.ws.mu.Lock()
	s.ws.rooms[roomID] = room
	s.ws.mu.Unlock()

	c.JSON(http.StatusCreated, gin.H{
		"roomId":   roomID,
		"name":     roomName,
		"isPublic": req.IsPublic,
		"creator":  req.UserID,
	})
}

func (s *Server) UpdateRoom(c *gin.Context) {
	roomID := c.Param("id")

	type updateRoomRequest struct {
		UserID   string `json:"userId" binding:"required"`
		Name     string `json:"name,omitempty"`
		IsPublic *bool  `json:"isPublic,omitempty"`
		Password string `json:"password,omitempty"`
	}

	var req updateRoomRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	s.ws.mu.Lock()
	defer s.ws.mu.Unlock()

	room, exists := s.ws.rooms[roomID]
	if !exists {
		// Room doesn't exist, ask for creation details
		c.JSON(http.StatusNotFound, gin.H{
			"error":         "Room not found",
			"createRoom":    true,
			"suggestedName": utils.GenerateRoomName(),
			"roomId":        roomID,
		})
		return
	}

	// Check if the user is the creator
	if room.Creator != req.UserID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Only the room creator can update the room",
		})
		return
	}

	// Update room properties
	if req.Name != "" {
		room.Name = req.Name
	}
	if req.IsPublic != nil {
		room.IsPublic = *req.IsPublic
	}
	if req.Password != "" {
		room.Password = req.Password
	}

	c.JSON(http.StatusOK, gin.H{
		"roomId":   room.ID,
		"name":     room.Name,
		"isPublic": room.IsPublic,
		"creator":  room.Creator,
	})
}

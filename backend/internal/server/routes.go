package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	rooms := make([]string, 0)
	for roomID, room := range s.ws.rooms {
		if room.IsPublic {
			rooms = append(rooms, roomID)
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
			"error": "Room not found",
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
		"creator": room.Creator,
		"members": len(room.Members),
	})
}

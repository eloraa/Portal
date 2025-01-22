package server

import (
	"net/http"
	"portal/internal/config"
	"portal/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	config   *config.Config
	Router   *gin.Engine
	ws       *WebSocketServer
	db       *db.Database
	upgrader websocket.Upgrader
}

func NewServer(cfg *config.Config, database *db.Database) *Server {
	server := &Server{
		config: cfg,
		Router: gin.Default(),
		ws:     NewWebSocketServer(),
		db:     database,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	// Configure CORS
	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Add NoRoute handler for 404 responses
	server.Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"path":  c.Request.URL.Path,
		})
	})

	server.SetupRoutes()
	return server
}

func (s *Server) CreateUser(c *gin.Context) {
	userID, err := s.db.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"userId": userID,
	})
}

func (s *Server) handleWebSocket(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not upgrade connection",
		})
		return
	}

	s.ws.HandleConnection(conn)
}

func (s *Server) Start() error {
	return s.Router.Run(s.config.ServerAddress)
}

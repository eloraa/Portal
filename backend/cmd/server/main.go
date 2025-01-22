package main

import (
	"log"
	"os"
	"portal/internal/config"
	"portal/internal/db"
	"portal/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	database, err := db.NewDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Initialize tables
	if err := database.InitializeTables(); err != nil {
		log.Fatal("Error initializing tables:", err)
	}

	// Create config
	cfg := &config.Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}
	if cfg.ServerAddress == "" {
		cfg.ServerAddress = ":8080"
	}

	// Initialize server
	s := server.NewServer(cfg, database)

	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := s.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/Go-lang-Blog-Apps/config"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/app"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// defer database.CloseDB(db)

	// Initialize Echo server
	server := app.NewServer(cfg, db)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("🚀 Starting server on %s\n", addr)

	if err := server.Start(addr); err != nil {
		log.Fatal(err)
	}
}

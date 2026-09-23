package database

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/Go-lang-Blog-Apps/config"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database connection
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w", err)
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.Comment{},
		&models.Post{},
	)
	if err != nil {
		log.Fatal("som error to auto migrate")
	}
	log.Println("Successfully connected to postgreSQL")
	return db, nil
}

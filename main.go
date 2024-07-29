package main

import (
	"log"

	"github.com/EmiliodDev/GinGuard/internal/config"
	"github.com/EmiliodDev/GinGuard/internal/models"
	"github.com/EmiliodDev/GinGuard/pkg/database"
)


func main() {
    config.LoadConfig()

    db, err := database.ConnectToDatabase()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    db.AutoMigrate(&models.User{}, &models.Role{})
}

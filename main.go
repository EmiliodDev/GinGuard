package main

import (
	"fmt"
	"log"

	"github.com/EmiliodDev/GinGuard/internal/config"
	"github.com/EmiliodDev/GinGuard/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
    config.LoadConfig()

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.AppConfig.Database.User,
        config.AppConfig.Database.Password,
        config.AppConfig.Database.Host,
        config.AppConfig.Database.Port,
        config.AppConfig.Database.Name,
    )
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    db.AutoMigrate(&models.User{}, &models.Role{})
}

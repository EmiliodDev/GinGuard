package database

import (
	"fmt"

	"github.com/EmiliodDev/GinGuard/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.AppConfig.Database.User,
        config.AppConfig.Database.Password,
        config.AppConfig.Database.Host,
        config.AppConfig.Database.Port,
        config.AppConfig.Database.Name,
    )
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

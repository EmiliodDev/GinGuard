package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username        string  `gorm:"uniqueIndex;not null"`
    Email           string  `gorm:"uniqueIndex;not null"`
    PasswordHash    string
    RoleID          uint
    Role            Role
}

type Role struct {
    gorm.Model
    Name            string  `gorm:"uniqueindex;not null"`
    Description     string
    Users           []User
}



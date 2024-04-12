package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Name     string
	Email    string
	Password string
}

type UserResponse struct {
    ID    uuid.UUID `json:"id"`
    Name  string    `json:"name"`
    Email string    `json:"email"`
}

type AddUser struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UpdateUser struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

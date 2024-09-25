package model

import (
	"errors"
	"time"
)

// User represents a user in the system.
type User struct {
	ID        int       `json:"id" gorm:"primary_key" example:"1"`                                              // ID of the user
	Username  string    `json:"username" gorm:"type:varchar(255);not null" example:"山田太郎"`                      // Username of the user
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null" example:"sample@gmail.com"` // Email of the user
	Password  string    `json:"password" gorm:"type:varchar(255)" example:"password"`                           // Password of the user
	CreatedAt time.Time `json:"created_at" example:"2024-09-25T11:46:43Z"`                                      // Creation timestamp
	Foods     []Food    `gorm:"foreignKey:UserID"`                                                              // Foods associated with the user
}

// UserResponse represents a response containing user information.
type UserResponse struct {
	ID        int       `json:"id" example:"1"`                            // ID of the user
	Username  string    `json:"username" example:"山田太郎"`                   // Username of the user
	Email     string    `json:"email" example:"sample@gmail.com"`          // Email of the user
	CreatedAt time.Time `json:"created_at" example:"2024-09-25T11:46:43Z"` // Creation timestamp
}

// UserRequest represents the request structure for creating or updating a user.
type UserRequest struct {
	Username string `json:"username" example:"山田太郎"`          // Username of the user
	Email    string `json:"email" example:"sample@gmail.com"` // Email of the user
	Password string `json:"password" example:"password"`      // Password of the user
}

var ErrInvalidPassword = errors.New("invalid password")

package model

import (
	"time"
)

// User represents a user in the system.
type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	Foods     []Food    `gorm:"foreignKey:UserID"`
}

// UserResponse represents a response containing user information.
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

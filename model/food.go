package model

import (
	"time"
)

type Food struct {
	ID             int       `json:"id" gorm:"primary_key"`
	Name           string    `json:"name" gorm:"not null"`
	UserID         int       `json:"user_id" gorm:"not null"`
	OriginalCode   int       `json:"original_code"`
	Quantity       int       `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	ExpirationDate time.Time `json:"expiration_date"`
	ImageURL       string    `json:"image_url"`
	Memo           string    `json:"memo"`
	User           User      `gorm:"foreignKey:UserID"`
}

type FoodResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	UserID         int       `json:"user_id"`
	OriginalCode   int       `json:"original_code"`
	Quantity       int       `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	ExpirationDate time.Time `json:"expiration_date"`
	ImageURL       string    `json:"image_url"`
	Memo           string    `json:"memo"`
}

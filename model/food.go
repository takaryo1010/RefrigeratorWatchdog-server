package model

import (
	"time"
)

// Food represents a food item in the database.
type Food struct {
	ID             int       `json:"id" gorm:"primary_key" example:"1"` // ID of the food item
	Name           string    `json:"name" gorm:"not null" example:"オレンジ"` // Name of the food item
	UserID         int       `json:"user_id" gorm:"not null" example:"1"` // User ID associated with the food item
	OriginalCode   int       `json:"original_code" example:"12456456"` // Original code of the food item
	Quantity       float64       `json:"quantity" example:"5.5"` // Quantity of the food item
	CreatedAt      time.Time `json:"created_at" example:"2024-09-25T11:46:43Z"` // Creation timestamp
	ExpirationDate *time.Time `json:"expiration_date" example:"2024-12-15T00:00:00Z"` // Expiration date
	ImageURL       string    `json:"image_url" example:"images/orange.jpg"` // URL of the food item image
	Memo           string    `json:"memo" example:"新鮮なオレンジだったものです"` // Additional notes or memo
	Tag            string    `json:"tag" gorm:"type:enum('野菜', '肉', '魚', '乳製品','調味料','卵','飲料','果物','加工食品','その他');default:'その他'"` // Tag of the food item
	User           User      `gorm:"foreignKey:UserID"` // User associated with the food item
}

// FoodResponse represents the response structure for a food item.
type FoodResponse struct {
	ID             int       `json:"id" example:"1"` // ID of the food item
	Name           string    `json:"name" example:"オレンジ"` // Name of the food item
	UserID         int       `json:"user_id" example:"1"` // User ID associated with the food item
	OriginalCode   int       `json:"original_code" example:"12456456"` // Original code of the food item
	Quantity       float64       `json:"quantity" example:"5.5"` // Quantity of the food item
	CreatedAt      time.Time `json:"created_at" example:"2024-09-25T11:46:43Z"` // Creation timestamp
	ExpirationDate *time.Time `json:"expiration_date" example:"2024-12-15T00:00:00Z"` // Expiration date
	ImageURL       string    `json:"image_url" example:"images/orange.jpg"` // URL of the food item image
	Tag 		  string    `json:"tag" example:"果物"` // Tag of the food item
	Memo           string    `json:"memo" example:"新鮮なオレンジだったものです"` // Additional notes or memo
}

// FoodRequest represents the request structure for creating a new food item.
type FoodRequest struct {
	Name           string    `json:"name" example:"オレンジ"` // Name of the food item
	UserID         int       `json:"user_id" example:"1"` // User ID associated with the food item
	OriginalCode   int       `json:"original_code" example:"12456456"` // Original code of the food item
	Quantity       float64       `json:"quantity" example:"5.5"` // Quantity of the food item
	ExpirationDate *time.Time `json:"expiration_date" example:"2024-12-15T00:00:00Z"` // Expiration date
	ImageURL       string    `json:"image_url" example:"images/orange.jpg"` // URL of the food item image
	Tag 		  string    `json:"tag" example:"果物"` // Tag of the food item'野菜', '肉', '魚', '乳製品','調味料','卵','飲料','果物','加工食品','その他'
	Memo           string    `json:"memo" example:"新鮮なオレンジだったものです"` // Additional notes or memo
}

package models

import (
	"time"
)

// Product represents a product in the store
type Product struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

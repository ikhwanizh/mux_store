package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint       `json:"user_id"` // Foreign key for User
	User      User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"foreignKey:CartID"` // Correct foreign key definition
}

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`    // Foreign key for Cart
	ProductID uint    `json:"product_id"` // Foreign key for Product
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
}

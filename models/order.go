package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
}

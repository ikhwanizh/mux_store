package models

// Cart represents a user's shopping cart
type Cart struct {
	ID     int        `json:"id" gorm:"primaryKey"`
	UserID int        `json:"user_id"`
	Items  []CartItem `json:"items,omitempty" gorm:"foreignKey:CartID"`
}

// CartItem represents an item in a user's shopping cart
type CartItem struct {
	ID        int `json:"id" gorm:"primaryKey"`
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

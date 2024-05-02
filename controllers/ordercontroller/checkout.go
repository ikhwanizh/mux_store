package ordercontroller

import (
	"net/http"
	"online-store-backend/helper"
	"online-store-backend/models"
)

var responseJson = helper.ResponseJson

func Checkout(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uint)
	if !ok || userID == 0 {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Retrieve the user's cart
	var cart models.Cart
	if err := models.DB.Preload("User").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		http.Error(w, "Failed to retrieve user cart: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if cart.ID == 0 {
		http.Error(w, "No cart found for the user", http.StatusNotFound)
		return
	}

	// Retrieve all cart items for the found cart
	var cartItems []models.CartItem
	if err := models.DB.Preload("Product").Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		http.Error(w, "Failed to load cart items: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate the total price
	totalPrice := calculateTotalPrice(cartItems)

	// Create the order
	order := models.Order{
		UserID:     cart.UserID, // Ensure this uses the user ID from the cart, not context directly
		TotalPrice: totalPrice,
	}
	if err := models.DB.Create(&order).Error; err != nil {
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear the cart
	if err := models.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		http.Error(w, "Failed to clear cart: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created order details
	responseJson(w, order, http.StatusCreated)
}

func calculateTotalPrice(cartItems []models.CartItem) float64 {
	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.Product.Price * float64(item.Quantity)
	}
	return totalPrice
}

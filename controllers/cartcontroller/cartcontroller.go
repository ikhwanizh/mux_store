package cartcontroller

import (
	"encoding/json"
	"net/http"
	"online-store-backend/helper"
	"online-store-backend/models"
	"strconv"

	"github.com/gorilla/mux"
)

var responseJson = helper.ResponseJson
var responseError = helper.ResponseError

// AddToCart adds a new item to the user's shopping cart
func AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uint) // Assume you have a middleware to extract UserID and put it in context
	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve or create a cart for the user
	var cart models.Cart
	if err := models.DB.Where(models.Cart{UserID: userID}).FirstOrCreate(&cart).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Assign the correct CartID to the CartItem
	cartItem.CartID = cart.ID
	if err := models.DB.Create(&cartItem).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, cartItem, http.StatusCreated)
}

func ViewCart(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("userID").(uint)
	var cartItems []models.CartItem
	if err := models.DB.Preload("Product").Where("cart_id = (SELECT id FROM carts WHERE user_id = ?)", userID).Find(&cartItems).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, cartItems, http.StatusOK)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemID := vars["cartItemID"]

	// Convert cartItemID from string to int
	id, err := strconv.Atoi(cartItemID)
	if err != nil {
		responseError(w, "Invalid cart item ID", http.StatusBadRequest)
		return
	}

	// Attempt to delete the cart item using a safe WHERE condition
	result := models.DB.Where("id = ?", id).Delete(&models.CartItem{})
	if result.Error != nil {
		responseError(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		responseError(w, "Cart item not found", http.StatusNotFound)
		return
	}

	responseJson(w, "Cart item deleted successfully", http.StatusOK)
}

// package ordercontroller

// import (
// 	"errors"
// 	"net/http"
// 	"online-store-backend/helper"
// 	"online-store-backend/models"
// 	"time"

// 	"gorm.io/gorm"
// )

// var responseJson = helper.ResponseJson
// var responseError = helper.ResponseError
// var getUserID = helper.GetUserIDFromContext

// func Checkout(w http.ResponseWriter, r *http.Request) {
// 	userID, ok := getUserID(r.Context())
// 	if !ok {
// 		responseError(w, "User not authenticated", http.StatusBadRequest)
// 		return
// 	}

// 	cartID, err := getCartIDForUser(userID)
// 	if err != nil {
// 		responseError(w, "Failed to retrieve cart", http.StatusInternalServerError)
// 		return
// 	}

// 	var cartItems []models.CartItem
// 	if err := models.DB.Where("cart_id = ?", cartID).Preload("Product").Find(&cartItems).Error; err != nil {
// 		responseError(w, "Failed to get cart items", http.StatusInternalServerError)
// 		return
// 	}

// 	totalPrice := 0.0
// 	for _, item := range cartItems {
// 		totalPrice += float64(item.Quantity) * item.Product.Price
// 	}

// 	order := models.Order{
// 		UserID:     userID,
// 		TotalPrice: totalPrice,
// 		CreatedAt:  time.Now(),
// 	}

// 	tx := models.DB.Begin()
// 	if err := tx.Create(&order).Error; err != nil {
// 		tx.Rollback()
// 		responseError(w, "Failed to create order", http.StatusInternalServerError)
// 		return
// 	}

// 	if err := tx.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error; err != nil {
// 		tx.Rollback()
// 		responseError(w, "Failed to clear cart", http.StatusInternalServerError)
// 		return
// 	}

// 	tx.Commit()
// 	responseJson(w, order, http.StatusOK)
// }

// func getCartIDForUser(userID int) (int, error) {
// 	var cart models.Cart
// 	result := models.DB.First(&cart, "user_id = ?", userID)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		cart.UserID = userID
// 		if err := models.DB.Create(&cart).Error; err != nil {
// 			return 0, err
// 		}
// 	} else if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return cart.ID, nil
// }

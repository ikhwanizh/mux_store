package productcontroller

import (
	"encoding/json"
	"net/http"
	"online-store-backend/helper"
	"online-store-backend/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var responseJson = helper.ResponseJson
var responseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseJson(w, products, http.StatusOK)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Product not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	responseJson(w, product, http.StatusOK)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Create the product in the database
	if err := models.DB.Create(&product).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Reload the product with the category preloaded to include category details in the response
	var loadedProduct models.Product
	if err := models.DB.Preload("Category").First(&loadedProduct, product.ID).Error; err != nil {
		responseError(w, "Failed to load product with category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, loadedProduct, http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Product not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Save(&product).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, product, http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Product not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := models.DB.Delete(&product).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, product, http.StatusOK)

}

// Filter products by category
func FilterByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Product not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var products []models.Product
	if err := models.DB.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, products, http.StatusOK)
}

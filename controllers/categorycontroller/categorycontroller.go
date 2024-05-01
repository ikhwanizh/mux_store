package categorycontroller

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
	var categories []models.Category
	if err := models.DB.Find(&categories).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseJson(w, categories, http.StatusOK)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category models.Category
	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Category not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	responseJson(w, category, http.StatusOK)
}

func Create(w http.ResponseWriter, r *http.Request) {

	var category models.Category

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&category).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, category, http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category models.Category
	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Category not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Save(&category).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, category, http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category models.Category
	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Category not found", http.StatusNotFound)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := models.DB.Delete(&category).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, category, http.StatusOK)
}

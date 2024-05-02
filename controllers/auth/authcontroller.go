package auth

import (
	"encoding/json"
	"net/http"
	"online-store-backend/config"
	"online-store-backend/helper"
	"online-store-backend/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var responseJson = helper.ResponseJson
var responseError = helper.ResponseError

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := models.DB.Create(&userInput).Error; err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJson(w, userInput, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var loginInput models.User
	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginInput); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Where("username = ?", loginInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, "Invalid username or password", http.StatusUnauthorized)
			return
		default:
			responseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
		responseError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	//JWT
	expTime := time.Now().Add(1 * time.Minute)
	claims := &config.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		responseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		Expires:  expTime,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login successful"}
	responseJson(w, response, http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	response := map[string]string{"message": "Logout successful"}
	responseJson(w, response, http.StatusOK)
}

package main

import (
	"fmt"
	"net/http"
	"online-store-backend/controllers/auth"
	"online-store-backend/controllers/categorycontroller"
	"online-store-backend/controllers/productcontroller"
	"online-store-backend/middlewares"
	"online-store-backend/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDB()
	r := mux.NewRouter()

	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	api.HandleFunc("/products", productcontroller.Create).Methods("POST")
	api.HandleFunc("/products/{id}", productcontroller.Update).Methods("PUT")
	api.HandleFunc("/products/{id}", productcontroller.Delete).Methods("DELETE")

	api.HandleFunc("/categories", categorycontroller.Index).Methods("GET")
	api.HandleFunc("/categories/{id}", categorycontroller.Show).Methods("GET")
	api.HandleFunc("/categories", categorycontroller.Create).Methods("POST")
	api.HandleFunc("/categories/{id}", categorycontroller.Update).Methods("PUT")
	api.HandleFunc("/categories/{id}", categorycontroller.Delete).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

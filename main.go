package main

import (
	"net/http"
	"online-store-backend/controllers/productcontroller"
	"online-store-backend/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDB()
	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/products", productcontroller.Create).Methods("POST")
	r.HandleFunc("/products/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", productcontroller.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

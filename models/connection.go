package models

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the PostgreSQL database
func ConnectDB() {
	db, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=mux_store port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Category{}, &Product{}, &Cart{}, &CartItem{}, &Order{})
	DB = db
}

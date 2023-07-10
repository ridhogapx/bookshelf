package config

import (
	"bookshelf/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	ConnectionDB()
	ConnectRPC()
}

func ConnectionDB() {
	var err error
	dsn := "host=localhost user=root password=root dbname=bookshelf port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Model Initiator
	DB.AutoMigrate(model.Book{})
	DB.AutoMigrate(model.Owner{})
}

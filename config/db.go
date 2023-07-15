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
	/* Should be changed to DB Container IP Address! */
	dsn := "host=172.17.0.2 user=root password=root dbname=bookshelf port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Model Initiator
	DB.AutoMigrate(model.Book{})
	DB.AutoMigrate(model.Owner{})
}

func ExConnect() {
	/*
	* This is example of connection
	* Just to test. Should not used in production!
	 */

	var err error
	dsn := "host=localhost user=root password=root dbname=bookshelf port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	DB.AutoMigrate(model.Book{})
	DB.AutoMigrate(model.Owner{})

}

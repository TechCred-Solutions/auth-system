package db

import (
	"auth-system/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable to hold the database connection

// InitDB initializes the database connection and performs auto-migration
func InitDB() {
	// Open a connection to the SQLite database
	database, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Automatically migrate the User model to keep the schema up to date
	database.AutoMigrate(&models.User{})
	
	// Assign the database connection to the global variable
	DB = database
}

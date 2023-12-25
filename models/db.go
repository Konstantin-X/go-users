package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbFile = "users.db"

var DB *gorm.DB

func Setup() {
	DB, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Printf("--- connected to DB\n")

	// Perform database migration
	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

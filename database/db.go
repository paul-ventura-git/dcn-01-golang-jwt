package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Error conectando a la base de datos:", err)
	}

	DB = db
	log.Println("✅ Base de datos conectada.")
}

package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		dsn = "host=db user=postgres password=postgres dbname=library port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanina bağlanirken hata oluştu:", err)
	}

	fmt.Println("✅ Veritabanina başariyla bağlanildi!")
	DB = db
}

package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// dbConfig holds the configuration for the database connection.
type dbConfig struct {
	user     string
	password string
	name     string
	port     string
}

func (c dbConfig) dsn() string {
	return fmt.Sprintf("host=db user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.user, c.password, c.name, c.port)
}

func loadDBConfig() (dbConfig, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		return dbConfig{}, fmt.Errorf("DB_USER environment variable is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return dbConfig{}, fmt.Errorf("DB_PASSWORD environment variable is not set")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return dbConfig{}, fmt.Errorf("DB_NAME environment variable is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return dbConfig{}, fmt.Errorf("PORT environment variable is not set")
	}

	return dbConfig{
		user:     user,
		password: password,
		name:     name,
		port:     port,
	}, nil
}

func Connect() {
	config, err := loadDBConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	db, err := gorm.Open(postgres.Open(config.dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanina bağlanirken hata oluştu:", err)
	}

	fmt.Println("✅ Veritabanina başariyla bağlanildi!")
	DB = db
}

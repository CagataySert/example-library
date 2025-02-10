package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/CagataySert/library-system/internal/models"
)

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
		return dbConfig{}, fmt.Errorf("DB_PORT environment variable is not set")
	}

	return dbConfig{
		user:     user,
		password: password,
		name:     name,
		port:     port,
	}, nil
}

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		config, err := loadDBConfig()
		if err != nil {
			log.Fatalf("Failed to load database configuration: %v", err)
		}

		db, err := gorm.Open(postgres.Open(config.dsn()), &gorm.Config{})
		if err != nil {
			log.Fatal("Veritabanina bağlanirken hata oluştu:", err)
		}

		log.Println("✅ Veritabanina başariyla bağlanildi!")
		dbInstance = db

		// Migration Process
		mErr := db.AutoMigrate(
			&models.Book{},
		)

		if mErr != nil {
			log.Println("❌ Migration failed:", mErr)
		} else {
			log.Println("✅ Database migrated successfully!")
		}
	})

	return dbInstance
}

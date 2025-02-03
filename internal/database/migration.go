package database

import (
	"fmt"

	"github.com/CagataySert/library-system/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Book{})
	if err != nil {
		fmt.Println("❌ Migration işlemi başarısız oldu:", err)
	} else {
		fmt.Println("✅ Migration başarıyla tamamlandı!")
	}
}

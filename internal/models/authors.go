package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Books []Book `json:"books"` // One author can have many books
}

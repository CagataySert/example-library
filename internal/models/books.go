package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string `json:"title"`
	AuthorID      uint   `json:"author_id"` // Foreign key to Author
	PublishedDate string `json:"published_date"`
	Author        Author `gorm:"foreignKey:AuthorID"` // Relationship to Author
}

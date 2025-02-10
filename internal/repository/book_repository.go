package repositories

import (
	"github.com/CagataySert/library-system/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetById(id int) (*models.Book, error)
	GetBooks() ([]models.Book, error)
	Create(book models.Book) error
}

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) BookRepository {
	return &bookRepository{DB: DB}
}

func (r *bookRepository) GetById(id int) (*models.Book, error) {
	var book models.Book

	err := r.DB.First(&book, id).Error

	return &book, err
}

func (r *bookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.DB.Find(&books).Error

	return books, err
}

func (r *bookRepository) Create(book models.Book) error {
	return r.DB.Create(book).Error
}

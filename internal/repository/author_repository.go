package repositories

import (
	"github.com/CagataySert/library-system/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetById(id int) (*models.Author, error)
	GetAuthors() ([]models.Author, error)
	Create(author models.Author) error
}

type authorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(DB *gorm.DB) AuthorRepository {
	return &authorRepository{DB: DB}
}

func (r *authorRepository) GetById(id int) (*models.Author, error) {
	var author models.Author

	err := r.DB.First(&author, id).Error

	return &author, err
}

func (r *authorRepository) GetAuthors() ([]models.Author, error) {
	var authors []models.Author

	err := r.DB.Find(&authors).Error

	return authors, err
}

func (r *authorRepository) Create(author models.Author) error {
	return r.DB.Create(author).Error
}

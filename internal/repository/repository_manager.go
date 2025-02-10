package repositories

import "gorm.io/gorm"

type RepositoryManager interface {
	Book() BookRepository
	Author() AuthorRepository
}

type repositoryManager struct {
	DB         *gorm.DB
	bookRepo   BookRepository
	authorRepo AuthorRepository
}

func NewRepositoryManager(DB *gorm.DB) RepositoryManager {
	return &repositoryManager{
		DB:         DB,
		bookRepo:   NewBookRepository(DB),
		authorRepo: NewAuthorRepository(DB),
	}
}

func (r *repositoryManager) Book() BookRepository {
	return r.bookRepo
}

func (r *repositoryManager) Author() AuthorRepository {
	return r.authorRepo
}

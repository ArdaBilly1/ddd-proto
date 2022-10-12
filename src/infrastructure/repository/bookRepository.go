package repository

import (
	model "ddd-proto/src/domain/model/book"
	"ddd-proto/src/domain/repository"
	"time"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepositoryContract {
	return bookRepository{db: db}
}

func (r bookRepository) Create(book model.Book) error {
	data := new(model.Book)

	data.Name = book.Name
	data.Type = book.Name
	data.CreatedAt = time.Now()

	err := r.db.Model(book).Create(&data).Error

	return err
}

func (r bookRepository) GetBookById(id int) (data model.Book, err error) {
	err = r.db.Model(data).Where("id", id).Find(&data).Error

	return
}

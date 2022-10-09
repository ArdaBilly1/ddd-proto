package repository

import "ddd-proto/src/domain/model"

type BookRepositoryContract interface {
	Create(book model.Book) error
	GetBookById(id int) (data model.Book, err error)
}

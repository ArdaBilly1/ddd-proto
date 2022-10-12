package repository

import model "ddd-proto/src/domain/model/book"

type BookRepositoryContract interface {
	Create(book model.Book) error
	GetBookById(id int) (data model.Book, err error)
}

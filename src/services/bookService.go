package services

import (
	"ddd-proto/src/domain/model/book"
	model "ddd-proto/src/domain/model/book"
	"ddd-proto/src/domain/repository"
)

type bookService struct {
	repository repository.BookRepositoryContract
}

type BookServiceContract interface {
	CreateNew(req book.RequestBookCreate) error
}

func NewBookService(repo repository.BookRepositoryContract) BookServiceContract {
	return bookService{repository: repo}
}

func (r bookService) CreateNew(req book.RequestBookCreate) error {
	data := new(model.Book)

	data.Name = req.Name
	data.Type = req.Type

	err := r.repository.Create(*data)

	return err
}

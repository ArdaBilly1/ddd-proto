package services

import (
	"ddd-proto/src/domain/model"
	"ddd-proto/src/domain/repository"
	"ddd-proto/src/interface/http/handler/v1/request"
)

type bookService struct {
	repository repository.BookRepositoryContract
}

type BookServiceContract interface {
}

func NewBookService(repo repository.BookRepositoryContract) BookServiceContract {
	return bookService{repository: repo}
}

func (r bookService) CreateNew(req request.RequestBookCreate) error {
	data := new(model.Book)

	data.Name = req.Name
	data.Type = req.Type

	err := r.repository.Create(*data)

	return err
}

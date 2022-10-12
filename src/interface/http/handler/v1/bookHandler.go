package v1

import (
	"ddd-proto/src/domain/model/book"
	utilities "ddd-proto/src/infrastructure/utilitIes"
	"ddd-proto/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	service services.BookServiceContract
}

type BookHandlerContract interface {
	Create(c echo.Context) error
}

func NewBookHandler(svc services.BookServiceContract) BookHandlerContract {
	return bookHandler{service: svc}
}

func (h bookHandler) Create(c echo.Context) error {
	req := new(book.RequestBookCreate)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return utilities.SetResponse(
			c,
			http.StatusBadRequest,
			utilities.MsgErrorFailedInsert,
			err,
		)
	}

	if err := h.service.CreateNew(*req); err != nil {
		return utilities.SetResponse(
			c,
			http.StatusBadRequest,
			utilities.MsgErrorFailedInsert,
			err,
		)
	}

	return utilities.SetResponse(
		c,
		http.StatusCreated,
		utilities.MsgSuccessCreate,
	)
}

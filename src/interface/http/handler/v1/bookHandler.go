package v1

import (
	utilities "ddd-proto/src/infrastructure/utilitIes"
	"ddd-proto/src/interface/http/handler/v1/request"
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
	req := new(request.RequestBookCreate)
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

	return utilities.SetResponse(
		c,
		http.StatusCreated,
		utilities.MsgSuccessCreate,
	)
}

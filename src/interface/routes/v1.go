package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func V1(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo says: im fine :)")
	})
}

package routes

import (
	"ddd-proto/config"
	"ddd-proto/src/infrastructure/repository"
	v1 "ddd-proto/src/interface/http/handler/v1"
	"ddd-proto/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func V1(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo says: im fine :)")
	})

	bookConnector := bookConnector(config.MysqlDB)
	book := v1.Group("/book")
	book.POST("", bookConnector.Create)
}

// Connector
func bookConnector(db *gorm.DB) v1.BookHandlerContract {
	repo := repository.NewBookRepository(db)
	svc := services.NewBookService(repo)
	return v1.NewBookHandler(svc)
}

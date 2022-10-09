package utilities

import "github.com/labstack/echo/v4"

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data; omitempty"`
}

func SetResponse(c echo.Context, code int, msg string, data ...interface{}) error {
	response := new(DefaultResponse)

	response.Message = msg
	response.Data = data

	return c.JSON(code, response)
}

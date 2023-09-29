package auth

import (
	"belajar_golang/models/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    nil,
	})
}

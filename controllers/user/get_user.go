package user

import (
	"belajar_golang/configs"
	basemodel "belajar_golang/models/base"
	"belajar_golang/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []user.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status:  false,
			Message: "Failed get from database!",
			Data:    users,
		})
	}
	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    users,
	})
}

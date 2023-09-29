package auth

import (
	"belajar_golang/configs"
	"belajar_golang/middleware"
	"belajar_golang/models/base"
	usermodel "belajar_golang/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var user usermodel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := configs.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "failed add to Database",
			Data:    nil,
		})
	}

	var authReponse = usermodel.ResponseAuth{
		user.Id,
		user.Name,
		user.Email,
		middleware.GenerateTokenJwt(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success Register",
		Data:    authReponse,
	})
}

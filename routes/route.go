package routes

import (
	authcontroller "belajar_golang/controllers/auth"
	usercontroller "belajar_golang/controllers/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", authcontroller.LoginController)
	e.POST("/register", authcontroller.RegisterController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	eAuth.GET("/users", usercontroller.GetUsersController)

}

package main

import (
	"belajar_golang/configs"
	"belajar_golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)	
	e.Start(":8000")
}





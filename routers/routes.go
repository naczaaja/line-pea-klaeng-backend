package routers

import (
	"main/handlers"

	"github.com/labstack/echo/v4"
)

func GetHello(e *echo.Echo) {
	e.GET("/", handlers.GetHelloWorld)
}

func PostClientDataRegister(e *echo.Echo) {
	e.POST("/client-data-register", handlers.PostClientDataRegister)
}
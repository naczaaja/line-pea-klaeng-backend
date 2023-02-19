package main

import (
	"main/db"
	"main/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.ConnectDb()

	e := echo.New()

	//CORS config for all domain
	e.Use(middleware.CORS())

	routers.GetHello(e)

	e.Logger.Fatal(e.Start(":1323"))
}

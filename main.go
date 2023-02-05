package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	e := echo.New()

	// Attach the logger to Echo
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.POST("/users", saveUser)

	e.Logger.Fatal(e.Start(":1323"))
}

// FORM: application/x-www-form-urlencoded
func saveUser(c echo.Context) error {
	imageAvatar := c.FormValue("imageAvatar")
	ca := c.FormValue("ca")
	pea_meter := c.FormValue("pea_meter")
	return c.String(http.StatusOK, "imageAvatar:"+imageAvatar+", ca:"+ca+", pea_meter:"+pea_meter)
}

type User struct {
	ImageAvatar string `json:"imageAvatar" xml:"imageAvatar" form:"imageAvatar" query:"imageAvatar"`
	CA          int    `json:"ca" xml:"ca" form:"ca" query:"ca"`
	Pea_numbet  int    `json:"pea_number" xml:"pea_number" form:"pea_number" query:"pea_number"`
}

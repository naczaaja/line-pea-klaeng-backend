package handlers

import (
	"main/db"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func PostClientDataRegister(c echo.Context) error {
	client_data_register := models.ClientDataRegister{}
	if err := c.Bind(&client_data_register); err != nil {
		return err
	}
	client_data_register_db := models.ClientDataRegisterDB{
		LineId:      client_data_register.LineId,
		ImageAvatar: client_data_register.ImageAvatar,
		IDcard:      client_data_register.IDcard,
	}
	result := db.DB.Db.Save(&client_data_register_db)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.RowsAffected)
	}
	return c.JSON(http.StatusOK, client_data_register_db)
}

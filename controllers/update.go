package controllers

import (
	"net/http"
	"strconv"

	"api-login/models"

	"github.com/labstack/echo"
)

func UpdateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	models.Users[id].Name = u.Name
	return c.JSON(http.StatusOK, models.Users[id])
}

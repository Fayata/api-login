package controllers

import (
	"api-login/models"
	"net/http"

	"github.com/labstack/echo"
)
func CreateUser(c echo.Context) error {
	u := &models.User{
		ID: models.Seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	models.Users[u.ID] = u
	models.Seq++
	return c.JSON(http.StatusCreated, u)
}
package controllers

import (
	"net/http"
	"strconv"

	"api-login/models"

	"github.com/labstack/echo"
)

func DeleteUser(c echo.Context) error {
	
	id, _ := strconv.Atoi(c.Param("id"))
	delete(models.Users, id)
	return c.NoContent(http.StatusNoContent)
}
package controllers

import (
	"net/http"

	"github.com/KamillKAPLAN/intern_project/server/models"
	"github.com/labstack/echo"
)

func GetCategory(c echo.Context) error {
	id := c.Param("id")

	category, err := models.GetCategory(id)
	if err != nil {
		return c.String(http.StatusNotFound, id)
	}

	return c.String(http.StatusOK, category.Name)
}

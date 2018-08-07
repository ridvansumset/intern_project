package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ridvansumset/intern_project/server/models"
)

func GetMyCategory(c echo.Context) error {
	id := c.Param("id")

	category, err := models.GetMyCategory(id)
	if err != nil {
		return c.String(http.StatusNotFound, id)
	}

	return c.String(http.StatusOK, category.Name)
}

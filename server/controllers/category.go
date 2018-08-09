package controllers

import (
	"net/http"

	"github.com/KamillKAPLAN/intern_project/server/models"
	"github.com/labstack/echo"
)

func GetCategory(c echo.Context) error {
	id := c.Param("category_id")

	category, err := models.GetCategory(id)
	if err != nil {
		return c.String(http.StatusNotFound, id+"bulunamadı.")
	}
	return c.JSON(http.StatusOK, category)
}

func ListCategory(c echo.Context) error {
	categories, err := models.ListCategory()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

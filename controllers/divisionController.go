package controllers

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDivisions(c echo.Context) error {
	divisions, err := database.GetDivisions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, divisions)
}

func GetDivision(c echo.Context) error {
	id := c.Param("id")
	division, err := database.GetDivision(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, division)
}

func CreateDivision(c echo.Context) error {
	division := models.Division{}
	c.Bind(&division)
	err := database.CreateDivision(&division)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, division)
}

func UpdateDivision(c echo.Context) error {
	division := models.Division{}
	c.Bind(&division)
	err := database.UpdateDivision(&division)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, division)
}

func DeleteDivision(c echo.Context) error {
	division := models.Division{}
	c.Bind(&division)
	err := database.DeleteDivision(&division)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, division)
}

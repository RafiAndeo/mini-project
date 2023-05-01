package controllers

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRoles(c echo.Context) error {
	roles, err := database.GetRoles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, roles)
}

func GetRole(c echo.Context) error {
	id := c.Param("id")
	role, err := database.GetRole(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, role)
}

func CreateRole(c echo.Context) error {
	role := models.Role{}
	c.Bind(&role)
	err := database.CreateRole(&role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, role)
}

func UpdateRole(c echo.Context) error {
	role := models.Role{}
	c.Bind(&role)
	err := database.UpdateRole(&role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, role)
}

func DeleteRole(c echo.Context) error {
	role := models.Role{}
	c.Bind(&role)
	err := database.DeleteRole(&role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, role)
}

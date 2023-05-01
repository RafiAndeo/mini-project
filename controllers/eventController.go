package controllers

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEvents(c echo.Context) error {
	events, err := database.GetEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, events)
}

func GetEvent(c echo.Context) error {
	id := c.Param("id")
	event, err := database.GetEvent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, event)
}

func CreateEvent(c echo.Context) error {
	event := models.Event{}
	c.Bind(&event)
	err := database.CreateEvent(&event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, event)
}

func UpdateEvent(c echo.Context) error {
	event := models.Event{}
	c.Bind(&event)
	err := database.UpdateEvent(&event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, event)
}

func DeleteEvent(c echo.Context) error {
	event := models.Event{}
	c.Bind(&event)
	err := database.DeleteEvent(&event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, event)
}

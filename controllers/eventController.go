package controllers

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetEventsController(c echo.Context) error {
	var events []models.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all events",
		"events":  events,
	})
}

func GetEventController(c echo.Context) error {
	event := models.Event{}
	EventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&event, EventId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get event by id",
		"event":   event,
	})
}

func CreateEventController(c echo.Context) error {
	event := models.Event{}
	c.Bind(&event)

	if err := config.DB.Save(&event).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new event",
		"event":   event,
	})
}

func UpdateEventController(c echo.Context) error {
	event := models.Event{}
	c.Bind(&event)
	EventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", EventId).Updates(&event).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update event",
		"event":   event,
	})
}

func DeleteEventController(c echo.Context) error {
	event := models.Event{}
	EventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", EventId).Delete(&event).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete event",
		"event":   event,
	})
}

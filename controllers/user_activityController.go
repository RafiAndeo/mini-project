package controllers

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserActivities(c echo.Context) error {
	user_activities, err := database.GetUserActivities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user_activities)
}

func GetUserActivity(c echo.Context) error {
	id := c.Param("id")
	user_activity, err := database.GetUserActivity(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user_activity)
}

func CreateUserActivity(c echo.Context) error {
	user_activity := models.UserActivity{}
	c.Bind(&user_activity)
	err := database.CreateUserActivity(&user_activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user_activity)
}

func UpdateUserActivity(c echo.Context) error {
	user_activity := models.UserActivity{}
	c.Bind(&user_activity)
	err := database.UpdateUserActivity(&user_activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user_activity)
}

func DeleteUserActivity(c echo.Context) error {
	user_activity := models.UserActivity{}
	c.Bind(&user_activity)
	err := database.DeleteUserActivity(&user_activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user_activity)
}

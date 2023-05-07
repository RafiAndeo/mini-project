package controllers

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserActivityController(c echo.Context) error {
	var userActivity []models.UserActivity

	if err := config.DB.Find(&userActivity).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get all user activity",
		"userActivity": userActivity,
	})
}

func GetUserActivityByUserController(c echo.Context) error {
	userActivity := models.UserActivity{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&userActivity, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get user activity by id",
		"userActivity": userActivity,
	})
}

func CreateUserActivityController(c echo.Context) error {
	userActivity := models.UserActivity{}
	c.Bind(&userActivity)

	if err := config.DB.Save(&userActivity).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success create new user activity",
		"userActivity": userActivity,
	})
}

func UpdateUserActivityController(c echo.Context) error {
	userActivity := models.UserActivity{}
	c.Bind(&userActivity)
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Save(&userActivity).Where("id = ?", UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update user activity",
		"userActivity": userActivity,
	})
}

func DeleteUserActivityController(c echo.Context) error {
	userActivity := models.UserActivity{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Delete(&userActivity, UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user activity",
	})
}

package controllers

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDivisionsController(c echo.Context) error {
	var divisions []models.Division

	if err := config.DB.Find(&divisions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get all divisions",
		"divisions": divisions,
	})
}

func GetDivisionController(c echo.Context) error {
	division := models.Division{}
	DivisionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&division, DivisionId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get division by id",
		"division": division,
	})
}

func CreateDivisionController(c echo.Context) error {
	division := models.Division{}
	c.Bind(&division)

	if err := config.DB.Save(&division).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new division",
		"division": division,
	})
}

func UpdateDivisionController(c echo.Context) error {
	division := models.Division{}
	c.Bind(&division)
	DivisionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", DivisionId).Updates(&division).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success update division",
		"division": division,
	})
}

func DeleteDivisionController(c echo.Context) error {
	division := models.Division{}
	DivisionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", DivisionId).Delete(&division).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success delete division",
		"division": division,
	})
}

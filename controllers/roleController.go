package controllers

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetRolesController(c echo.Context) error {
	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all roles",
		"roles":   roles,
	})
}

func GetRoleController(c echo.Context) error {
	role := models.Role{}
	RoleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&role, RoleId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get role by id",
		"role":    role,
	})
}

func CreateRoleController(c echo.Context) error {
	role := models.Role{}
	c.Bind(&role)

	if err := config.DB.Save(&role).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new role",
		"role":    role,
	})
}

func UpdateRoleController(c echo.Context) error {
	role := models.Role{}
	c.Bind(&role)
	RoleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.Where("id = ?", RoleId).Updates(&role).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update role by id",
		"role":    role,
	})
}

func DeleteRoleController(c echo.Context) error {
	role := models.Role{}
	RoleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.Delete(&role, RoleId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete role by id",
	})
}

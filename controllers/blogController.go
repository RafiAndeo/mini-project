package controllers

import (
	"mini-project/config"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

func GetBlogController(c echo.Context) error {
	blog := models.Blog{}
	BlogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&blog, BlogId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blog":    blog,
	})
}

func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)
	BlogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", BlogId).Updates(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	blog := models.Blog{}
	BlogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := config.DB.Where("id = ?", BlogId).Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

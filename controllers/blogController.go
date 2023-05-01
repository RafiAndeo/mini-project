package controllers

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBlogs(c echo.Context) error {
	blogs, err := database.GetBlogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, blogs)
}

func GetBlog(c echo.Context) error {
	id := c.Param("id")
	blog, err := database.GetBlog(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, blog)
}

func CreateBlog(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)
	err := database.CreateBlog(&blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, blog)
}

func UpdateBlog(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)
	err := database.UpdateBlog(&blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, blog)
}

func DeleteBlog(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)
	err := database.DeleteBlog(&blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, blog)
}

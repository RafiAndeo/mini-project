package controllers

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	products, err := database.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := database.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	err := database.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	err := database.UpdateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	err := database.DeleteProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

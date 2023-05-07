package controllers

import (
	"mini-project/config"
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	user := models.User{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&user, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	hashPassword, err2 := utils.HashPassword(user.Password)
	if err2 != nil {
		return err2
	}

	user.Password = hashPassword

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	user := models.User{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.First(&user, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	c.Bind(&user)
	if err1 := config.DB.Save(&user).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}

func DeleteUserController(c echo.Context) error {
	user := models.User{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := config.DB.Delete(&user, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	token, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var dbUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to find user record")
	}

	dbUser.Token = ""
	if t, ok := token.(string); ok {
		dbUser.Token = t
	}
	if err := config.DB.Model(&dbUser).Update("token", dbUser.Token).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to save token to database")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"token":   token,
	})
}

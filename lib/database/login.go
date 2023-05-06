package database

import (
	"mini-project/config"
	"mini-project/middlewares"
	"mini-project/models"
	"mini-project/utils"
)

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	var userDB models.User
	if err = config.DB.First(&userDB, "email = ?", user.Email).Error; err != nil {
		return nil, err
	}

	if err = utils.ComparePassword(userDB.Password, user.Password); err != nil {
		return nil, err
	}

	Token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		return nil, err
	}

	return Token, nil
}

package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetUserActivities() ([]models.UserActivity, error) {
	var user_activities []models.UserActivity
	if err := config.DB.Find(&user_activities).Error; err != nil {
		return user_activities, err
	}
	return user_activities, nil
}

func GetUserActivity(id string) (models.UserActivity, error) {
	var user_activity models.UserActivity
	if err := config.DB.Where("id = ?", id).First(&user_activity).Error; err != nil {
		return user_activity, err
	}
	return user_activity, nil
}

func CreateUserActivity(user_activity *models.UserActivity) error {
	if err := config.DB.Create(&user_activity).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserActivity(user_activity *models.UserActivity) error {
	if err := config.DB.Save(&user_activity).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserActivity(user_activity *models.UserActivity) error {
	if err := config.DB.Delete(&user_activity).Error; err != nil {
		return err
	}
	return nil
}

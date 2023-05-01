package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}

func GetRole(id string) (models.Role, error) {
	var role models.Role
	if err := config.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func CreateRole(role *models.Role) error {
	if err := config.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRole(role *models.Role) error {
	if err := config.DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRole(role *models.Role) error {
	if err := config.DB.Delete(&role).Error; err != nil {
		return err
	}
	return nil
}

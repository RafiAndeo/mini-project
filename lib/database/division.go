package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetDivisions() ([]models.Division, error) {
	var divisions []models.Division
	if err := config.DB.Find(&divisions).Error; err != nil {
		return divisions, err
	}
	return divisions, nil
}

func GetDivision(id string) (models.Division, error) {
	var division models.Division
	if err := config.DB.Where("id = ?", id).First(&division).Error; err != nil {
		return division, err
	}
	return division, nil
}

func CreateDivision(division *models.Division) error {
	if err := config.DB.Create(&division).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDivision(division *models.Division) error {
	if err := config.DB.Save(&division).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDivision(division *models.Division) error {
	if err := config.DB.Delete(&division).Error; err != nil {
		return err
	}
	return nil
}

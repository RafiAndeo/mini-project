package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetEvents() ([]models.Event, error) {
	var events []models.Event
	if err := config.DB.Find(&events).Error; err != nil {
		return events, err
	}
	return events, nil
}

func GetEvent(id string) (models.Event, error) {
	var event models.Event
	if err := config.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}
	return event, nil
}

func CreateEvent(event *models.Event) error {
	if err := config.DB.Create(&event).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEvent(event *models.Event) error {
	if err := config.DB.Save(&event).Error; err != nil {
		return err
	}
	return nil
}

func DeleteEvent(event *models.Event) error {
	if err := config.DB.Delete(&event).Error; err != nil {
		return err
	}
	return nil
}

package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func GetProduct(id string) (models.Product, error) {
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func CreateProduct(product *models.Product) error {
	if err := config.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *models.Product) error {
	if err := config.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *models.Product) error {
	if err := config.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

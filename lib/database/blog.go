package database

import (
	"mini-project/config"
	"mini-project/models"
)

func GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := config.DB.Find(&blogs).Error; err != nil {
		return blogs, err
	}
	return blogs, nil
}

func GetBlog(id string) (models.Blog, error) {
	var blog models.Blog
	if err := config.DB.Where("id = ?", id).First(&blog).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func CreateBlog(blog *models.Blog) error {
	if err := config.DB.Create(&blog).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBlog(blog *models.Blog) error {
	if err := config.DB.Save(&blog).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBlog(blog *models.Blog) error {
	if err := config.DB.Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}

package config

import (
	"fmt"
	"mini-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "pass",
		DB_Port:     "3306",
		DB_Host:     "mysql",
		DB_Name:     "ase_database",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Blog{})
	DB.AutoMigrate(&models.Division{})
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserActivity{})
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_ID    uint       `json:"user_id" form:"user_id" gorm:"primary_key"`
	Roles      []Role     `json:"roles" form:"roles" gorm:"foreignKey:User_ID"`
	Divisions  []Division `json:"divisions" form:"divisions" gorm:"foreignKey:User_ID"`
	Email      string     `json:"email" form:"email" gorm:"not null"`
	Name       string     `json:"name" form:"name" gorm:"not null"`
	Photo_url  string     `json:"photo_url" form:"photo_url"`
	Created_at time.Time  `json:"created_at" form:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at time.Time  `json:"updated_at" form:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Deleted_at time.Time  `json:"deleted_at" form:"deleted_at" gorm:"default:NULL"`
}

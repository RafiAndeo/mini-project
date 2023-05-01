package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role_ID    uint      `json:"role_id" form:"role_id" gorm:"primary_key"`
	User_ID    uint      `json:"user_id" form:"user_id" gorm:"foreignkey:User_ID"`
	User       User      `gorm:"references:User_ID"`
	Name       string    `json:"name" form:"name" gorm:"not null"`
	Created_at time.Time `json:"created_at" form:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at time.Time `json:"updated_at" form:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Deleted_at time.Time `json:"deleted_at" form:"deleted_at" gorm:"default:NULL"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Division struct {
	gorm.Model
	Division_ID   uint      `json:"division_id" form:"division_id" gorm:"primary_key"`
	User_ID       uint      `json:"user_id" form:"user_id" gorm:"foreignkey:User_ID"`
	User          User      `gorm:"references:User_ID"`
	Name          string    `json:"name" form:"name" gorm:"not null"`
	Thumbnail_url string    `json:"thumbnail_url" form:"thumbnail_url"`
	Created_at    time.Time `json:"created_at" form:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at    time.Time `json:"updated_at" form:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Deleted_at    time.Time `json:"deleted_at" form:"deleted_at" gorm:"default:NULL"`
}

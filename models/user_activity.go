package models

import (
	"gorm.io/gorm"
)

type UserActivity struct {
	gorm.Model
	Name         string   `json:"name" form:"name"`
	Tags         []string `json:"tags" form:"tags" gorm:"type:enum('Event','Blog','Product');default:'Event'"`
	ThumbnailURL string   `json:"thumbnail_url" form:"thumbnail_url"`
	Description  string   `json:"description" form:"description"`
	Status       []string `json:"status" form:"status" gorm:"type:enum('Started','Postponed', 'Canceled', 'Finished');default:'Started'"`
	UserID       uint     `json:"-" form:"-" gorm:"unique;not null"`
}

package models

import (
	"gorm.io/gorm"
)

type Division struct {
	gorm.Model
	Name         []string `json:"name" form:"name" gorm:"type:enum('Pembina','Asisten','Member');default:'member';not null"`
	ThumbnailURL string   `json:"thumbnail_url" form:"thumbnail_url"`
	UserID       uint     `json:"-" form:"-" gorm:"unique;not null"`
}

package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title         string   `json:"title" form:"title"`
	ThumbnailURL  string   `json:"thumbnail_url" form:"thumbnail_url"`
	Tags          []string `json:"tags" form:"tags" gorm:"type:enum('Competition','Recruitment','Others');default:'Competition'"`
	Description   string   `json:"description" form:"description"`
	StartDatetime string   `json:"start_datetime" form:"start_datetime"`
	EndDatetime   string   `json:"end_datetime" form:"end_datetime"`
	CreatedBy     string   `json:"created_by" form:"created_by"`
	UpdatedBy     string   `json:"updated_by" form:"updated_by"`
}

package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title             string   `json:"title" form:"title"`
	Tags              []string `json:"tags" form:"tags" gorm:"type:enum('Technology','UI/UX Design','Microservice', 'Others');default:'Technology'"`
	ThumbnailURL      string   `json:"thumbnail_url" form:"thumbnail_url"`
	Description       string   `json:"description" form:"description"`
	BlogDate          string   `json:"blog_date" form:"blog_date"`
	EstimatedReadTime string   `json:"estimated_read_time" form:"estimated_read_time"`
	CreatedBy         string   `json:"created_by" form:"created_by"`
	UpdatedBy         string   `json:"updated_by" form:"updated_by"`
}

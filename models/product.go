package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string   `json:"name" form:"name"`
	Description  string   `json:"description" form:"description"`
	ThumbnailURL string   `json:"thumbnail_url" form:"thumbnail_url"`
	Link         string   `json:"link" form:"link"`
	Tags         []string `json:"tags" form:"tags" gorm:"type:enum('Game','Web','Android', 'UI/UX');default:'Game'"`
	PublishedBy  string   `json:"published_by" form:"published_by"`
	PublishedAt  string   `json:"published_at" form:"published_at"`
	CreatedBy    string   `json:"created_by" form:"created_by"`
	UpdatedBy    string   `json:"updated_by" form:"updated_by"`
}

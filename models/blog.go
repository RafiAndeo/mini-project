package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Blog_ID             uint      `json:"blog_id" form:"blog_id" gorm:"primary_key"`
	Title               string    `json:"title" form:"title" gorm:"not null"`
	Tags                []string  `json:"tags" form:"tags" gorm:"type:VARCHAR(255)[]"`
	Thumbnail_url       string    `json:"thumbnail_url" form:"thumbnail_url"`
	Description         string    `json:"description" form:"description"`
	Blog_date           string    `json:"blog_date" form:"blog_date"`
	Estimated_read_time string    `json:"estimated_read_time" form:"estimated_read_time"`
	Created_by          string    `json:"created_by" form:"created_by"`
	Updated_by          string    `json:"updated_by" form:"updated_by" gorm:"default:NULL"`
	Created_at          time.Time `json:"created_at" form:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at          time.Time `json:"updated_at" form:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Deleted_at          time.Time `json:"deleted_at" form:"deleted_at" gorm:"default:NULL"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Product_ID    uint      `json:"product_id" form:"product_id" gorm:"primary_key"`
	Name          string    `json:"name" form:"name" gorm:"not null"`
	Description   string    `json:"description" form:"description"`
	Thumbnail_url string    `json:"thumbnail_url" form:"thumbnail_url"`
	Link          string    `json:"link" form:"link"`
	Tags          []string  `json:"tags" form:"tags" gorm:"type:VARCHAR(255)[]"`
	Published_by  string    `json:"published_by" form:"published_by"`
	Published_at  string    `json:"published_at" form:"published_at" gorm:"autoCreateTime"`
	Created_by    string    `json:"created_by" form:"created_by"`
	Updated_by    string    `json:"updated_by" form:"updated_by" gorm:"default:NULL"`
	Created_at    time.Time `json:"created_at" form:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at    time.Time `json:"updated_at" form:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Deleted_at    time.Time `json:"deleted_at" form:"deleted_at" gorm:"default:NULL"`
}

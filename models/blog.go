package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagBType string

const (
	Technology   tagBType = "Technology"
	UIUXDesign   tagBType = "UI/UX Design"
	Microservice tagBType = "Microservice"
)

func (t *tagBType) Scan(value interface{}) error {
	*t = tagBType(value.([]byte))
	return nil
}

func (t tagBType) Value() (driver.Value, error) {
	return string(t), nil
}

type Blog struct {
	gorm.Model
	Title             string   `json:"title" form:"title"`
	Tags              tagBType `json:"tags" form:"tags" gorm:"type:enum('Technology', 'UI/UX Design', 'Microservice');default:'Technology'"`
	ThumbnailURL      string   `json:"thumbnail_url" form:"thumbnail_url"`
	Description       string   `json:"description" form:"description"`
	BlogDate          string   `json:"blog_date" form:"blog_date"`
	EstimatedReadTime string   `json:"estimated_read_time" form:"estimated_read_time"`
	CreatedBy         string   `json:"created_by" form:"created_by"`
	UpdatedBy         string   `json:"updated_by" form:"updated_by"`
	UserActivityID    uint     `json:"user_activity_id" form:"user_activity_id"`
}

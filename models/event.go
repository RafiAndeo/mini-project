package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagEType string

const (
	Competition tagEType = "Competition"
	Recruitment tagEType = "Recruitment"
)

func (t *tagEType) Scan(value interface{}) error {
	*t = tagEType(value.([]byte))
	return nil
}

func (t tagEType) Value() (driver.Value, error) {
	return string(t), nil
}

type Event struct {
	gorm.Model
	Title          string   `json:"title" form:"title"`
	ThumbnailURL   string   `json:"thumbnail_url" form:"thumbnail_url"`
	Tags           tagEType `json:"tags" form:"tags" gorm:"type:enum('Competition', 'Recruitment')"`
	Description    string   `json:"description" form:"description"`
	StartDatetime  string   `json:"start_datetime" form:"start_datetime"`
	EndDatetime    string   `json:"end_datetime" form:"end_datetime"`
	CreatedBy      string   `json:"created_by" form:"created_by"`
	UpdatedBy      string   `json:"updated_by" form:"updated_by"`
	UserActivityID uint     `json:"-" form:"-" gorm:"uniqueIndex:idx_event_user_activity"`
}

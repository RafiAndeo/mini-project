package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagDType string

const (
	Pembina tagDType = "Pembina"
	Asisten tagDType = "Asisten"
	Member  tagDType = "Member"
)

func (t *tagDType) Scan(value interface{}) error {
	*t = tagDType(value.([]byte))
	return nil
}

func (t tagDType) Value() (driver.Value, error) {
	return string(t), nil
}

type Division struct {
	gorm.Model
	Name         tagDType `json:"name" form:"name" gorm:"type:enum('Pembina', 'Asisten', 'Member')"`
	ThumbnailURL string   `json:"thumbnail_url" form:"thumbnail_url"`
	UserID       uint     `json:"-" form:"-"`
}

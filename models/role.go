package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagRType string

const (
	Pembina tagRType = "Pembina"
	Asisten tagRType = "Asisten"
	Member  tagRType = "Member"
)

func (t *tagRType) Scan(value interface{}) error {
	*t = tagRType(value.([]byte))
	return nil
}

func (t tagRType) Value() (driver.Value, error) {
	return string(t), nil
}

type Role struct {
	gorm.Model
	Name   tagRType `json:"name" form:"name" gorm:"type:enum('Pembina', 'Asisten', 'Member');default:'Member'"`
	UserID uint     `json:"-" form:"-"`
}

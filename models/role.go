package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	UserID uint   `json:"-" form:"-"`
}

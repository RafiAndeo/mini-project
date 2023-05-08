package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `json:"name" form:"name" gorm:"type:enum('Pembina', 'Asisten', 'Member');default:'Member'"`
	UserID uint   `json:"user_id" form:"user_id"`
}

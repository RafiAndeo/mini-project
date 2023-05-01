package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string     `json:"name" form:"name"`
	Email      string     `json:"email" form:"email"`
	PhotoURL   string     `json:"photo_url" form:"photo_url"`
	RoleID     uint       `json:"-" form:"-" gorm:"unique;not null"`
	DivisionID uint       `json:"-" form:"-" gorm:"unique;not null"`
	Roles      []Role     `json:"roles" form:"roles" gorm:"not null"`
	Divisions  []Division `json:"divisions" form:"divisions" gorm:"type:enum('Pembina','Asisten','Member');default:'member'"`
}

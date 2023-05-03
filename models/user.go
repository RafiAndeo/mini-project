package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string   `json:"name" form:"name"`
	Email      string   `json:"email" form:"email"`
	Password   string   `json:"password" form:"password"`
	PhotoURL   string   `json:"photo_url" form:"photo_url"`
	RoleID     uint     `json:"-" form:"-"`
	DivisionID uint     `json:"-" form:"-"`
	Roles      []Role   `json:"roles" form:"roles"`
	Divisions  tagDType `json:"divisions" form:"divisions" gorm:"type:enum('Pembina', 'Asisten', 'Member')"`
	Token      string   `json:"token" form:"token"`
}

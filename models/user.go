package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name" form:"name"`
	Email    string   `json:"email" form:"email" gorm:"not null;unique"`
	Password string   `json:"password" form:"password" gorm:"not null"`
	PhotoURL string   `json:"photo_url" form:"photo_url"`
	Role     string   `json:"role" form:"role" gorm:"type:enum('Pembina', 'Asisten', 'Member');default:'Member'"`
	Division tagDType `json:"division" form:"division" gorm:"type:enum('Backend', 'Frontend', 'UI/UX', 'Mobile', 'Audio Composer', 'Game Designer', 'Game Programmer', 'Game Art');default:'Backend'"`
	Token    string   `json:"token" form:"token"`
}

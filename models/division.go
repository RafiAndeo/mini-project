package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagDType string

const (
	Backend        tagDType = "Backend"
	Frontend       tagDType = "Frontend"
	UIUX           tagDType = "UI/UX"
	Mobile         tagDType = "Mobile"
	AudioComposer  tagDType = "Audio Composer"
	GameDesigner   tagDType = "Game Designer"
	GameProgrammer tagDType = "Game Programmer"
	GameArt        tagDType = "Game Art"
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
	Name         tagDType `json:"name" form:"name" gorm:"type:enum('Backend', 'Frontend', 'UI/UX', 'Mobile', 'Audio Composer', 'Game Designer', 'Game Programmer', 'Game Art');default:'Backend'"`
	ThumbnailURL string   `json:"thumbnail_url" form:"thumbnail_url"`
	UserID       uint     `json:"user_id" form:"user_id"`
}

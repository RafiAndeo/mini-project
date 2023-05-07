package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagPType string

const (
	Game    tagPType = "Game"
	Web     tagPType = "Web"
	Android tagPType = "Android"
)

func (t *tagPType) Scan(value interface{}) error {
	*t = tagPType(value.([]byte))
	return nil
}

func (t tagPType) Value() (driver.Value, error) {
	return string(t), nil
}

type Product struct {
	gorm.Model
	Name           string   `json:"name" form:"name"`
	Description    string   `json:"description" form:"description"`
	ThumbnailURL   string   `json:"thumbnail_url" form:"thumbnail_url"`
	Link           string   `json:"link" form:"link"`
	Tags           tagPType `json:"tags" form:"tags" gorm:"type:enum('Game', 'Web', 'Android');default:'Game'"`
	PublishedBy    string   `json:"published_by" form:"published_by"`
	PublishedAt    string   `json:"published_at" form:"published_at"`
	CreatedBy      string   `json:"created_by" form:"created_by"`
	UpdatedBy      string   `json:"updated_by" form:"updated_by"`
	UserActivityID uint     `json:"-" form:"-"`
}

package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type tagUAType string

type statusType string

const (
	Events   tagUAType = "Event"
	Blogs    tagUAType = "Blog"
	Products tagUAType = "Product"
)

const (
	Started   statusType = "Started"
	Postponed statusType = "Postponed"
	Canceled  statusType = "Canceled"
	Finished  statusType = "Finished"
)

func (t *tagUAType) Scan(value interface{}) error {
	*t = tagUAType(value.([]byte))
	return nil
}

func (t tagUAType) Value() (driver.Value, error) {
	return string(t), nil
}

func (t *statusType) Scan(value interface{}) error {
	*t = statusType(value.([]byte))
	return nil
}

func (t statusType) Value() (driver.Value, error) {
	return string(t), nil
}

type UserActivity struct {
	gorm.Model
	Name         string     `json:"name" form:"name"`
	Tags         tagUAType  `json:"tags" form:"tags" gorm:"enum('Event', 'Blog', 'Product');default:'Event'"`
	ThumbnailURL string     `json:"thumbnail_url" form:"thumbnail_url"`
	Description  string     `json:"description" form:"description"`
	Status       statusType `json:"status" form:"status" gorm:"enum('Started', 'Postponed', 'Canceled', 'Finished');default:'Started'"`
	UserID       uint       `json:"user_id" form:"user_id"`
}

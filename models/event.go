package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Event_ID       uint      `json:"event_id" form:"event_id" gorm:"primary_key"`
	Title          string    `json:"title" form:"title" gorm:"not null"`
	Tags           []string  `json:"tags" form:"tags" gorm:"type:VARCHAR(255)[]"`
	Description    string    `json:"description" form:"description"`
	Start_datetime string    `json:"start_datetime" form:"start_datetime"`
	End_datetime   string    `json:"end_datetime" form:"end_datetime"`
	Created_by     string    `json:"created_by" form:"created_by"`
	Updated_by     string    `json:"updated_by" form:"updated_by" gorm:"default:NULL"`
	Created_at     time.Time `json:"created_at" form:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated_at     time.Time `json:"updated_at" form:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Deleted_at     time.Time `json:"deleted_at" form:"deleted_at" gorm:"default:NULL"`
}

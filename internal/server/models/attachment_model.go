package models

import (
	"time"
)

type Attachment struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	TaskID    int64     `gorm:"index" json:"task_id"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
	Task      Task      `gorm:"foreignKey:TaskID" json:"task"`
}

package models

import (
	"time"
)

type Comment struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	TaskID    int64     `gorm:"index" json:"task_id"`
	UserID    int64     `gorm:"index" json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Task      Task      `gorm:"foreignKey:TaskID" json:"task"`
}

package models

import (
	"time"
)

type Task struct {
	ID          int64        `gorm:"primaryKey" json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	Priority    int          `json:"priority"`
	DueDate     time.Time    `json:"due_date"`
	AssignedTo  int64        `json:"assigned_to"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	User        User         `gorm:"foreignKey:AssignedTo" json:"user"`
	Comments    []Comment    `gorm:"foreignKey:TaskID" json:"comments"`
	Attachments []Attachment `gorm:"foreignKey:TaskID" json:"attachments"`
}

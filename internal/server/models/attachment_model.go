package models

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID        uuid.UUID `json:"id"`
	TaskID    uuid.UUID `json:"task_id"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
}

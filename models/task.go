package models

import (
	"time"
)

type (
	// Task ...
	Task struct {
		ID        string    `json:"_id"`
		Rev       string    `json:"_rev,omitempty"`
		Content   string    `json:"content"`
		Tags      string    `json:"tags"`
		CreatedAt time.Time `json:"created_at"`
		IsDeleted bool      `json:"is_deleted"`
		IsDone    bool      `json:"is_done"`
	}
)

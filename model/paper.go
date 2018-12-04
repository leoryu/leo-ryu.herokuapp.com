package model

import (
	"time"
)

// Paper is used for paper management
type Paper struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Subject   string    `json:"subject"`
	Abstract  string    `json:"abstract"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}


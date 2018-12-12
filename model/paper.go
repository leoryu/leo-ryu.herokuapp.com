package model

import (
	"time"
)

type Paper struct {
	*Introduction
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	EditedAt  time.Time `json:"edited_at" bson:"edited_at"`
}

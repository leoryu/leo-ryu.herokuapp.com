package model

import (
	"time"
)

type Paper struct {
	Introduction introduction `json:"introduction" bson:"introduction"`
	Content      string       `json:"content" bson:"content"`
	CreatedAt    time.Time    `json:"created_at" bson:"created_at"`
	EditedAt     time.Time    `json:"edited_at" bson:"edited_at"`
}

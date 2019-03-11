package core

import (
	"context"
)

type (
	Paper struct {
		ID        string          `json:"id,omitempty"`
		Title     string          `json:"title,omitempty"`
		Tags      map[string]bool `json:"tags,omitempty"`
		Abstract  string          `json:"abstract,omitempty"`
		Content   string          `json:"content,omitempty"`
		CreatedAt int64           `json:"created_at,omitempty"`
		UpdatedAt int64           `json:"edited_at,omitempty"`
	}

	PaperStore interface {
		Find(ctx context.Context, id string) (*Paper, error)
		List(ctx context.Context, limit, page int) ([]*Paper, error)
		ListByTag(ctx context.Context, tag string, limit, page int) ([]*Paper, error)
		Create(ctx context.Context, paper *Paper) error
		Update(ctx context.Context, paper *Paper) error
		Delete(ctx context.Context, id string) error
	}
)

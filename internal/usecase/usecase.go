package usecase

import (
	"context"
	"thought-RestAPI/internal/domain"
)

type ThoughtRepository interface {
	GetRandomThought(ctx context.Context) (*domain.Thought, error)
	CreateThought(ctx context.Context, text, author string) (int64, error)
	DeleteThought(ctx context.Context, id int64) error
	UpdateThought(ctx context.Context, id int64, text string, author string) error
}

type Thought struct {
	repo ThoughtRepository
}

func NewThought(repo ThoughtRepository) *Thought {
	return &Thought{repo: repo}
}

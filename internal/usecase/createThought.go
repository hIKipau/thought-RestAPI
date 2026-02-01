package usecase

import (
	"context"
	"fmt"
)

type CreateThoughtInput struct {
	Text   string
	Author string
}

func (t *Thought) CreateThought(ctx context.Context, input CreateThoughtInput) (int64, error) {
	const op = "internal/usecase/CreateThought"

	ID, err := t.repo.CreateThought(ctx, input.Text, input.Author)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return ID, nil
}

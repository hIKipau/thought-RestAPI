package usecase

import (
	"context"
	"fmt"
)

type UpdateThoughtInput struct {
	ID     int64
	Text   string
	Author string
}

func (t *Thought) UpdateThought(ctx context.Context, input UpdateThoughtInput) error {
	const op = "internal/usecase/UpdateThought"

	err := t.repo.UpdateThought(ctx, input.ID, input.Text, input.Author)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

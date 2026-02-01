package usecase

import (
	"context"
	"fmt"
)

func (t *Thought) DeleteThought(ctx context.Context, id int64) error {
	const op = "internal/usecase/DeleteThought"

	err := t.repo.DeleteThought(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

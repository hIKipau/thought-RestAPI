package usecase

import (
	"context"
	"fmt"
	"thought-RestAPI/internal/dto"
)

func (t *Thought) DeleteThought(ctx context.Context, input dto.DeleteThoughtInput) error {
	const op = "internal/usecase/DeleteThought"

	err := t.thoughtRepository.DeleteThought(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("%s: Cant delete thought, error: %s", op, err.Error())
	}

	return nil
}

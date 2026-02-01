package usecase

import (
	"context"
	"fmt"
	"thought-RestAPI/internal/dto"
)

func (t *Thought) UpdateThought(ctx context.Context, input dto.UpdateThoughtInput) error {
	const op = "internal/usecase/UpdateThought"

	err := t.thoughtRepository.UpdateThought(ctx, input.ID, input.Text, input.Author)
	if err != nil {
		return fmt.Errorf("%s: Cant update thought, error: %s", op, err.Error())
	}

	return nil
}

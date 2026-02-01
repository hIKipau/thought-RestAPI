package usecase

import (
	"context"
	"fmt"
	"thought-RestAPI/internal/dto"
)

func (t *Thought) CreateThought(ctx context.Context, input dto.CreateThoughtInput) (dto.CreateThoughtOutput, error) {
	const op = "internal/usecase/CreateThought"

	ID, err := t.thoughtRepository.CreateThought(ctx, input.Text, input.Author)
	if err != nil {
		return dto.CreateThoughtOutput{}, fmt.Errorf("%s: Cant create thought, error: %s", op, err.Error())
	}

	output := dto.CreateThoughtOutput{ID: ID}
	return output, nil
}

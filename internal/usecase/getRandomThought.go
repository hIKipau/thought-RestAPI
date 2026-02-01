package usecase

import (
	"context"
	"fmt"
	"thought-RestAPI/internal/dto"
)

func (t *Thought) GetRandomThought(ctx context.Context) (dto.GetRandomThoughtOutput, error) {
	const op = "internal/usecase/GetRandomThought"

	thought, err := t.thoughtRepository.GetRandomThought(ctx)
	if err != nil {
		return dto.GetRandomThoughtOutput{}, fmt.Errorf("%s: Cant get thought, error: %s", op, err.Error())
	}

	output := dto.GetRandomThoughtOutput{ID: thought.ID,
		Text:   thought.Text,
		Author: thought.Author,
	}

	return output, nil
}

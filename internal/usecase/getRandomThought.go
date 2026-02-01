package usecase

import (
	"context"
	"fmt"
	"thought-RestAPI/internal/domain"
)

func (t *Thought) GetRandomThought(ctx context.Context) (*domain.Thought, error) {
	const op = "internal/usecase/GetRandomThought"

	thought, err := t.repo.GetRandomThought(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return thought, nil
}

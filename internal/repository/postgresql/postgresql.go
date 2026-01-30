package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"thoughtRestApi/internal/model"
)

type PostgreSQL struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *PostgreSQL {
	return &PostgreSQL{
		db: db,
	}
}

func (pgsql *PostgreSQL) GetRandomThought(ctx context.Context) (*model.Thought, error) {
	row := pgsql.db.QueryRow(ctx,
		`SELECT id, text, author 
		 FROM thoughts 
		 WHERE id >= (SELECT floor(random() * (SELECT max(id) FROM thoughts))::bigint)
		 ORDER BY id 
		 LIMIT 1`)

	var t model.Thought
	err := row.Scan(&t.ID, &t.Text, &t.Author)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no thoughts found: %w", err)
		}
		return nil, fmt.Errorf("get random thought: %w", err)
	}

	return &t, nil
}

func (p *PostgreSQL) CreateThought(
	ctx context.Context,
	text, author string,
) (int64, error) {

	var id int64

	err := p.db.QueryRow(
		ctx,
		"INSERT INTO thoughts (text, author) VALUES ($1, $2) RETURNING id",
		text,
		author,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("create thought: %w", err)
	}

	return id, nil
}

func (pgsql *PostgreSQL) DeleteThought(ctx context.Context, id int64) error {
	cmdTag, err := pgsql.db.Exec(ctx, "DELETE FROM thoughts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("delete thought: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("delete thought: no row with id %d", id)
	}
	return nil
}

func (pgsql *PostgreSQL) UpdateThought(ctx context.Context, id int64, text string, author string) error {
	cmdTag, err := pgsql.db.Exec(
		ctx,
		"UPDATE thoughts SET text = $1, author = $2 WHERE id = $3",
		text,
		author,
		id,
	)
	if err != nil {
		return fmt.Errorf("update thought: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("update thought: no row with id %d", id)
	}
	return nil
}

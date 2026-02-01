package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"thought-RestAPI/internal/domain"
)

type PostgreSQL struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, databaseURL string, log *slog.Logger) (*PostgreSQL, error) {
	log.Info("Connecting to database")
	conn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database %s", err.Error())
	}
	log.Info("Successfully connected to database")
	return &PostgreSQL{
		db: conn,
	}, nil
}

func (pgsql *PostgreSQL) Close() {
	pgsql.db.Close()
}

func (pgsql *PostgreSQL) GetRandomThought(ctx context.Context) (*domain.Thought, error) {
	row := pgsql.db.QueryRow(ctx,
		`SELECT id, text, author 
		 FROM thoughts 
		 WHERE id >= (SELECT floor(random() * (SELECT max(id) FROM thoughts))::bigint)
		 ORDER BY id 
		 LIMIT 1`)

	var t domain.Thought
	err := row.Scan(&t.ID, &t.Text, &t.Author)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no thoughts found: %w", err)
		}
		return nil, fmt.Errorf("get random thought: %w", err)
	}

	return &t, nil
}

func (pgsql *PostgreSQL) CreateThought(ctx context.Context, text, author string) (int64, error) {

	var id int64

	err := pgsql.db.QueryRow(
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

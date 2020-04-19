package pokedex

import (
	"context"
	"database/sql"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jbowes/vice"

	"github.com/clwiseman/letsgopokemon/internal/models"
)

// Pokedex is a directory of pokemon.
type Pokedex struct {
	db     *sql.DB
	sb     squirrel.StatementBuilderType
}

// NewPokedex creates a new Pokedex, connecting it to the postgres server on
// the URL provided.
func NewPokedex() *Pokedex {
	pgURL := os.Getenv("DATABASE_URL")
	c, err := pgx.ParseURI(pgURL)
	if err != nil {
		panic("failed to parse database url")
	}

	db := stdlib.OpenDB(c)

	return &Pokedex{
		db:     db,
		sb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
	}
}

// Close releases any resources.
func (pd Pokedex) Close() error {
	return pd.db.Close()
}

// ListGenerations lists all the generations currently in the pokedex.
func (pd Pokedex) ListGenerations(ctx context.Context) ([]*models.Generation, error) {
	q := pd.sb.Select("id", "display_name").From("generation")

	rows, err := q.QueryContext(ctx)
	if err != nil {
		return nil, vice.Wrap(err, vice.Internal, "unable to list generations")
	}

	var generations []*models.Generation

	for rows.Next() {
		gen, err := scanGeneration(rows)
		if err != nil {
			return nil, vice.Wrap(err, vice.Internal, "error scanning row")
		}
		generations = append(generations, gen)
	}

	err = rows.Err()
	if err != nil {
		return nil, vice.Wrap(err, vice.Internal, "error during iteration")
	}

	err = rows.Close()
	if err != nil {
		return nil, vice.Wrap(err, vice.Internal, "error closing rows")
	}

	return generations, nil
}

// scanGeneration
func scanGeneration(row squirrel.RowScanner) (*models.Generation, error) {
	var gen *models.Generation
	err := row.Scan(gen.ID, gen.DisplayName)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, vice.Wrap(err, vice.Internal, "unable to scan row")
		}

		return nil, err
	}

	return gen, nil
}

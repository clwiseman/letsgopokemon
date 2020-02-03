package db

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/jbowes/vice"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"

	"github.com/clwiseman/letsgopokemon/internal/models"
)

// Pokedex is a directory of pokemon.
type Pokedex struct {
	logger *logrus.Logger
	db     *sql.DB
	sb     squirrel.StatementBuilderType
}

// NewPokedex creates a new Pokedex, connecting it to the postgres server on
// the URL provided.
func NewPokedex(logger *logrus.Logger, pgURL *url.URL) (*Pokedex, error) {
	c, err := pgx.ParseURI((pgURL.String()))
	if err != nil {
		return nil, err
	}

	c.Logger = logrusadapter.NewLogger(logger)
	db := stdlib.OpenDB(c)

	return &Pokedex{
		logger: logger,
		db:     db,
		sb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
	}, nil
}

// Close releases any resources.
func (pd Pokedex) Close() error {
	return pd.db.Close()
}

// // GetPokemonByID retrieves a pokemon by it's id (pokedex) number.
// func (pd Pokedex) GetPokemonByID(ctx context.Context, id *graphql) (err error) {
// 	return nil
// }

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
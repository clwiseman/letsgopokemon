package pokedex

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/clwiseman/letsgopokemon/internal/db_models"
)

// Pokedex is a directory of pokemon.
type Pokedex struct {
	db *pg.DB
}

// NewPokedex creates a new Pokedex, connecting it to the postgres server on
// the URL provided.
func NewPokedex(ctx context.Context, pgURL string) (*Pokedex, error) {
	dbOptions, err := pg.ParseURL(pgURL)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("could not parse connection URL.")
		return nil, err
	}

	db := pg.Connect(dbOptions)
	err = db_models.CreateSchema(db)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("could not create database schema.")
		return nil, err
	}

	err = db_models.InsertDefaults(db)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("could not insert default values in the database.")
		return nil, err
	}

	return &Pokedex{
		db: db,
	}, nil
}

// Close releases any resources.
func (pd Pokedex) Close() error {
	return pd.db.Close()
}

// GetPokemonByID retrieves a pokemon by it's id (pokedex) number.
func (pd Pokedex) GetPokemonByID(ctx context.Context, id int64) (db_models.Pokemon, error) {
	pokemon := db_models.Pokemon{
		BaseModel: db_models.BaseModel{
			Id: id, // This will fetch the right pokemon
		},
	}

	err := pd.db.WithContext(ctx).Select(&pokemon)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Could not query pokemon")
		return db_models.Pokemon{}, errors.Wrap(err, "could not query pokemon")
	}

	return pokemon, nil
}

// ListGenerations lists all the generations currently in the pokedex.
func (pd Pokedex) ListGenerations(ctx context.Context) ([]db_models.Generation, error) {
	var generations []db_models.Generation

	err := pd.db.WithContext(ctx).Model(&generations).Select()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Could not query generations")
		return nil, errors.Wrap(err, "could not query generations")
	}

	return generations, nil
}

// GetSession gets the session for the current game.
func (pd Pokedex) GetSession(ctx context.Context, userId int64) (db_models.Session, error) {
	session := db_models.Session{}

	err := pd.db.
		WithContext(ctx).
		Model(&session).
		Where("session.user_id = ?", userId).
		Select()
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Could not query game session")
		return db_models.Session{}, errors.Wrap(err, "could not query game session")
	}

	return session, nil
}

// CreateGame creates a new user and a new game session.
func (pd Pokedex) CreateGame(ctx context.Context, name string) (db_models.Session, error) {
	user := db_models.User{
		Name: name,
	}

	err := pd.db.WithContext(ctx).Insert(&user)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Could not create a new user")
		return db_models.Session{}, errors.Wrap(err, "could not create a new user")
	}

	session := db_models.Session{
		Users: []*db_models.User{&user},
	}

	err = pd.db.WithContext(ctx).Insert(&session)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Could not create a new session")
		return db_models.Session{}, errors.Wrap(err, "could not create a new session")
	}

	return session, nil
}

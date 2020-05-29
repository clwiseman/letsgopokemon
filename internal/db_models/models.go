package db_models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// BaseModel is a generic struct that implements the time based indices and the ID
type BaseModel struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

var _ pg.BeforeInsertHook = (*BaseModel)(nil)

// BeforeInsert is a hook that will set the times when an element is about to get inserted.
func (m *BaseModel) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return ctx, nil
}

var _ pg.BeforeUpdateHook = (*BaseModel)(nil)

// BeforeUpdate is a hook that will set the times when an element is about to get updated.
func (m *BaseModel) BeforeUpdate(ctx context.Context) (context.Context, error) {
	m.UpdatedAt = time.Now()
	return ctx, nil
}

// Session is a game session owning users and rounds.
type Session struct {
	BaseModel

	Users  []*User
	Rounds []*Round
}

// User is a user within a game session
type User struct {
	BaseModel

	Name    string
	Session *Session
}

// Generation is a pokemon generation
type Generation struct {
	BaseModel

	Name     string
	Pokemons []*Pokemon
}

// Pokemon represents a drawable pokemon, owned by a generation.
type Pokemon struct {
	BaseModel

	Name       string
	ImageURL   string
	Generation *Generation
}

// Round is a game round within a game session. Each round will have multiple turns for each users in the session.
type Round struct {
	BaseModel

	Session *Session
	Turns   []*Turn
}

// Turn is a turn within a round. Each players takes on turn drawing their pokemon.
type Turn struct {
	BaseModel

	User    *User
	Pokemon *Pokemon
	Drawing string
	Round   *Round
}

// CreateSchema creates the database schemas based on the models.
func CreateSchema(db *pg.DB) error {
	tables := []interface{}{
		(*Session)(nil),
		(*User)(nil),
		(*Generation)(nil),
		(*Pokemon)(nil),
		(*Round)(nil),
		(*Turn)(nil),
	}

	for _, model := range tables {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

package testing

import (
	"context"

	"github.com/clwiseman/letsgopokemon/internal/db_models"
)

// FakePokedex contains fake pokedex for testing.
type FakePokedex struct {
	ListGenerationsFunc func(ctx context.Context) ([]db_models.Generation, error)
}

// ListGenerations fake.
func (f *FakePokedex) ListGenerations(ctx context.Context) ([]db_models.Generation, error) {
	return f.ListGenerations(ctx)
}

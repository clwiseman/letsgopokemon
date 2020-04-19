package resolver

import (
	"context"

	"github.com/clwiseman/letsgopokemon/internal/models"
	pd "github.com/clwiseman/letsgopokemon/internal/pokedex"
	"github.com/jbowes/vice"
)

// Resolver tracks the state
type Resolver struct {
	generations []*models.Generation
}

type QueryResolver interface {
	Generations(ctx context.Context) ([]*models.Generation, error)
	Pokemon(ctx context.Context, input string) (*models.Pokemon, error)
}

type MutationResolver interface {
	CreateGame(ctx context.Context, input models.CreateGameInput) (*models.CreateGamePayload, error)
	JoinGame(ctx context.Context, input models.JoinGameInput) (*models.JoinGamePayload, error)
	StartTurn(ctx context.Context, input models.StartTurnInput) (*models.StartTurnPayload, error)
	EndTurn(ctx context.Context, input models.EndTurnInput) (*models.EndTurnPayload, error)
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateGame(ctx context.Context, input models.CreateGameInput) (*models.CreateGamePayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) JoinGame(ctx context.Context, input models.JoinGameInput) (*models.JoinGamePayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) StartTurn(ctx context.Context, input models.StartTurnInput) (*models.StartTurnPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) EndTurn(ctx context.Context, input models.EndTurnInput) (*models.EndTurnPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Generations(ctx context.Context) ([]*models.Generation, error) {
	gens, err := pd.Pokedex.ListGenerations(ctx)
	if err != nil {
		return nil, vice.Wrap(err, vice.NotFound, "no generations found")
	}

	r.generations = gens

	return r.generations, nil
}
func (r *queryResolver) Pokemon(ctx context.Context, input string) (*models.Pokemon, error) {
	panic("not implemented")
}

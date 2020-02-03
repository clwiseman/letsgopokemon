package generated

import (
	"context"

	"github.com/clwiseman/letsgopokemon/internal/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

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
	panic("not implemented")
}
func (r *queryResolver) Pokemon(ctx context.Context, input string) (*models.Pokemon, error) {
	panic("not implemented")
}

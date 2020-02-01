package graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateGame(ctx context.Context, input CreateGameInput) (*CreateGamePayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) JoinGame(ctx context.Context, input JoinGameInput) (*JoinGamePayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) StartTurn(ctx context.Context, input StartTurnInput) (*StartTurnPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) EndTurn(ctx context.Context, input EndTurnInput) (*EndTurnPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Generations(ctx context.Context) (*Generation, error) {
	panic("not implemented")
}
func (r *queryResolver) Pokemon(ctx context.Context, input string) (*Pokemon, error) {
	panic("not implemented")
}

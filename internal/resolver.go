package generated

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/clwiseman/letsgopokemon/internal/converters"
	"github.com/clwiseman/letsgopokemon/internal/generated"
	"github.com/clwiseman/letsgopokemon/internal/models"
	"github.com/clwiseman/letsgopokemon/internal/pokedex"
)

type Observer struct {
	UserId  string
	Message chan *models.GameSession
}

type Room struct {
	Id        string
	Session   *models.GameSession
	Observers map[string]Observer
}

type Resolver struct {
	pokedx *pokedex.Pokedex

	Rooms map[string]*Room
	mu    sync.Mutex // nolint: structcheck
}

func NewResolver(ctx context.Context, pokedexURL string) (*Resolver, error) {
	pokedx, err := pokedex.NewPokedex(ctx, pokedexURL)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		pokedx: pokedx,
	}, nil
}

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

func (r *queryResolver) Generations(ctx context.Context) ([]*models.Generation, error) {
	result, err := r.pokedx.ListGenerations(ctx)
	if err != nil {
		return nil, err
	}

	return converters.ConvertGenerations(result), nil
}

func (r *queryResolver) Pokemon(ctx context.Context, input string) (*models.Pokemon, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("could not convert inputted id")
		return nil, fmt.Errorf("invalid ID provided")
	}

	result, err := r.pokedx.GetPokemonByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return converters.ConvertPokemon(result), nil
}

func (r *subscriptionResolver) Session(ctx context.Context, userID string) (<-chan *models.GameSession, error) {
	id, err := strconv.Atoi(userID)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("could not convert inputted id")
		return nil, fmt.Errorf("invalid User ID provided")
	}

	result, err := r.pokedx.GetSession(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	session := converters.ConvertSession(result)

	r.mu.Lock()
	room := r.Rooms[session.ID]
	if room == nil {
		room = &Room{
			Id:        session.ID,
			Session:   session,
			Observers: map[string]Observer{},
		}
		r.Rooms[session.ID] = room
	}
	r.mu.Unlock()

	events := make(chan *models.GameSession, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(room.Observers, session.ID)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	room.Observers[session.ID] = Observer{
		UserId:  userID,
		Message: events,
	}
	r.mu.Unlock()

	return events, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

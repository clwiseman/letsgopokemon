package converters

import (
	"github.com/clwiseman/letsgopokemon/internal/db_models"
	"github.com/clwiseman/letsgopokemon/internal/models"
)

func ConvertSession(session db_models.Session) *models.GameSession {
	converted := &models.GameSession{
		ID: string(session.Id),
	}

	if session.Users != nil {
		temp := make([]db_models.User, len(session.Users))

		for index, el := range session.Users {
			temp[index] = *el
		}

		converted.Users = ConvertUsers(temp)
	}

	if session.Rounds != nil {
		temp := make([]db_models.Round, len(session.Rounds))

		for index, el := range session.Rounds {
			temp[index] = *el
		}

		converted.Rounds = ConvertRounds(temp)
	}

	if converted.Rounds != nil && len(converted.Rounds) > 0 {
		converted.CurrentRound = converted.Rounds[0]
	}

	return converted
}

func ConvertUser(user db_models.User) *models.User {
	return &models.User{
		ID:   string(user.Id),
		Name: user.Name,
	}
}

func ConvertUsers(users []db_models.User) []*models.User {
	converted := make([]*models.User, len(users))

	for index, user := range users {
		converted[index] = ConvertUser(user)
	}

	return converted
}

func ConvertGeneration(generation db_models.Generation) *models.Generation {
	converted := &models.Generation{
		ID:          string(generation.Id),
		DisplayName: generation.Name,
	}

	if generation.Pokemons != nil {
		temp := make([]db_models.Pokemon, len(generation.Pokemons))

		for index, el := range generation.Pokemons {
			temp[index] = *el
		}

		converted.Pokemons = ConvertPokemons(temp)
	}

	return converted
}

func ConvertGenerations(generations []db_models.Generation) []*models.Generation {
	converted := make([]*models.Generation, len(generations))

	for index, generation := range generations {
		converted[index] = ConvertGeneration(generation)
	}

	return converted
}

func ConvertPokemon(pokemon db_models.Pokemon) *models.Pokemon {
	converted := &models.Pokemon{
		ID:    string(pokemon.Id),
		Name:  pokemon.Name,
		Image: pokemon.ImageURL,
	}

	if pokemon.Generation != nil {
		converted.Generation = ConvertGeneration(*pokemon.Generation)
	}

	return converted
}

func ConvertPokemons(pokemons []db_models.Pokemon) []*models.Pokemon {
	converted := make([]*models.Pokemon, len(pokemons))

	for index, pokemon := range pokemons {
		converted[index] = ConvertPokemon(pokemon)
	}

	return converted
}

func ConvertRound(round db_models.Round) *models.GameRound {
	converted := &models.GameRound{
		ID: string(round.Id),
	}

	if round.Session != nil {
		converted.Session = ConvertSession(*round.Session)
	}

	if round.Turns != nil {
		converted.Turns = ConvertTurns(round.Turns)
	}

	return converted
}

func ConvertRounds(rounds []db_models.Round) []*models.GameRound {
	converted := make([]*models.GameRound, len(rounds))

	for index, round := range rounds {
		converted[index] = ConvertRound(round)
	}

	return converted
}

func ConvertTurn(turn db_models.Turn) *models.Turn {
	converted := &models.Turn{
		ID:      string(turn.Id),
		Drawing: &models.Drawing{URL: turn.Drawing},
	}

	if turn.Pokemon != nil {
		converted.Pokemon = ConvertPokemon(*turn.Pokemon)
	}

	if turn.User != nil {
		converted.User = ConvertUser(*turn.User)
	}

	return converted
}

func ConvertTurns(turns []*db_models.Turn) []*models.Turn {
	converted := make([]*models.Turn, len(turns))

	for index, turn := range turns {
		converted[index] = ConvertTurn(*turn)
	}

	return converted
}

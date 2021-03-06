// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

// CreateGameInput represents the input to the createGame mutation
type CreateGameInput struct {
	NewUser *NewUser `json:"newUser"`
}

// CreateGamePayload represents the output to the createGame mutation
type CreateGamePayload struct {
	GameSession *GameSession `json:"gameSession"`
}

// Drawing represents the drawing done by a user on the canvas
type Drawing struct {
	URL string `json:"url"`
}

// EndTurnInput represents the input to the endTurn mutation
type EndTurnInput struct {
	TurnID     string      `json:"turnId"`
	NewDrawing *NewDrawing `json:"newDrawing"`
}

// EndTurnPayload represents the output to the endTurn mutation
type EndTurnPayload struct {
	Turn *Turn `json:"turn"`
}

// GameInvite represents the information needed to add a new player to a session
type GameInvite struct {
	SessionID string `json:"sessionId"`
	UserName  string `json:"userName"`
}

// GameRound is the collection of player turns that occurs at the same time
type GameRound struct {
	ID      string       `json:"id"`
	Session *GameSession `json:"session"`
	Turns   []*Turn      `json:"turns"`
}

// GameSession represents the session users can join to play the game together
type GameSession struct {
	ID           string       `json:"id"`
	Users        []*User      `json:"users"`
	Rounds       []*GameRound `json:"rounds"`
	CurrentRound *GameRound   `json:"currentRound"`
}

// Generation represents the pokemon generational period.
type Generation struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"displayName"`
	Pokemons    []*Pokemon `json:"pokemons"`
}

// JoinGameInput represents the input to the joinGame mutation
type JoinGameInput struct {
	Invite *GameInvite `json:"invite"`
}

// JoinGamePayload represents the output to the joinGame mutation
type JoinGamePayload struct {
	Session *GameSession `json:"session"`
}

// NewDrawing represents the drawing done by a user on the canvas
type NewDrawing struct {
	Drawing string `json:"drawing"`
}

// StartTurnInput represents the information needed to start a turn
type NewTurn struct {
	UserID  string `json:"userId"`
	RoundID string `json:"roundId"`
}

// NewUser represents a new user to be created
type NewUser struct {
	DisplayName string `json:"displayName"`
}

// Pokemon represents the pokedex information for a pokemon.
type Pokemon struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Image      string      `json:"image"`
	Generation *Generation `json:"generation"`
}

// StartTurnInput represents the input to the startTurn mutation
type StartTurnInput struct {
	NewTurn *NewTurn `json:"newTurn"`
}

// StartTurnPayload represents the output to the startTurn mutation
type StartTurnPayload struct {
	Turn *Turn `json:"turn"`
}

// A Turn represents the gameplay for a user within a round
type Turn struct {
	ID      string   `json:"id"`
	User    *User    `json:"user"`
	Pokemon *Pokemon `json:"pokemon"`
	Drawing *Drawing `json:"drawing"`
}

// User represents a player in the game
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

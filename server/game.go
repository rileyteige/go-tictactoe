package server

import (
	"fmt"
	"net/http"
)

type Team string

const (
	X = Team("X")
	O = Team("O")
)

type Player struct {
	Id   Guid
	Team Team
}

type Game struct {
	Id      Guid
	Current Team
	X       Guid
	O       Guid
}

type PlayerGame struct {
	Game   Guid
	Player Player
}

var games = map[Guid]*Game{}

type StartParameters struct {
	Team string
}

func newGame() (game *Game) {
	game = new(Game)
	game.Id = generateGuid()
	game.Current = X
	game.X = emptyGuid()
	game.O = emptyGuid()
	games[game.Id] = game
	return
}

func getGame(w http.ResponseWriter, r *http.Request) {
	vars := pullVars(r)
	id, err := parseGuid(vars["id"])
	if err != nil {
		errorStandardText(w, http.StatusBadRequest)
		return
	}

	game := games[id]
	if game == nil {
		http.NotFound(w, r)
		return
	}

	writeJson(w, game)
}

func startGame(w http.ResponseWriter, r *http.Request) {
	var args StartParameters
	err := readJson(r, &args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Starting game with player %v\n", args.Team)

	team, playerId := Team(args.Team), generateGuid()
	game := newGame()
	switch team {
	case X:
		game.X = playerId
	case O:
		game.O = playerId
	}

	pg := PlayerGame{Game: game.Id, Player: Player{Id: playerId, Team: team}}

	writeJson(w, pg)
}

type JoinParameters struct {
	Team string
	Game Guid
}

func joinGame(w http.ResponseWriter, r *http.Request) {
	var args JoinParameters
	err := readJson(r, &args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Player %v joining game %v\n", args.Team, args.Game)

	team, playerId := Team(args.Team), generateGuid()
	game := games[args.Game]
	if game == nil {
		http.NotFound(w, r)
		return
	}

	switch team {
	case X:
		if game.X != emptyGuid() {
			errorStandardText(w, http.StatusForbidden)
			return
		}

		game.X = playerId
	case O:
		if game.O != emptyGuid() {
			errorStandardText(w, http.StatusForbidden)
			return
		}

		game.O = playerId
	}

	pg := PlayerGame{Game: game.Id, Player: Player{Id: playerId, Team: team}}
	writeJson(w, pg)
}

package server

import (
	"fmt"
	"net/http"
)

type Game struct {
	Id Guid
}

var games = map[Guid]*Game{}

type StartParameters struct {
	Player string
}

func newGame() (game *Game) {
	game = new(Game)
	game.Id = generateGuid()
	games[game.Id] = game
	return
}

func getGame(w http.ResponseWriter, r *http.Request) {
	vars := pullVars(r)
	id := Guid(vars["id"])

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

	fmt.Printf("Starting game with player %v\n", args.Player)

	game := newGame()

	writeJson(w, game)
}

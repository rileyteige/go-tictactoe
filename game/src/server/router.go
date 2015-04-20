package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func pullVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func buildRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/ping", ping).Methods(methodGet)
	r.HandleFunc("/game", startGame).Methods(methodPost)
	r.HandleFunc("/game/{id}", getGame).Methods(methodGet)

	return r
}

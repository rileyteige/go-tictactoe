package server

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != methodGet {
		http.Error(w, "Method not allowed", 405)
		return
	}

	//	echo := r.URL.Query().Get("txt")
	//	if echo == "" {
	//		http.Error(w, "Bad request", http.StatusBadRequest)
	//	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
}

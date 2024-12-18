package main

import (
	"context"
	"net/http"
)

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Cannot erase users", nil)
		return
	}

	err := cfg.db.ResetUsers(context.Background())
	if err != nil {
		respondWithError(w, 500, "Something went wrong", err)
		return
	}

	cfg.fileServerHits.Store(0)
	w.WriteHeader(200)
	w.Write([]byte("Hits reset to 0 and chirpy database reset to initial state."))
}

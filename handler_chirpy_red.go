package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpyRed(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	params := parameters{}
	decodeJSONBody(w, req, &params)

	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}

	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user_id", err)
		return
	}

	err = cfg.db.UpdateChirpyRedStatus(context.Background(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to find user", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

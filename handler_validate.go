package main

import (
	"encoding/json"
	"net/http"
)

func handlerChirpsValidate(w http.ResponseWriter, req *http.Request) {
	type Chirp struct {
		Body string `json:"body"`
	}

	type validResponse struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(req.Body)
	chirp := Chirp{}
	err := decoder.Decode(&chirp)
	if err != nil {
		respondWithError(w, 500, "Failed to decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(chirp.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, validResponse{
		Valid: true,
	})
}

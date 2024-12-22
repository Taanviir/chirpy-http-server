package main

import (
	"context"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirpsByID(w http.ResponseWriter, req *http.Request) {
	chirpID, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID entered", err)
		return
	}

	chirp, err := cfg.db.GetChirpById(context.Background(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to find chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, req *http.Request) {
	dbChirps, err := cfg.db.GetChirps(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get chirps", err)
		return
	}

	authorID := uuid.Nil
	authorIDString := req.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author id", err)
			return
		}
	}

	sortOrder := "asc"
	sortOrderParam := req.URL.Query().Get("sort")
	if sortOrderParam == "desc" {
		sortOrder = "desc"
	}

	chirps := []Chirp{}
	for _, chirp := range dbChirps {
		if authorID != uuid.Nil && chirp.UserID != authorID {
			continue
		}

		chirps = append(chirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	sort.Slice(chirps, func(i int, j int) bool {
		if sortOrder == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

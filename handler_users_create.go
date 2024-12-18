package main

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, req *http.Request) {
	type requestValues struct {
		Email string `json:"email"`
	}

	userInfo := requestValues{}
	decodeJSONBody(w, req, &userInfo)

	user, err := cfg.db.CreateUsers(context.Background(), userInfo.Email)
	if err != nil {
		respondWithError(w, 500, "Something went wrong", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}

package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Taanviir/chirpy/internal/auth"
	"github.com/Taanviir/chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, req *http.Request) {
	type createUserRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	userInfo := createUserRequest{}
	decodeJSONBody(w, req, &userInfo)

	hashedPassword, err := auth.HashPassword(userInfo.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error creating user", err)
		return
	}

	user, err := cfg.db.CreateUsers(context.Background(), database.CreateUsersParams{
		Email: userInfo.Email,
		HashedPassword: hashedPassword,
	})
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

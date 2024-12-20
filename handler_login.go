package main

import (
	"context"
	"net/http"

	"github.com/Taanviir/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, req *http.Request) {
	type loginUserRequest struct {
		Password string `json:"password"`
		Email string `json:"email"`
	}

	userInfo := loginUserRequest{}
	decodeJSONBody(w, req, &userInfo)

	user, err := cfg.db.GetUserByEmail(context.Background(), userInfo.Email)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Incorrect email or password", err)
		return
	}

	if err := auth.CheckPasswordHash(userInfo.Password, user.HashedPassword); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	})
}

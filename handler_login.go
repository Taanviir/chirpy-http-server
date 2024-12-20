package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Taanviir/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, req *http.Request) {
	type loginUserRequest struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
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

	timeToExpire := userInfo.ExpiresInSeconds
	if timeToExpire > int(time.Hour.Seconds()) || timeToExpire == 0 {
		timeToExpire = int(time.Hour.Seconds())
	}

	token, err := auth.MakeJWT(user.ID, cfg.tokenSecret, time.Duration(timeToExpire)*time.Second)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Failed to make auth token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		JWT:       token,
	})
}

package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CoDanTheBarbarian/chirpy/internal/auth"
	"github.com/CoDanTheBarbarian/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.dbQueries.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email", err)
		return
	}
	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect password", err)
		return
	}
	dbUser := User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	}
	token, err := auth.MakeJWT(dbUser.ID, cfg.jwtSecret, time.Duration(1)*time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate token", err)
		return
	}
	dbUser.Token = token
	refTkn, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate refresh token", err)
		return
	}
	refreshToken, err := cfg.dbQueries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refTkn,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Duration(60 * 24 * time.Hour)),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}
	dbUser.RefreshToken = refreshToken.Token
	respondWithJSON(w, http.StatusOK, dbUser)
}

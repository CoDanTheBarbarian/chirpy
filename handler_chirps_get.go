package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")
	sort := r.URL.Query().Get("sort")
	reverse := false
	if sort == "desc" {
		reverse = true
	}
	if authorID != "" {
		userID, err := uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Couldn't parse user ID", err)
			return
		}
		dbchirps, err := cfg.dbQueries.GetChirpsForUser(r.Context(), userID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
			return
		}
		chirps := make([]Chirp, len(dbchirps))
		if reverse {
			for i, dbchirp := range dbchirps {
				chirps[len(dbchirps)-i-1] = Chirp{
					ID:        dbchirp.ID,
					CreatedAt: dbchirp.CreatedAt,
					UpdatedAt: dbchirp.UpdatedAt,
					Body:      dbchirp.Body,
					UserID:    dbchirp.UserID,
				}
			}
		} else {
			for i, dbchirp := range dbchirps {
				chirps[i] = Chirp{
					ID:        dbchirp.ID,
					CreatedAt: dbchirp.CreatedAt,
					UpdatedAt: dbchirp.UpdatedAt,
					Body:      dbchirp.Body,
					UserID:    dbchirp.UserID,
				}
			}
		}

		respondWithJSON(w, 200, chirps)
		return
	}
	dbchirps, err := cfg.dbQueries.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}

	chirps := make([]Chirp, len(dbchirps))
	if reverse {
		for i, dbchirp := range dbchirps {
			chirps[len(dbchirps)-i-1] = Chirp{
				ID:        dbchirp.ID,
				CreatedAt: dbchirp.CreatedAt,
				UpdatedAt: dbchirp.UpdatedAt,
				Body:      dbchirp.Body,
				UserID:    dbchirp.UserID,
			}
		}
	} else {
		for i, dbchirp := range dbchirps {
			chirps[i] = Chirp{
				ID:        dbchirp.ID,
				CreatedAt: dbchirp.CreatedAt,
				UpdatedAt: dbchirp.UpdatedAt,
				Body:      dbchirp.Body,
				UserID:    dbchirp.UserID,
			}
		}
	}

	respondWithJSON(w, 200, chirps)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse chirp ID", err)
		return
	}
	dbchirp, err := cfg.dbQueries.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "Chirp not found", err)
		return
	}
	chirp := Chirp{
		ID:        dbchirp.ID,
		CreatedAt: dbchirp.CreatedAt,
		UpdatedAt: dbchirp.UpdatedAt,
		Body:      dbchirp.Body,
		UserID:    dbchirp.UserID,
	}
	respondWithJSON(w, 200, chirp)
}

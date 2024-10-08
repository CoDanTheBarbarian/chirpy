package main

import "net/http"

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	dbchirps, err := cfg.dbQueries.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps", err)
		return
	}

	chirps := make([]Chirp, len(dbchirps))
	for i, dbchirp := range dbchirps {
		chirps[i] = Chirp{
			ID:        dbchirp.ID,
			CreatedAt: dbchirp.CreatedAt,
			UpdatedAt: dbchirp.UpdatedAt,
			Body:      dbchirp.Body,
			UserID:    dbchirp.UserID,
		}
	}
	respondWithJSON(w, 200, chirps)
}

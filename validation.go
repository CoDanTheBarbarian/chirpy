package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Body string `json:"body"`
	}

	type validResponse struct {
		Valid bool `json:"valid"`
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	var p params
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		errResp := errorResponse{
			Error: "Invalid JSON",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(dat)
		return
	}

	if len(p.Body) > 140 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		errResp := errorResponse{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(dat)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	validResp := validResponse{
		Valid: true,
	}
	dat, err := json.Marshal(validResp)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(dat)
}

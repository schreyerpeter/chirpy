package main

import (
	"chirpy/internal/auth"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsDeleteByID(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid authorization token", err)
		return
	}


	userId, err := auth.ValidateJWT(token, cfg.secretKey)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid authorization token", err)
		return
	}

	id := r.PathValue("chirpID")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID format", err)
		return
	}

	dbChirp, err := cfg.db.GetChirpsByID(r.Context(), uuidID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirp", err)
		return
	}

	if dbChirp.UserID != userId {
		respondWithError(w, http.StatusForbidden, "You are not authorized to delete this chirp", nil)
		return
	}

	err = cfg.db.DeleteChirpByID(r.Context(), uuidID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirp", err)
		return
	}


	respondWithJSON(w, http.StatusNoContent, nil)
}

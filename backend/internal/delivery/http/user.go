package http

import (
	"backend/internal/domain"
	"encoding/json"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user domain.UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	id, err := h.services.UsersService.SignUp(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]uint{"id": id})
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	tokens, err := h.services.UsersService.SignIn(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tokens)
}

func (h *Handler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Authorization")
	if refreshToken == "" {
		respondWithError(w, http.StatusUnauthorized, "empty refresh token")
		return
	}

	tokens, err := h.services.UsersService.RefreshTokens(refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tokens)
}

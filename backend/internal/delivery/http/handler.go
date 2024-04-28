package http

import (
	"backend/internal/service"
	"backend/pkg/auth"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	services *service.Services

	tokenManager auth.Manager
}

func NewHandler(services *service.Services, tokenManager auth.Manager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) InitRoutes() {
	router := chi.NewRouter()

	router.Route("/user", func(r chi.Router) {
		r.Post("/sign-up", h.SignUp)
		r.Post("/sign-in", h.SignIn)
		r.Post("/refresh-tokens", h.RefreshTokens)
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	var response []byte
	var err error

	if payload != nil {
		response, err = json.Marshal(payload)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Ошибка при маршалинге JSON")
			return
		}
	} else {
		response = []byte("{}")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]interface{}{"error": message})
}

package register_game_id_get

import (
	"context"
	"encoding/json"
	"net/http"
)

//go:generate ../../../bin/mockgen -destination=./deps_mocks_test.go -package=register_game_id_get -source=game_handler.go

type Service interface {
	RegisterGame(context.Context) (string, error)
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler { return &Handler{svc: svc} }

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/game", h.createGame)
}
func (h *Handler) createGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	id, err := h.svc.RegisterGame(ctx)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "failed to save game")
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(CreateGameResponse{ID: id})
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}

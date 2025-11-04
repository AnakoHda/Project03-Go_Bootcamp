package game_post

import (
	"Project03-Go_Bootcamp/internal/domain"

	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

//go:generate ../../../bin/mockgen -destination=./deps_mocks_test.go -package=game_post -source=game_handler.go

type Service interface {
	WriteNewState(ctx context.Context, prevId string, next domain.Game) (domain.Game, error)
	NextMove(ctx context.Context, id string) (domain.Game, error)
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/game/", h.postGame)
}

func (h *Handler) postGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/game/")
	if id == "" || strings.Contains(id, "/") {
		writeErr(w, http.StatusBadRequest, "bad game id")
		slog.Error("bad game id", "id", id)
		return
	}

	var req GameWebRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		slog.Error("invalid json", "err", err, "body", r.Body)
		return
	}

	ctx := r.Context()

	updatedGame, err := h.svc.WriteNewState(ctx, id, FromWeb(req))
	if err != nil {
		//Todo: distinguish error types
		writeErr(w, http.StatusNotFound, "game not found")
		slog.Error("WrireNewState: ", "err", err, "id ", id)
		return
	}
	//Human win or draw
	if updatedGame.GetWinner() != domain.NoneWinner {
		resp := ToWeb(updatedGame)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}
	updatedGame, err = h.svc.NextMove(ctx, id)
	if err != nil {
		writeErr(w, http.StatusNotFound, "game not found")
		slog.Error("game not found", "id", id, "err", err)
		return
	}
	if updatedGame.GetWinner() != domain.NoneWinner {
		resp := ToWeb(updatedGame)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}
	resp := ToWeb(updatedGame)
	_ = json.NewEncoder(w).Encode(resp)
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}

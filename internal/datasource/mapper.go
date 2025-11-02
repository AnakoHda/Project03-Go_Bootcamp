package datasource

import (
	"Project03-Go_Bootcamp/internal/domain"
)

func ToDS(g domain.Game) (GameDS, string) {
	return GameDS{
		board:  g.Board,
		turn:   g.Turn,
		winner: g.Winner,
	}, g.ID
}
func ToDomain(g GameDS, ID string) domain.Game {
	return domain.Game{
		ID:     ID,
		Board:  g.board,
		Turn:   g.turn,
		Winner: g.winner}
}

package game_post

import (
	"Project03-Go_Bootcamp/internal/domain"
)

func ToWeb(g domain.Game) GameWebResponse {
	return GameWebResponse{
		Board:  g.Board().Matrix(),
		Winner: g.GetWinner(),
	}
}

func FromWeb(g GameWebRequest) domain.Game {
	return domain.SetGame(domain.NewBoard(g.Board), domain.TurnO, domain.NoneWinner)
}

package datasource

import "Project03-Go_Bootcamp/internal/domain"

type GameDS struct {
	board  domain.Board
	turn   int
	winner int
}

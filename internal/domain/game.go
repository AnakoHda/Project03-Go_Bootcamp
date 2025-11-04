package domain

import (
	"errors"
)

const (
	TurnX = 10
	TurnO = 11
)
const (
	NoneWinner = 20
	Drow       = 21
	WinnerX    = 22
	WinnerO    = 23
)

type Game struct {
	board  Board
	turn   int
	winner int
}

func NewGame() Game {
	return Game{
		board:  Board{},
		turn:   TurnX,
		winner: NoneWinner,
	}
}
func SetGame(b Board, t, w int) Game {
	return Game{
		board:  b,
		turn:   t,
		winner: w,
	}
}
func (g *Game) ValidateNextState(next Game) error {
	if g.GetWinner() != NoneWinner || next.GetWinner() != NoneWinner {
		return errors.New("game finished")
	}

	prevXNum, prevONum := 0, 0
	nextXNum, nextONum := 0, 0
	changes := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a := g.board.GetPoint(i, j)
			b := next.board.GetPoint(i, j)

			if a == PointValueX {
				prevXNum++
			}
			if a == PointValueO {
				prevONum++
			}
			if b == PointValueX {
				nextXNum++
			}
			if b == PointValueO {
				nextONum++
			}

			if a != PointValueEmpty && a != b {
				return errors.New("attempt to overwrite existing cell")
			}

			if a != b {
				changes++
			}
		}
	}
	if changes != 1 {
		return errors.New("exactly one cell must change")
	}

	switch g.GetTurn() {
	case TurnX:
		if !(nextXNum == prevXNum+1 && nextONum == prevONum) {
			return errors.New("x must add exactly one mark")
		}
	case TurnO:
		if !(nextONum == prevONum+1 && nextXNum == prevXNum) {
			return errors.New("o must add exactly one mark")
		}
	default:
		return errors.New("invalid turn")
	}
	return nil
}

func (g *Game) SetGameWinner() (int, bool) {
	if winner, _ := g.board.CheckWinner(PointValueO); winner == true {
		g.winner = WinnerO
		return WinnerO, true
	}
	if winner, _ := g.board.CheckWinner(PointValueX); winner == true {
		g.winner = WinnerX
		return WinnerX, true
	}

	if full := g.board.IsFull(); full == true {
		g.winner = Drow
		return Drow, true
	}

	return -1, false
}

func (g *Game) GetWinner() int {
	return g.winner
}
func (g *Game) GetTurn() int {
	return g.turn
}
func (g *Game) NextTour() {
	if g.turn == TurnX {
		g.turn = TurnO
	} else if g.turn == TurnO {
		g.turn = TurnX
	}
}
func (g *Game) Board() *Board {
	return &g.board
}

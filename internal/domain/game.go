package domain

import "errors"

const (
	PointValueEmpty = 0
	PointValueO     = 1
	PointValueX     = 2

	GameStatusDraw = 3
)

type Game struct {
	ID     string //UUID
	Board  Board
	Turn   int
	Winner int
}

func ValidateBoards(prev, next Game) error {
	if prev.Winner != PointValueEmpty || next.Winner != PointValueEmpty {
		return errors.New("game finished")
	}

	prevX, prevO := 0, 0
	nextX, nextO := 0, 0
	changes := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a := prev.Board.Matrix[i][j]
			b := next.Board.Matrix[i][j]

			if a == PointValueX {
				prevX++
			}
			if a == PointValueO {
				prevO++
			}
			if b == PointValueX {
				nextX++
			}
			if b == PointValueO {
				nextO++
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

	switch prev.Turn {
	case PointValueX:
		if !(nextX == prevX+1 && nextO == prevO) {
			return errors.New("x must add exactly one mark")
		}
	case PointValueO:
		if !(nextO == prevO+1 && nextX == prevX) {
			return errors.New("o must add exactly one mark")
		}
	default:
		return errors.New("invalid turn")
	}
	return nil
}

func (g *Game) CheckGameCompete() bool {
	m := g.Board.Matrix

	for i := 0; i < 3; i++ {
		if m[i][0] != PointValueEmpty && m[i][0] == m[i][1] && m[i][1] == m[i][2] {
			g.Winner = m[i][0]
			return true
		}
		if m[0][i] != PointValueEmpty && m[0][i] == m[1][i] && m[1][i] == m[2][i] {
			g.Winner = m[0][i]
			return true
		}
	}
	if m[0][0] != PointValueEmpty && m[0][0] == m[1][1] && m[1][1] == m[2][2] {
		g.Winner = m[0][0]
		return true
	}
	if m[0][2] != PointValueEmpty && m[0][2] == m[1][1] && m[1][1] == m[2][0] {
		g.Winner = m[0][2]
		return true
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if m[i][j] == PointValueEmpty {
				return false // ещё можно играть
			}
		}
	}

	g.Winner = GameStatusDraw
	return true
}

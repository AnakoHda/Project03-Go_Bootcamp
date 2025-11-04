package domain

import "errors"

const (
	PointValueEmpty = 0
	PointValueO     = 1
	PointValueX     = 2
)

type Board struct {
	matrix [3][3]int `default:"{}"`
}

func NewBoard(m [3][3]int) Board {
	return Board{
		matrix: m,
	}
}
func (b *Board) Matrix() [3][3]int {
	return b.matrix
}
func (g *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.matrix[i][j] == PointValueEmpty {
				return false
			}
		}
	}
	return true
}
func (b *Board) CheckWinner(player int) (bool, error) {
	if player != PointValueO && player != PointValueX {
		return false, errors.New("invalid player")
	}
	m := b.matrix
	for i := 0; i < 3; i++ {
		if m[i][0] == player && m[i][1] == player && m[i][2] == player ||
			m[0][i] == player && m[1][i] == player && m[2][i] == player {
			return true, nil
		}
	}
	if m[0][0] == player && m[1][1] == player && m[2][2] == player ||
		m[0][2] == player && m[1][1] == player && m[2][0] == player {
		return true, nil
	}
	return false, nil
}
func (b *Board) GetPoint(i, j int) int {
	return b.matrix[i][j]
}
func (b *Board) SetPoint(i, j, point int) error {
	if point != PointValueX && point != PointValueO && point != PointValueEmpty {
		return errors.New("invalid point")
	}
	if point == PointValueEmpty {
		b.matrix[i][j] = point
		return nil
	}
	if b.matrix[i][j] == PointValueX || b.matrix[i][j] == PointValueO {
		return errors.New("point already occupied")
	}
	b.matrix[i][j] = point
	return nil
}

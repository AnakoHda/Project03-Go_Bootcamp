package minimax

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"
	"errors"
)

type Repository interface {
	Get(ctx context.Context, id string) (domain.Game, error)
	Save(ctx context.Context, g domain.Game, id string) error
}
type Service struct {
	repo Repository
}

func New(repository Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) WriteNewState(ctx context.Context, prevId string, next domain.Game) (domain.Game, error) {
	game, err := s.repo.Get(ctx, prevId)
	if err != nil {
		return domain.Game{}, err
	}
	err = game.ValidateNextState(next)
	if err != nil {
		return domain.Game{}, errors.New("ValidateNextState " + err.Error())
	}
	game = next
	_, _ = game.SetGameWinner()
	if err = s.repo.Save(ctx, game, prevId); err != nil {
		return domain.Game{}, err
	}
	return game, nil
}
func (s *Service) NextMove(ctx context.Context, id string) (domain.Game, error) {
	g, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.Game{}, err
	}

	if g.GetWinner() != domain.NoneWinner {
		return domain.Game{}, errors.New("game over, winner ")
	}

	if g.GetTurn() != domain.TurnO {
		return domain.Game{}, errors.New("turn X")
	}

	bestScore := -999
	var move [2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if err := g.Board().SetPoint(i, j, domain.PointValueO); err == nil {
				score := minimax(&g, 0, false)
				_ = g.Board().SetPoint(i, j, domain.PointValueEmpty)
				if score > bestScore {
					bestScore = score
					move = [2]int{i, j}
				}
			}
		}
	}
	_ = g.Board().SetPoint(move[0], move[1], domain.PointValueO)
	g.NextTour()
	_, _ = g.SetGameWinner()
	if err = s.repo.Save(ctx, g, id); err != nil {
		return domain.Game{}, err
	}
	return g, nil
}

func minimax(g *domain.Game, depth int, isMaximizing bool) int {
	if a, _ := g.Board().CheckWinner(domain.PointValueO); a {
		return 10 - depth
	}
	if a, _ := g.Board().CheckWinner(domain.PointValueX); a {
		return depth - 10
	}
	if a := g.Board().IsFull(); a {
		return 0
	}

	if isMaximizing {
		best := -999
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if g.Board().GetPoint(i, j) == domain.PointValueEmpty {
					_ = g.Board().SetPoint(i, j, domain.PointValueO)
					score := minimax(g, depth+1, false)
					_ = g.Board().SetPoint(i, j, domain.PointValueEmpty)
					if score > best {
						best = score
					}
				}
			}
		}
		return best
	} else {
		best := 999
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if g.Board().GetPoint(i, j) == domain.PointValueEmpty {
					_ = g.Board().SetPoint(i, j, domain.PointValueX)
					score := minimax(g, depth+1, true)
					_ = g.Board().SetPoint(i, j, domain.PointValueEmpty)
					if best > score {
						best = score
					}
				}
			}
		}
		return best
	}
}

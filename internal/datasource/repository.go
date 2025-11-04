package datasource

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"
	"errors"
	"sync"
)

type MainRepo struct {
	data sync.Map
}

func New() *MainRepo {
	return &MainRepo{
		data: sync.Map{},
	}
}

func (m *MainRepo) Save(ctx context.Context, g domain.Game, id string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		game := ToDS(g)
		m.data.Store(id, game)
		return nil
	}
}

func (m *MainRepo) Get(ctx context.Context, id string) (domain.Game, error) {
	select {
	case <-ctx.Done():
		return domain.Game{}, ctx.Err()
	default:
		value, ok := m.data.Load(id)
		if !ok {
			return domain.Game{}, errors.New("game not found" + id)
		}
		return ToDomain(value.(GameDS)), nil
	}
}

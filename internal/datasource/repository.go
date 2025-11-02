package datasource

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"

	"errors"
)

type MainRepo struct {
	store *Store
}

func NewRepository(store *Store) *MainRepo {
	return &MainRepo{store: store}
}

func (m *MainRepo) Save(ctx context.Context, g domain.Game) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		game, id := ToDS(g)
		m.store.data.Store(id, game)
		return nil
	}
}
func (m *MainRepo) Get(ctx context.Context, id string) (domain.Game, error) {
	select {
	case <-ctx.Done():
		return domain.Game{}, ctx.Err() // отмена или timeout
	default:
		value, ok := m.store.data.Load(id)
		if !ok {
			return domain.Game{}, errors.New("game not found")
		}
		return ToDomain(value.(GameDS), id), nil
	}
}

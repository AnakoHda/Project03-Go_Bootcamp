package datasource

import "sync"

type Store struct {
	data sync.Map
}

func NewStore() *Store { return &Store{} }

func (s *Store) Put(id string, e GameDS) {
	s.data.Store(id, e)
}

func (s *Store) Load(id string) (GameDS, bool) {
	v, ok := s.data.Load(id)
	if !ok {
		return GameDS{}, false
	}
	return v.(GameDS), true
}

func (s *Store) Delete(id string) {
	s.data.Delete(id)
}

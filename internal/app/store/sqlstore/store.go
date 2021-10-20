package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"http-rest-api/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository

	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

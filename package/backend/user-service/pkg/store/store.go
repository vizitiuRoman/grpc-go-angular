package store

import (
	"github.com/user-service/pkg/store/pg"
)

type Store struct {
	PG *pg.DB

	User UserRepo
}

func NewStore() (*Store, error) {
	pgDB, err := pg.Dial()
	if err != nil {
		return nil, err
	}
	return &Store{
		PG:   pgDB,
		User: pg.NewUserRepo(pgDB),
	}, nil
}

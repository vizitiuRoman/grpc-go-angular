package store

import "github.com/user-service/pkg/store/pg"

type Store struct {
	User UserRepo
}

func NewStore() (*Store, error) {
	pgDB, err := pg.Dial()
	if err != nil {
		return nil, err
	}
	return &Store{
		User: pg.NewUserRepo(pgDB),
	}, nil
}

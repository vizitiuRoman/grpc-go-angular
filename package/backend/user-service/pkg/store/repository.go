package store

import (
	. "github.com/user-service/pkg/domain"
)

// UserRepo is a store for users
//go:generate mockery --dir . --name UserRepo --output ./mocks
type UserRepo interface {
	CreateUser(*User) (*User, error)
	GetUser(uint64) (*User, error)
	UpdateUser(*User) (*User, error)
	DeleteUser(uint64) error
	GetUserByEmail(string) (*User, error)
}

package services

import (
	. "github.com/user-service/pkg/domain"
)

// UserService is a service for users
//go:generate mockery --dir . --name UserService --output ./mocks
type UserService interface {
	CreateUser(model *User) (*User, error)
	GetUser(uint64) (*User, error)
	UpdateUser(*User) (*User, error)
	DeleteUser(uint64) error
	GetUserByEmail(string) (*User, error)
	VerifyUser(string, string) (*User, error)
}

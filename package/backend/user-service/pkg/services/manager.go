package services

import (
	"github.com/user-service/pkg/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	User UserService
}

// NewManager creates new service manager
func NewManager(store *store.Store) *Manager {
	return &Manager{
		User: NewUserService(store.User),
	}
}

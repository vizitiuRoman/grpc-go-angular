package services

import (
	"fmt"

	"github.com/pkg/errors"
	. "github.com/user-service/pkg/models"
	"github.com/user-service/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

type UserWebService struct {
	store *store.Store
}

// NewUserWebService creates a new user web service
func NewUserWebService(store *store.Store) *UserWebService {
	return &UserWebService{
		store: store,
	}
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (svc *UserWebService) CreateUser(user *User) (*User, error) {
	password, err := hashPassword(user.Password)
	if err != nil {
		return &User{}, errors.New("svc.user.CreateUser hashPassword error")
	}
	user.Password = string(password)
	createdUser, err := svc.store.User.CreateUser(user)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.CreateUser error")
	}
	return createdUser, err
}

func (svc *UserWebService) GetUser(userID uint64) (*User, error) {
	user, err := svc.store.User.GetUser(userID)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUser error")
	}
	return user, nil
}

func (svc *UserWebService) UpdateUser(user *User) error {
	foundUser, err := svc.store.User.GetUser(user.ID)
	if err != nil {
		return errors.Wrap(err, "svc.user.GetUser error")
	}
	if foundUser == nil {
		return errors.New(fmt.Sprintf("User '%d' not found", user.ID))
	}

	err = VerifyPassword(foundUser.Password, user.Password)
	if err != nil {
		return errors.Wrap(err, "Incorrect password")
	}

	password, err := hashPassword(user.Password)
	if err != nil {
		return errors.New("svc.user.UpdateUser hashPassword error")
	}

	user.Password = string(password)
	err = svc.store.User.UpdateUser(user)
	if err != nil {
		return errors.Wrap(err, "svc.user.UpdateUser error")
	}

	return nil
}

func (svc *UserWebService) DeleteUser(userID uint64) error {
	foundUser, err := svc.store.User.GetUser(userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.GetUser error")
	}
	if foundUser == nil {
		return errors.New(fmt.Sprintf("User '%d' not found", userID))
	}
	err = svc.store.User.DeleteUser(userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.DeleteUser error")
	}
	return nil
}

func (svc *UserWebService) GetUserByEmail(email string) (*User, error) {
	user, err := svc.store.User.GetUserByEmail(email)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUserByEmail error")
	}
	return user, nil
}

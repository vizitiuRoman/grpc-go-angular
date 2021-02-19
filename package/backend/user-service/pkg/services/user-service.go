package services

import (
	"fmt"

	"github.com/pkg/errors"
	. "github.com/user-service/pkg/domain"
	"github.com/user-service/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

type userWebService struct {
	store *store.Store
}

func NewUserWebService(store *store.Store) UserService {
	return &userWebService{
		store: store,
	}
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (svc *userWebService) CreateUser(user *User) (*User, error) {
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

func (svc *userWebService) GetUser(userID uint64) (*User, error) {
	user, err := svc.store.User.GetUser(userID)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUser error")
	}
	return user, nil
}

func (svc *userWebService) UpdateUser(user *User) (*User, error) {
	foundUser, err := svc.store.User.GetUser(user.ID)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUser error")
	}
	if foundUser == nil {
		return &User{}, errors.New(fmt.Sprintf("User '%d' not found", user.ID))
	}

	err = verifyPassword(foundUser.Password, user.Password)
	if err != nil {
		return &User{}, errors.Wrap(err, "Incorrect password")
	}

	password, err := hashPassword(user.Password)
	if err != nil {
		return &User{}, errors.New("svc.user.UpdateUser hashPassword error")
	}

	user.Password = string(password)
	updatedUser, err := svc.store.User.UpdateUser(user)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.UpdateUser error")
	}

	return updatedUser, nil
}

func (svc *userWebService) DeleteUser(userID uint64) error {
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

func (svc *userWebService) GetUserByEmail(email string) (*User, error) {
	user, err := svc.store.User.GetUserByEmail(email)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUserByEmail error")
	}
	return user, nil
}

func (svc *userWebService) VerifyUser(email string, password string) (*User, error) {
	foundUser, err := svc.store.User.GetUserByEmail(email)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUserByEmail error")
	}

	err = verifyPassword(foundUser.Password, password)
	if err != nil {
		return &User{}, errors.Wrap(err, "verify user error")
	}

	return foundUser, nil
}

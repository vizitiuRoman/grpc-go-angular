package services

import (
	"fmt"

	"github.com/pkg/errors"
	. "github.com/user-service/pkg/domain"
	"github.com/user-service/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo store.UserRepo
}

func NewUserService(repo store.UserRepo) UserService {
	return &userService{repo}
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (svc *userService) CreateUser(user *User) (*User, error) {
	password, err := hashPassword(user.Password)
	if err != nil {
		return &User{}, errors.New("svc.user.CreateUser hashPassword error")
	}
	user.Password = string(password)
	createdUser, err := svc.repo.CreateUser(user)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.CreateUser error")
	}
	return createdUser, err
}

func (svc *userService) GetUser(userID uint64) (*User, error) {
	user, err := svc.repo.GetUser(userID)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUser error")
	}
	return user, nil
}

func (svc *userService) UpdateUser(user *User) (*User, error) {
	foundUser, err := svc.repo.GetUser(user.ID)
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
	updatedUser, err := svc.repo.UpdateUser(user)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.UpdateUser error")
	}

	return updatedUser, nil
}

func (svc *userService) DeleteUser(userID uint64) error {
	foundUser, err := svc.repo.GetUser(userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.GetUser error")
	}
	if foundUser == nil {
		return errors.New(fmt.Sprintf("User '%d' not found", userID))
	}
	err = svc.repo.DeleteUser(userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.DeleteUser error")
	}
	return nil
}

func (svc *userService) GetUserByEmail(email string) (*User, error) {
	user, err := svc.repo.GetUserByEmail(email)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUserByEmail error")
	}
	return user, nil
}

func (svc *userService) VerifyUser(email string, password string) (*User, error) {
	foundUser, err := svc.repo.GetUserByEmail(email)
	if err != nil {
		return &User{}, errors.Wrap(err, "svc.user.GetUserByEmail error")
	}

	err = verifyPassword(foundUser.Password, password)
	if err != nil {
		return &User{}, errors.Wrap(err, "verify user error")
	}

	return foundUser, nil
}

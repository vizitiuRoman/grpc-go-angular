package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/user-service/pkg/domain"
	"github.com/user-service/pkg/store"
	"github.com/user-service/pkg/store/mocks"
)

func TestCreateUser(t *testing.T) {
	input := &User{
		Email:    "roma@roma",
		Password: "qwertyui",
	}

	tests := []struct {
		name         string
		expectations func(userRepo *mocks.UserRepo)
		input        *User
		err          error
	}{
		{
			name: "Valid and create",
			expectations: func(userRepo *mocks.UserRepo) {
				userRepo.On("CreateUser", input).Return(nil, nil)
			},
			input: input,
		},
		{
			name: "Valid and create",
			expectations: func(userRepo *mocks.UserRepo) {
				userRepo.On("CreateUser", input).Return(input, nil)
			},
			input: input,
		},
		{
			name: "Store error",
			expectations: func(userRepo *mocks.UserRepo) {
				userRepo.On("CreateUser", input).Return(
					nil, errors.New("some error"),
				)
			},
			input: input,
			err:   errors.New("svc.user.CreateUser error: some error"),
		},
	}
	for _, test := range tests {
		t.Logf("Running: %s", test.name)

		userRepo := &mocks.UserRepo{}
		svc := NewUserService(&store.Store{User: userRepo})
		test.expectations(userRepo)

		_, err := svc.CreateUser(test.input)
		if err != nil {
			if test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		}
		userRepo.AssertExpectations(t)
	}
}

func TestGetUser(t *testing.T) {
	input := &User{
		ID:       1,
		Email:    "roma@roma",
		Password: "qwertyui",
	}

	tests := []struct {
		name         string
		expectations func(userRepo *mocks.UserRepo)
		input        *User
		err          error
	}{
		{
			name: "Valid and found",
			expectations: func(userRepo *mocks.UserRepo) {
				userRepo.On("GetUser", input.ID).Return(input, nil)
			},
			input: input,
		},
		{
			name: "Valid user id and not found",
			expectations: func(userRepo *mocks.UserRepo) {
				userRepo.On("GetUser", input.ID).Return(nil, nil)
			},
			input: input,
			err:   errors.New("User 1 not found: resource not found"),
		},
		{
			name: "Store error",
			expectations: func(userRepo *mocks.UserRepo) {
				userRepo.On("GetUser", input.ID).Return(
					nil, errors.New("some error"),
				)
			},
			input: input,
			err:   errors.New("svc.user.GetUser error: some error"),
		},
	}
	for _, test := range tests {
		t.Logf("Running: %s", test.name)

		userRepo := &mocks.UserRepo{}
		svc := NewUserService(&store.Store{User: userRepo})
		test.expectations(userRepo)

		_, err := svc.GetUser(test.input.ID)
		if err != nil {
			if test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		}
		userRepo.AssertExpectations(t)
	}
}

func TestUpdateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}

func TestGetUserByEmail(t *testing.T) {

}

func TestVerifyUser(t *testing.T) {

}

package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_service "go.mod/mocks"
	"go.mod/model"
	"testing"
)

func TestServiceGetAllUsers(t *testing.T) {
	storeReturn := map[string]*model.User{
		"Burcu": {
			Username: "Burcu",
			Balance:  55,
		},
		"Sevinç": {
			Username: "Sevinç",
			Balance:  66,
		},
	}

	store := mock_service.NewMockIUserLocalStorage(gomock.NewController(t))
	store.EXPECT().
		GetAllUsers().
		Return(storeReturn, nil).
		Times(1)

	service := NewService(store, 0, -100)
	users, err := service.GetAllUsers()

	expectedUsers := make([]*model.User, 0)
	expectedUsers = append(expectedUsers, &model.User{
		Username: "Burcu",
		Balance:  55,
	}, &model.User{
		Username: "Sevinç",
		Balance:  66,
	})

	assert.Equal(t, &expectedUsers, &users)
	assert.Nil(t, err)
}

func TestServiceRepoReturnError(t *testing.T) {
	serviceError := errors.New("test error")
	store := mock_service.NewMockIUserLocalStorage(gomock.NewController(t))
	store.EXPECT().
		GetAllUsers().
		Return(nil, serviceError).
		Times(1)

	service := NewService(store, 0, -100)
	users, err := service.GetAllUsers()

	assert.ErrorIs(t, err, serviceError)
	assert.Nil(t, users)
}

func TestServiceGetUser(t *testing.T) {
	storeReturn := &model.User{
		Username: "Burcu",
		Balance:  88,
	}

	store := mock_service.NewMockIUserLocalStorage(gomock.NewController(t))
	store.EXPECT().
		GetUser("Burcu").
		Return(storeReturn, nil).
		Times(1)

	service := NewService(store, initialBalanceAmount, minimumBalanceAmount)
	users, err := service.GetUser("Burcu")

	assert.Equal(t, storeReturn, users)
	assert.Nil(t, err)
}

func TestServiceCreateUser(t *testing.T) {
	expecUser := model.User{
		Username: "Burcu",
		Balance:  initialBalanceAmount,
	}

	store := mock_service.NewMockIUserLocalStorage(gomock.NewController(t))
	store.EXPECT().
		GetUser("Burcu").
		Return(nil, errors.New("Burxu")).
		Times(1)
	store.EXPECT().
		CreateUser(&expecUser).
		Return(&expecUser).
		Times(1)

	service := NewService(store, initialBalanceAmount, minimumBalanceAmount)
	newUser := service.CreateUser("Burcu")

	assert.Equal(t, &expecUser, newUser)
}

func TestServiceUpdateUser(t *testing.T) {
	User := model.User{
		Username: "Burcu",
		Balance:  111,
	}

	expecUser := model.User{
		Username: "Burcu",
		Balance:  333,
	}

	store := mock_service.NewMockIUserLocalStorage(gomock.NewController(t))
	store.EXPECT().
		GetUser(User.Username).
		Return(&User, nil).
		Times(1)
	store.EXPECT().
		UpdateUser(&expecUser).
		Return(&expecUser).
		Times(1)

	service := NewService(store, initialBalanceAmount, minimumBalanceAmount)

	updatedUser, err := service.UpdateUser(expecUser.Username, 222)

	assert.Equal(t, &expecUser, updatedUser)
	assert.Nil(t, err)
}

func TestServiceUpdateUserMinimumBalanceCheck(t *testing.T) {
	User := model.User{
		Username: "Burcu",
		Balance:  111,
	}

	expecUser := model.User{
		Username: "Burcu",
		Balance:  333,
	}

	store := mock_service.NewMockIUserLocalStorage(gomock.NewController(t))
	store.EXPECT().
		GetUser(User.Username).
		Return(&User, nil).
		Times(1)

	service := NewService(store, initialBalanceAmount, minimumBalanceAmount)

	updatedUser, err := service.UpdateUser(expecUser.Username, -444)

	assert.NotNil(t, err)
	assert.Nil(t, updatedUser)
}

const (
	initialBalanceAmount = 0
	minimumBalanceAmount = -100
)

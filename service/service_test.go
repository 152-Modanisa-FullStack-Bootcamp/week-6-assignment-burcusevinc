package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_service "go.mod/mocks"
	"go.mod/model"
	"testing"
)

func TestServiceReturnAllUsers(t *testing.T) {
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
		Return(storeReturn).
		Times(1)

	service := NewService(store, 0, -100)
	users := service.GetAllUsers()

	expectedUsers := make([]*model.User, 0)
	expectedUsers = append(expectedUsers, &model.User{
		Username: "Burcu",
		Balance:  55,
	}, &model.User{
		Username: "Sevinç",
		Balance:  66,
	})

	assert.Equal(t, &expectedUsers, &users)
}

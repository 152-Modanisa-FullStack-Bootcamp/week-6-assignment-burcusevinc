package repository

import (
	"github.com/stretchr/testify/assert"
	"go.mod/model"
	"testing"
)

func TestRepositoryGetAllUsers(t *testing.T) {
	store := NewUserLocalStorage()

	length := len(store.GetAllUsers())

	assert.Equal(t, 0, length)
}

func TestRepositoryGetUser(t *testing.T) {
	store := NewUserLocalStorage()

	User := model.User{
		Username: "Burcu",
		Balance:  88,
	}

	store.CreateUser(&User)

	user, err := store.GetUser(User.Username)

	assert.Equal(t, &User, user)
	assert.Nil(t, err)
}

func TestRepositoryGetUserExist(t *testing.T) {
	store := NewUserLocalStorage()

	user, err := store.GetUser("none")

	assert.Equal(t, "User is not found!", err.Error())
	assert.Nil(t, user)
}

func TestRepositoryCreateUser(t *testing.T) {
	store := NewUserLocalStorage()

	User := model.User{
		Username: "Burcu",
		Balance:  88,
	}

	users := store.GetAllUsers()

	length := len(users)
	store.CreateUser(&User)

	assert.Equal(t, length+1, len(users))
}

func TestRepositoryUpdateUser(t *testing.T) {
	store := NewUserLocalStorage()

	User := model.User{
		Username: "Burcu",
		Balance:  88,
	}

	store.CreateUser(&User)

	UpdatedUser := model.User{
		Username: "Burcu",
		Balance:  99,
	}

	store.UpdateUser(&UpdatedUser)

	user, _ := store.GetUser(UpdatedUser.Username)

	assert.Equal(t, &UpdatedUser, user)
}

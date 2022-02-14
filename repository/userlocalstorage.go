package repository

import (
	"errors"
	"go.mod/model"
)

//Repository package interface
type IUserLocalStorage interface {
	GetAllUsers() (map[string]*model.User, error)
	GetUser(username string) (*model.User, error)
	CreateUser(user *model.User) *model.User
	UpdateUser(user *model.User) *model.User
}

//Repository package struct implements model User struct map
type UserLocalStorage struct {
	Users map[string]*model.User
}

//it returns all user data
func (u *UserLocalStorage) GetAllUsers() (map[string]*model.User, error) {
	users := u.Users
	return users, nil
}

//it returns wanted user or an error
func (u *UserLocalStorage) GetUser(username string) (*model.User, error) {
	for _, user := range u.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("User is not found!")
}

//it returns username
func (u *UserLocalStorage) CreateUser(user *model.User) *model.User {
	u.Users[user.Username] = &model.User{
		Username: user.Username,
		Balance:  user.Balance,
	}
	return u.Users[user.Username]
}

//it returns username
func (u *UserLocalStorage) UpdateUser(user *model.User) *model.User {
	u.Users[user.Username] = &model.User{
		Username: user.Username,
		Balance:  user.Balance,
	}
	return u.Users[user.Username]
}

//constructure function
func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		Users: make(map[string]*model.User),
	}
}

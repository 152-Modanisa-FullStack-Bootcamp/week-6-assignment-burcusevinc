package service

import (
	"errors"
	"fmt"
	"go.mod/model"
	"go.mod/repository"
)

//Service package interface
type IUserService interface {
	GetAllUsers() []*model.User
	GetUser(username string) (*model.User, error)
	CreateUser(username string) *model.User
	UpdateUser(username string, balance int) (*model.User, error)
}

//Service package struct implements store package's interface and integer values.
type UserService struct {
	store                repository.IUserLocalStorage
	InitialBalanceAmount int
	MinimumBalanceAmount int
}

func (u *UserService) GetAllUsers() []*model.User {
	//using repository's function, it returns all users data
	getUsers := u.store.GetAllUsers()

	//list of users
	users := make([]*model.User, 0, len(getUsers))
	//making new users struct
	for _, user := range getUsers {
		users = append(users, user)
	}
	return users
}

func (u *UserService) GetUser(username string) (*model.User, error) {
	//using repository's function, it returns username data
	user, err := u.store.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserService) CreateUser(username string) *model.User {
	//it returns the wanted user
	existUser, err := u.store.GetUser(username)
	if err == nil {
		return existUser
	}
	//create new user struct, balance amount is, initialbalanceamount
	user := &model.User{
		Username: username,
		Balance:  u.InitialBalanceAmount,
	}

	return u.store.CreateUser(user)
}

func (u *UserService) UpdateUser(username string, balance int) (*model.User, error) {
	//using repository's function, it returns user data
	getUsers, err := u.GetUser(username)
	if err != nil {
		return nil, err
	}

	//initialize result value, add balance value with user value.
	result := getUsers.Balance + (balance)

	//if result, smaller than minimum balance amount, return an error
	if result < u.MinimumBalanceAmount {
		return nil, errors.New(fmt.Sprintf("Wallet balance amount never belows %d", u.MinimumBalanceAmount))
	}
	//create new user struct
	user := &model.User{
		Username: username,
		Balance:  balance,
	}

	//using repository's function, it returns user data
	updatedUser := u.store.UpdateUser(user)
	return updatedUser, nil
}

//constructure function
func NewService(store repository.IUserLocalStorage, iba int, mba int) IUserService {
	return &UserService{
		store:                store,
		InitialBalanceAmount: iba,
		MinimumBalanceAmount: mba,
	}
}

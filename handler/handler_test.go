package handler

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_service "go.mod/mocks"
	"go.mod/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	initialBalanceAmount = 0
	minimumBalanceAmount = -100
)

func TestHandlerServiceGetAllUsers(t *testing.T) {
	service := mock_service.NewMockIUserService(gomock.NewController(t))
	serviceReturn := []*model.User{
		{
			Username: "",
			Balance:  0,
		},
	}
	service.EXPECT().
		GetAllUsers().
		Return(serviceReturn, nil).
		Times(1)

	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAllUsers(w, r)

	actual := []*model.User{{
		Username: "",
		Balance:  0,
	}}
	json.Unmarshal(w.Body.Bytes(), actual)

	assert.Equal(t, actual, serviceReturn)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestHandlerServiceGetAllUserReturnError(t *testing.T) {
	service := mock_service.NewMockIUserService(gomock.NewController(t))

	serviceError := errors.New("test error")

	service.EXPECT().
		GetAllUsers().
		Return(nil, serviceError).
		Times(1)

	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAllUsers(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, w.Body.String(), serviceError.Error())
}

func TestHandlerServiceGetUser(t *testing.T) {
	service := mock_service.NewMockIUserService(gomock.NewController(t))
	serviceReturn := &model.User{
		Username: "Burcu",
		Balance:  333,
	}

	service.EXPECT().
		GetUser("").
		Return(serviceReturn, nil).
		Times(1)

	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetUser(w, r)

	actual := &model.User{
		Username: "Burcu",
		Balance:  333,
	}

	json.Unmarshal(w.Body.Bytes(), actual)

	assert.Equal(t, serviceReturn, actual)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestHandlerServiceGetUserReturnError(t *testing.T) {
	service := mock_service.NewMockIUserService(gomock.NewController(t))

	serviceError := errors.New("null")

	service.EXPECT().
		GetUser("").
		Return(nil, serviceError).
		Times(1)

	handler := NewWalletHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetUser(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusNotFound)
	assert.Equal(t, w.Body.String(), serviceError.Error())
}

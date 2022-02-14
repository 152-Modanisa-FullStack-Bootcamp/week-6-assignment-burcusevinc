// Code generated by MockGen. DO NOT EDIT.
// Source: ./handler/handler.go

// Package mock_handler is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWalletHandler is a mock of IWalletHandler interface.
type MockIWalletHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletHandlerMockRecorder
}

// MockIWalletHandlerMockRecorder is the mock recorder for MockIWalletHandler.
type MockIWalletHandlerMockRecorder struct {
	mock *MockIWalletHandler
}

// NewMockIWalletHandler creates a new mock instance.
func NewMockIWalletHandler(ctrl *gomock.Controller) *MockIWalletHandler {
	mock := &MockIWalletHandler{ctrl: ctrl}
	mock.recorder = &MockIWalletHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWalletHandler) EXPECT() *MockIWalletHandlerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIWalletHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateUser", w, r)
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIWalletHandlerMockRecorder) CreateUser(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIWalletHandler)(nil).CreateUser), w, r)
}

// GetAllUsers mocks base method.
func (m *MockIWalletHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllUsers", w, r)
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIWalletHandlerMockRecorder) GetAllUsers(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIWalletHandler)(nil).GetAllUsers), w, r)
}

// GetUser mocks base method.
func (m *MockIWalletHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUser", w, r)
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIWalletHandlerMockRecorder) GetUser(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIWalletHandler)(nil).GetUser), w, r)
}

// UpdateUser mocks base method.
func (m *MockIWalletHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUser", w, r)
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIWalletHandlerMockRecorder) UpdateUser(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIWalletHandler)(nil).UpdateUser), w, r)
}

// WalletMethods mocks base method.
func (m *MockIWalletHandler) WalletMethods(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WalletMethods", w, r)
}

// WalletMethods indicates an expected call of WalletMethods.
func (mr *MockIWalletHandlerMockRecorder) WalletMethods(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletMethods", reflect.TypeOf((*MockIWalletHandler)(nil).WalletMethods), w, r)
}
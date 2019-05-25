// Code generated by MockGen. DO NOT EDIT.
// Source: logic/usersService.go

// Package mock_logic is a generated GoMock package.
package mock_logic

import (
	gomock "github.com/golang/mock/gomock"
	errs "github.com/studtool/common/errs"
	models "github.com/studtool/documents-service/models"
	reflect "reflect"
)

// MockUsersService is a mock of UsersService interface
type MockUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceMockRecorder
}

// MockUsersServiceMockRecorder is the mock recorder for MockUsersService
type MockUsersServiceMockRecorder struct {
	mock *MockUsersService
}

// NewMockUsersService creates a new mock instance
func NewMockUsersService(ctrl *gomock.Controller) *MockUsersService {
	mock := &MockUsersService{ctrl: ctrl}
	mock.recorder = &MockUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersService) EXPECT() *MockUsersServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockUsersService) AddUser(u *models.User) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", u)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockUsersServiceMockRecorder) AddUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUsersService)(nil).AddUser), u)
}

// CheckUserExists mocks base method
func (m *MockUsersService) CheckUserExists(u *models.User) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExists", u)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// CheckUserExists indicates an expected call of CheckUserExists
func (mr *MockUsersServiceMockRecorder) CheckUserExists(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExists", reflect.TypeOf((*MockUsersService)(nil).CheckUserExists), u)
}

// DeleteUser mocks base method
func (m *MockUsersService) DeleteUser(u *models.User) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", u)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUsersServiceMockRecorder) DeleteUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUsersService)(nil).DeleteUser), u)
}
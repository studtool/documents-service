// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/usersRepository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	gomock "github.com/golang/mock/gomock"
	errs "github.com/studtool/common/errs"
	types "github.com/studtool/common/types"
	models "github.com/studtool/documents-service/models"
	reflect "reflect"
)

// MockUsersRepository is a mock of UsersRepository interface
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockUsersRepository) AddUser(u *models.User) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", u)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockUsersRepositoryMockRecorder) AddUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUsersRepository)(nil).AddUser), u)
}

// CheckUserExistsByID mocks base method
func (m *MockUsersRepository) CheckUserExistsByID(userID types.ID) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExistsByID", userID)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// CheckUserExistsByID indicates an expected call of CheckUserExistsByID
func (mr *MockUsersRepositoryMockRecorder) CheckUserExistsByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExistsByID", reflect.TypeOf((*MockUsersRepository)(nil).CheckUserExistsByID), userID)
}

// DeleteUserByID mocks base method
func (m *MockUsersRepository) DeleteUserByID(userID types.ID) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByID", userID)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// DeleteUserByID indicates an expected call of DeleteUserByID
func (mr *MockUsersRepositoryMockRecorder) DeleteUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockUsersRepository)(nil).DeleteUserByID), userID)
}

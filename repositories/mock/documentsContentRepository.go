// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/documentsContentRepository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	gomock "github.com/golang/mock/gomock"
	errs "github.com/studtool/common/errs"
	types "github.com/studtool/common/types"
	models "github.com/studtool/documents-service/models"
	reflect "reflect"
)

// MockDocumentsContentRepository is a mock of DocumentsContentRepository interface
type MockDocumentsContentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDocumentsContentRepositoryMockRecorder
}

// MockDocumentsContentRepositoryMockRecorder is the mock recorder for MockDocumentsContentRepository
type MockDocumentsContentRepositoryMockRecorder struct {
	mock *MockDocumentsContentRepository
}

// NewMockDocumentsContentRepository creates a new mock instance
func NewMockDocumentsContentRepository(ctrl *gomock.Controller) *MockDocumentsContentRepository {
	mock := &MockDocumentsContentRepository{ctrl: ctrl}
	mock.recorder = &MockDocumentsContentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocumentsContentRepository) EXPECT() *MockDocumentsContentRepositoryMockRecorder {
	return m.recorder
}

// GetDocumentContent mocks base method
func (m *MockDocumentsContentRepository) GetDocumentContent(documentID types.ID, content *models.DocumentContent) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentContent", documentID, content)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// GetDocumentContent indicates an expected call of GetDocumentContent
func (mr *MockDocumentsContentRepositoryMockRecorder) GetDocumentContent(documentID, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentContent", reflect.TypeOf((*MockDocumentsContentRepository)(nil).GetDocumentContent), documentID, content)
}

// UpdateDocumentContent mocks base method
func (m *MockDocumentsContentRepository) UpdateDocumentContent(documentID types.ID, content *models.DocumentContent) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDocumentContent", documentID, content)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// UpdateDocumentContent indicates an expected call of UpdateDocumentContent
func (mr *MockDocumentsContentRepositoryMockRecorder) UpdateDocumentContent(documentID, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDocumentContent", reflect.TypeOf((*MockDocumentsContentRepository)(nil).UpdateDocumentContent), documentID, content)
}
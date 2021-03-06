// Code generated by MockGen. DO NOT EDIT.
// Source: logic/documentsContentService.go

// Package mock_logic is a generated GoMock package.
package mock_logic

import (
	gomock "github.com/golang/mock/gomock"
	errs "github.com/studtool/common/errs"
	logic "github.com/studtool/documents-service/logic"
	reflect "reflect"
)

// MockDocumentsContentService is a mock of DocumentsContentService interface
type MockDocumentsContentService struct {
	ctrl     *gomock.Controller
	recorder *MockDocumentsContentServiceMockRecorder
}

// MockDocumentsContentServiceMockRecorder is the mock recorder for MockDocumentsContentService
type MockDocumentsContentServiceMockRecorder struct {
	mock *MockDocumentsContentService
}

// NewMockDocumentsContentService creates a new mock instance
func NewMockDocumentsContentService(ctrl *gomock.Controller) *MockDocumentsContentService {
	mock := &MockDocumentsContentService{ctrl: ctrl}
	mock.recorder = &MockDocumentsContentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocumentsContentService) EXPECT() *MockDocumentsContentServiceMockRecorder {
	return m.recorder
}

// GetDocumentContent mocks base method
func (m *MockDocumentsContentService) GetDocumentContent(params *logic.GetDocumentContentParams) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentContent", params)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// GetDocumentContent indicates an expected call of GetDocumentContent
func (mr *MockDocumentsContentServiceMockRecorder) GetDocumentContent(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentContent", reflect.TypeOf((*MockDocumentsContentService)(nil).GetDocumentContent), params)
}

// UpdateDocumentContent mocks base method
func (m *MockDocumentsContentService) UpdateDocumentContent(params *logic.UpdateDocumentContentParams) *errs.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDocumentContent", params)
	ret0, _ := ret[0].(*errs.Error)
	return ret0
}

// UpdateDocumentContent indicates an expected call of UpdateDocumentContent
func (mr *MockDocumentsContentServiceMockRecorder) UpdateDocumentContent(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDocumentContent", reflect.TypeOf((*MockDocumentsContentService)(nil).UpdateDocumentContent), params)
}

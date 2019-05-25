package sfake

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/logic"
)

type DocumentsInfoService struct{}

func NewDocumentsInfoService() *DocumentsInfoService {
	return &DocumentsInfoService{}
}

func (s *DocumentsInfoService) AddDocumentInfo(params logic.AddDocumentInfoParams) *errs.Error {
	return errs.NewNotImplementedError(s.AddDocumentInfo)
}

func (s *DocumentsInfoService) GetDocumentInfo(params logic.GetDocumentInfoParams) *errs.Error {
	return errs.NewNotImplementedError(s.GetDocumentInfo)
}

func (s *DocumentsInfoService) UpdateDocumentTitle(params logic.UpdateDocumentTitleParams) *errs.Error {
	return errs.NewNotImplementedError(s.UpdateDocumentTitle)
}

func (s *DocumentsInfoService) UpdateDocumentSubject(params logic.UpdateDocumentSubjectParams) *errs.Error {
	return errs.NewNotImplementedError(s.UpdateDocumentSubject)
}

func (s *DocumentsInfoService) DeleteDocument(params logic.DeleteDocumentParams) *errs.Error {
	return errs.NewNotImplementedError(s.DeleteDocument)
}

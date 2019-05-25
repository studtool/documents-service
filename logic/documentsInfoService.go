package logic

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type DocumentsInfoService interface {
	AddDocumentInfo(params AddDocumentInfoParams) *errs.Error
	GetDocumentInfo(params GetDocumentInfoParams) *errs.Error
	UpdateDocumentTitle(params UpdateDocumentTitleParams) *errs.Error
	UpdateDocumentSubject(params UpdateDocumentSubjectParams) *errs.Error
	DeleteDocument(params DeleteDocumentParams) *errs.Error
}

type AddDocumentInfoParams struct {
	UserID   types.ID
	Document *models.DocumentInfo
}

type GetDocumentInfoParams struct {
	UserID   types.ID
	Document *models.DocumentInfo
}

type UpdateDocumentTitleParams struct {
	UserID types.ID
	Update *models.DocumentTitleUpdate
}

type UpdateDocumentSubjectParams struct {
	UserID types.ID
	Update *models.DocumentSubjectUpdate
}

type DeleteDocumentParams struct {
	UserID     types.ID
	DocumentID types.ID
}

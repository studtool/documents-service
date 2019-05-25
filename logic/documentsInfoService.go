package logic

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"
	"github.com/studtool/documents-service/repositories"

	"github.com/studtool/documents-service/models"
)

type DocumentsInfoService interface {
	AddDocumentInfo(params AddDocumentInfoParams) *errs.Error
	GetDocumentInfo(params GetDocumentInfoParams) *errs.Error
	GetDocumentsInfo(params GetDocumentsInfoParams) *errs.Error
	UpdateDocumentTitle(params UpdateDocumentTitleParams) *errs.Error
	UpdateDocumentSubject(params UpdateDocumentSubjectParams) *errs.Error
	DeleteDocument(params DeleteDocumentParams) *errs.Error
}

type AddDocumentInfoParams struct {
	UserID       types.ID
	DocumentInfo *models.DocumentInfo
}

type GetDocumentInfoParams struct {
	UserID       types.ID
	DocumentInfo *models.DocumentInfo
}

type GetDocumentsInfoParams struct {
	UserID        types.ID
	OwnerID       types.ID
	Subject       string
	Page          repositories.Page
	DocumentsInfo *models.DocumentsInfo
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

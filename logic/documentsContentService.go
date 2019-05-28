package logic

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type DocumentsContentService interface {
	GetDocumentContent(params *GetDocumentContentParams) *errs.Error
	UpdateDocumentContent(params *UpdateDocumentContentParams) *errs.Error
}

type GetDocumentContentParams struct {
	UserID          types.ID
	DocumentID      types.ID
	DocumentContent *models.DocumentContent
}

type UpdateDocumentContentParams struct {
	UserID          types.ID
	DocumentID      types.ID
	DocumentContent *models.DocumentContent
}

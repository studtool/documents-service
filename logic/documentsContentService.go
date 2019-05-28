package logic

import (
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type DocumentsContentService interface {
	GetDocumentContent()
	UpdateDocumentContent()
}

type GetDocumentContentParams struct {
	UserID          string
	DocumentContent *models.DocumentContent
}

type UpdateDocumentContentParams struct {
	UserID          types.ID
	DocumentID      types.ID
	DocumentContent *models.DocumentContent
}

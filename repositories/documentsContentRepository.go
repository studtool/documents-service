package repositories

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type DocumentsContentRepository interface {
	GetDocumentContent(documentID types.ID, content *models.DocumentContent) *errs.Error
	UpdateDocumentContent(documentID types.ID, content *models.DocumentContent) *errs.Error
}

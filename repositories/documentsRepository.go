package repositories

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
)

type ContentRepository interface {
	AddDocument(userID string, block *models.DocumentBlock) *errs.Error
	GetDocument(documentID string, userID string) (models.DocumentBlocks, *errs.Error)
	UpdateDocument(userID string, block *models.DocumentBlock) *errs.Error
	DeleteDocument(documentId string, userID string) *errs.Error
}

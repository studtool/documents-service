package repositories

import (
	"github.com/studtool/common/errs"
)

type DocumentsRepository interface {
	AddDocument(documentId string, data []byte) *errs.Error
	GetDocument(documentId string) ([]byte, *errs.Error)
	UpdateDocument(documentId string, data []byte) *errs.Error
	DeleteDocument(documentId string) *errs.Error
}

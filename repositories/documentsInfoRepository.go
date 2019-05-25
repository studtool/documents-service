package repositories

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type DocumentsInfoRepository interface {
	AddDocumentInfo(info *models.DocumentInfo) *errs.Error

	GetDocumentInfoByID(info *models.DocumentInfo) *errs.Error

	GetDocumentsInfoByOwnerID(ownerID types.ID, page Page) (models.DocumentsInfo, *errs.Error)
	GetDocumentsInfoByOwnerIDAndSubject(info *models.DocumentsInfo, page Page) (models.DocumentsInfo, *errs.Error)

	UpdateDocumentTitleByID(update *models.DocumentTitleUpdate) *errs.Error
	UpdateDocumentSubjectID(update *models.DocumentSubjectUpdate) *errs.Error

	DeleteDocumentByID(documentID types.ID) *errs.Error
}

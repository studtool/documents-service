package repositories

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
)

type Page struct {
	Index int32
	Size  int32
}

type DocumentsInfoRepository interface {
	AddDocumentInfo(info *models.DocumentInfoFull) *errs.Error

	GetDocumentInfo(documentID string) (*models.DocumentInfo, *errs.Error)
	GetDocumentInfoFull(documentID string) (*models.DocumentInfoFull, *errs.Error)

	GetDocumentsInfo(userID string, ownerID string, subject *string, page Page) (models.DocumentsInfo, *errs.Error)
	DeleteDocumentsInfo(ownerID string, subject *string) *errs.Error

	AddDocumentMember(documentID string, member *models.Permission) *errs.Error
	UpdateDocumentMemberPrivilege(documentID string, member *models.Permission) *errs.Error
	DeleteDocumentMember(documentID string, memberId string) *errs.Error

	AddDocumentUpdateToHistory(documentID string, info *models.UpdateInfo) *errs.Error
}

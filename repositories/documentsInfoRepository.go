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

	GetDocumentInfo(documentId string) (*models.DocumentInfo, *errs.Error)
	GetDocumentInfoFull(documentId string) (*models.DocumentInfoFull, *errs.Error)

	GetDocumentsInfo(ownerId string, subject *string, page Page) ([]models.DocumentInfo, *errs.Error)
	DeleteDocumentsInfo(ownerId string, subject *string) *errs.Error

	UpdateDocumentMeta(documentId string, meta *models.DocumentMeta) *errs.Error

	AddDocumentMember(documentId string, member *models.MemberInfo) *errs.Error
	UpdateDocumentMemberPrivilege(documentId string, member *models.MemberInfo) *errs.Error
	DeleteDocumentMember(documentId string, memberId string) *errs.Error

	AddDocumentUpdateToHistory(documentId string, info *models.UpdateInfo) *errs.Error
}

package postgres

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
)

type DocumentsInfoRepository struct {
	conn        *Connection
	notFoundErr *errs.Error
}

func NewDocumentsInfoRepository(conn *Connection) *DocumentsInfoRepository {
	return &DocumentsInfoRepository{
		conn:        conn,
		notFoundErr: errs.NewNotFoundError("profile not found"),
	}
}

func (r *DocumentsInfoRepository) AddDocumentInfo(info *models.DocumentInfoFull) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) GetDocumentInfo(documentId string) (*models.DocumentInfo, *errs.Error) {
	panic("implement me")
}

func (r *DocumentsInfoRepository) GetDocumentInfoFull(documentId string) (*models.DocumentInfoFull, *errs.Error) {
	panic("implement me")
}

func (r *DocumentsInfoRepository) GetDocumentsInfo(ownerId string,
	subject *string, page repositories.Page) ([]models.DocumentInfo, *errs.Error) {

	panic("implement me")
}

func (r *DocumentsInfoRepository) DeleteDocumentsInfo(ownerId string, subject *string) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) UpdateDocumentMeta(documentId string, meta *models.DocumentMeta) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) AddDocumentMember(documentId string, member *models.MemberInfo) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) UpdateDocumentMemberPrivilege(documentId string,
	member *models.MemberInfo) *errs.Error {

	panic("implement me")
}

func (r *DocumentsInfoRepository) DeleteDocumentMember(documentId string, memberId string) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) AddDocumentUpdateToHistory(documentId string, info *models.UpdateInfo) *errs.Error {
	panic("implement me")
}

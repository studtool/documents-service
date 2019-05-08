package mongo

import (
	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/config"
)

const (
	SelectTimeout = time.Second
	UpdateTimeout = 2 * time.Second
)

const (
	idField = "_id"
)

type DocumentsInfoRepository struct {
	connection  *Connection
	collection  *mongo.Collection
	notFoundErr *errs.Error
}

func NewDocumentsInfoRepository(conn *Connection) *DocumentsInfoRepository {
	db := conn.client.Database(config.StorageDB.Value())
	return &DocumentsInfoRepository{
		connection:  conn,
		collection:  db.Collection("documents"),
		notFoundErr: errs.NewNotFoundError("document not found"),
	}
}

func (r *DocumentsInfoRepository) SaveDocumentInfo(info *models.DocumentInfoFull) *errs.Error {
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

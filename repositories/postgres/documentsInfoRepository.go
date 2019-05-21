package postgres

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/beans"
	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
)

type DocumentsInfoRepository struct {
	conn *Connection

	documentNotFound       *errs.Error
	documentTitleDuplicate *errs.Error
}

func NewDocumentsInfoRepository(conn *Connection) *DocumentsInfoRepository {
	return &DocumentsInfoRepository{
		conn: conn,

		documentNotFound:       errs.NewNotFoundError("document not found"),
		documentTitleDuplicate: errs.NewConflictError("document title duplicate"),
	}
}

func (r *DocumentsInfoRepository) AddDocumentInfo(info *models.DocumentInfoFull) *errs.Error {
	const query = `
		INSERT INTO document(id,title,owner_id,subject) VALUES($1,$2,$3,$4);
	`
	var err error

	id, err := uuid.NewRandom()
	if err != nil {
		return errs.New(err)
	}
	info.ID = types.ID(id.String())

	_, err = r.conn.db.Exec(query,
		&info.ID, &info.Title, &info.OwnerId, &info.Subject)
	if err != nil {
		if strings.Contains(err.Error(), "documents_title_owner_id_subject_unique") {
			return r.documentTitleDuplicate
		} else {
			return errs.New(err)
		}
	}

	return nil
}

func (r *DocumentsInfoRepository) GetDocumentInfo(documentId string) (*models.DocumentInfo, *errs.Error) {
	const query = `
		SELECT d.title, d.owner_id, d.subject
		FROM document d WHERE d.id = $1;
	`

	rows, err := r.conn.db.Query(query, &documentId)
	if err != nil {
		return nil, errs.New(err)
	}
	defer r.closeRows(rows)

	if !rows.Next() {
		return nil, r.documentNotFound
	}

	info := &models.DocumentInfo{
		ID: types.ID(documentId),
	}
	if err := rows.Scan(&info.Title, &info.OwnerId, &info.Subject); err != nil {
		return nil, errs.New(err)
	}

	return info, nil
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

func (r *DocumentsInfoRepository) closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		beans.Logger().Error(err)
	}
}

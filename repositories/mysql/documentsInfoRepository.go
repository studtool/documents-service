package mysql

import (
	"database/sql"
	"github.com/google/uuid"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/beans"
	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
)

type DocumentsInfoRepository struct {
	conn *Connection

	documentNotFound  *errs.Error
	documentsNotFound *errs.Error
}

func NewDocumentsInfoRepository(conn *Connection) *DocumentsInfoRepository {
	return &DocumentsInfoRepository{
		conn: conn,

		documentNotFound:  errs.NewNotFoundError("document not found"),
		documentsNotFound: errs.NewNotFoundError("documents not found"),
	}
}

func (r *DocumentsInfoRepository) AddDocumentInfo(info *models.DocumentInfoFull) *errs.Error {
	const query = `
		INSERT INTO document(id, title, owner_id, subject) VALUES(?,?,?,?);
	`
	var err error

	id, err := uuid.NewRandom()
	if err != nil {
		return errs.New(err)
	}
	info.ID = types.ID(id.String())

	_, err = r.conn.db.Exec(query,
		&info.ID, &info.Title, &info.OwnerID, &info.Subject)
	if err != nil {
		return errs.New(err)
	}

	return nil
}

func (r *DocumentsInfoRepository) GetDocumentInfo(documentID string) (*models.DocumentInfo, *errs.Error) {
	const query = `
		SELECT d.title, d.owner_id, d.subject
		FROM document d WHERE d.id = ?;
	`

	rows, err := r.conn.db.Query(query, &documentID)
	if err != nil {
		return nil, errs.New(err)
	}
	defer r.closeRows(rows)

	if !rows.Next() {
		return nil, r.documentNotFound
	}

	info := &models.DocumentInfo{
		ID: types.ID(documentID),
	}
	if err := rows.Scan(&info.Title, &info.OwnerID, &info.Subject); err != nil {
		return nil, errs.New(err)
	}

	return info, nil
}

func (r *DocumentsInfoRepository) GetDocumentInfoFull(
	documentID string,
) (*models.DocumentInfoFull, *errs.Error) {
	panic("implement me")
}

func (r *DocumentsInfoRepository) GetDocumentsInfo(
	ownerID string,
	subject *string,
	page repositories.Page,
) (models.DocumentsInfo, *errs.Error) {
	const query = `
		SELECT
			d.id,
			d.title,
			d.owner_id,
			d.subject
		FROM document d
		WHERE
			d.owner_id = ? AND
			d.subject = ?
		LIMIT ? OFFSET ?;
	`

	page.Index *= page.Size

	rows, err := r.conn.db.Query(query,
		&ownerID, subject, &page.Size, &page.Index,
	)
	if err != nil {
		return nil, errs.New(err)
	}
	defer r.closeRows(rows)

	documents := make([]models.DocumentInfo, 0)
	for rows.Next() {
		document := models.DocumentInfo{}
		if err := rows.Scan(
			&document.ID, &document.Title,
			&document.Subject, &document.OwnerID,
		); err != nil {
			return nil, errs.New(err)
		}
		documents = append(documents, document)
	}
	if len(documents) == 0 {
		return nil, r.documentsNotFound
	}

	return models.DocumentsInfo(documents), nil
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

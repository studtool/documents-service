package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type DocumentsInfoRepository struct {
	conn *Connection

	structLogger  logs.Logger
	reflectLogger logs.Logger

	docNotFoundErr  *errs.Error
	docsNotFoundErr *errs.Error
}

func NewDocumentsInfoRepository(conn *Connection) *DocumentsInfoRepository {
	r := &DocumentsInfoRepository{
		conn: conn,

		docNotFoundErr:  errs.NewNotFoundError("document not found"),
		docsNotFoundErr: errs.NewNotFoundError("documents not found"),
	}

	r.structLogger = srvutils.MakeStructLogger(r)
	r.reflectLogger = srvutils.MakeReflectLogger(r)

	r.structLogger.Info("initialized")

	return r
}

func (r *DocumentsInfoRepository) AddDocumentInfo(info *models.DocumentInfo) *errs.Error {
	const query = `
		INSERT INTO document(id,title,owner_id,subject) VALUES(?,?,?,?);
	`

	_, err := r.db().ExecContext(r.iCtx(), query,
		info.DocumentID, info.Title, info.OwnerID, info.Subject,
	)
	if err != nil {
		return r.wrapErr(err)
	}

	return nil
}

func (r *DocumentsInfoRepository) GetDocumentInfoByID(info *models.DocumentInfo) *errs.Error {
	const query = `
		SELECT
			d.title,
			d.owner_id,
			d.subject
		FROM document d
		WHERE
			d.id = ?;
	`

	row := r.db().QueryRowContext(r.sCtx(), query, info.DocumentID)
	err := row.Scan(
		&info.Title, &info.OwnerID, &info.Subject,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return r.docNotFoundErr
		} else {
			return r.wrapErr(err)
		}
	}

	return nil
}

const (
	selectDocumentFields = "d.id,d.owner_id,d.title,d.subject"
)

func (r *DocumentsInfoRepository) GetDocumentsInfoByOwnerID(
	ownerID types.ID, page repositories.Page) (models.DocumentsInfo, *errs.Error) {

	const query = `
		SELECT ` + selectDocumentFields + `
		FROM document d
		WHERE
			d.owner_id = ?
		LIMIT ? OFFSET ?;
	`

	page.Index *= page.Size
	return r.getDocumentsInfo(query, ownerID, page.Size, page.Index)
}

func (r *DocumentsInfoRepository) GetDocumentsInfoByOwnerIDAndSubject(
	ownerID types.ID, subject string, page repositories.Page) (models.DocumentsInfo, *errs.Error) {

	const query = `
		SELECT ` + selectDocumentFields + `
		FROM document d
		WHERE
			d.owner_id = ? AND
			d.subject = ?
		LIMIT ? OFFSET ?;
	`

	page.Index *= page.Size
	return r.getDocumentsInfo(query, ownerID, subject, page.Size, page.Index)
}

func (r *DocumentsInfoRepository) UpdateDocumentTitleByID(update *models.DocumentTitleUpdate) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) UpdateDocumentSubjectByID(update *models.DocumentSubjectUpdate) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) CheckDocumentExistsByIDAndOwnerID(documentID types.ID, ownerID types.ID) *errs.Error {
	const query = `
		SELECT EXISTS(
			SELECT * FROM document d
			WHERE d.id = ? AND d.owner_id = ?
		);
	`

	row := r.db().QueryRowContext(r.sCtx(), query, documentID, ownerID)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return r.docNotFoundErr
		} else {
			return r.wrapErr(err)
		}
	}

	return nil
}

func (r *DocumentsInfoRepository) DeleteDocumentByID(documentID types.ID) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) db() *sql.DB {
	return r.conn.db
}

func (r *DocumentsInfoRepository) getDocumentsInfo(
	query string, args ...interface{}) (models.DocumentsInfo, *errs.Error) {

	rows, err := r.db().Query(query, args...)
	if err != nil {
		return nil, r.wrapErr(err)
	}
	defer r.closeRows(rows)

	documents := make(models.DocumentsInfo, 0)
	for rows.Next() {
		document := models.DocumentInfo{}
		err := rows.Scan(
			&document.DocumentID, &document.OwnerID,
			&document.Title, &document.Subject,
		)
		if err != nil {
			return nil, r.wrapErr(err)
		}
		documents = append(documents, document)
	}
	if len(documents) == 0 {
		return nil, r.docsNotFoundErr
	}

	return documents, nil
}

func (r *DocumentsInfoRepository) closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		r.reflectLogger.Error(err)
	}
}

func (r *DocumentsInfoRepository) iCtx() context.Context {
	return r.ctx(500 * time.Millisecond)
}

func (r *DocumentsInfoRepository) sCtx() context.Context {
	return r.ctx(200 * time.Millisecond)
}

func (r *DocumentsInfoRepository) msCtx() context.Context {
	return r.ctx(time.Second)
}

func (r *DocumentsInfoRepository) ctx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	return ctx
}

func (r *DocumentsInfoRepository) wrapErr(err error) *errs.Error {
	r.reflectLogger.Error(err)
	return errs.New(err)
}

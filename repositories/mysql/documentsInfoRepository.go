package mysql

import (
	"context"
	"database/sql"
	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"
	"github.com/studtool/common/types"
	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
	"time"
)

type DocumentsInfoRepository struct {
	conn *Connection

	structLogger  logs.Logger
	reflectLogger logs.Logger

	docNotFoundErr  *errs.Error
	docsNotFoundErr *errs.Error
}

func NewDocumentsInfoRepository(conn *Connection) *DocumentsInfoRepository {
	structLogger := logs.NewStructLogger(
		logs.StructLoggerParams{
			Component: config.Component,
			Structure: "mysql.DocumentsRepository",
		},
	)

	structLogger.Info("initialization")

	return &DocumentsInfoRepository{
		conn: conn,

		docNotFoundErr:  errs.NewNotFoundError("document not found"),
		docsNotFoundErr: errs.NewNotFoundError("documents not found"),
	}
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
			r.logDocNotFound(info.DocumentID)
			return r.docNotFoundErr
		} else {
			return r.wrapErr(err)
		}
	}

	return nil
}

func (r *DocumentsInfoRepository) GetDocumentsInfoByOwnerID(ownerID types.ID, page repositories.Page) (models.DocumentsInfo, *errs.Error) {
	const query = `
		SELECT
			d.id,
			d.title,
			d.subject
		FROM document d
		WHERE
			d.owner_id = ?
		LIMIT ? OFFSET ?;
	`

	page.Index *= page.Size

	rows, err := r.db().QueryContext(r.msCtx(), query,
		ownerID, page.Size, page.Index,
	)
	if err != nil {
		return nil, r.wrapErr(err)
	}
	defer r.closeRows(rows)

	documents := make(models.DocumentsInfo, 0)
	for rows.Next() {
		document := models.DocumentInfo{
			OwnerID: ownerID,
		}
		err := rows.Scan(
			&document.DocumentID, &document.Title, &document.Subject,
		)
		if err != nil {
			return nil, r.wrapErr(err)
		}
	}
	if len(documents) == 0 {
		r.logDocsNotFound(ownerID)
		return nil, r.docsNotFoundErr
	}

	return documents, nil
}

func (r *DocumentsInfoRepository) GetDocumentsInfoByOwnerIDAndSubject(info *models.DocumentsInfo, page repositories.Page) (models.DocumentsInfo, *errs.Error) {
	panic("implement me")
}

func (r *DocumentsInfoRepository) UpdateDocumentTitleByID(update *models.DocumentTitleUpdate) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) UpdateDocumentSubjectID(update *models.DocumentSubjectUpdate) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) DeleteDocumentByID(documentID types.ID) *errs.Error {
	panic("implement me")
}

func (r *DocumentsInfoRepository) db() *sql.DB {
	return r.conn.db
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
	return r.ctx(100 * time.Microsecond)
}

func (r *DocumentsInfoRepository) msCtx() context.Context {
	return r.ctx(time.Second)
}

func (r *DocumentsInfoRepository) ctx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	return ctx
}

func (r *DocumentsInfoRepository) logDocsNotFound(ownerID types.ID) {
	r.structLogger.Warningf("documents [owner_id = %s] not found", ownerID)
}

func (r *DocumentsInfoRepository) logDocNotFound(documentID types.ID) {
	r.structLogger.Warningf("document [id = %s] not found", documentID)
}

func (r *DocumentsInfoRepository) wrapErr(err error) *errs.Error {
	r.reflectLogger.Error(err)
	return errs.New(err)
}

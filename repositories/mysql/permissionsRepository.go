package mysql

import (
	"errors"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type PermissionsRepository struct {
	conn *Connection

	noPermissionErr *errs.Error
}

const (
	scopeReadFlag  = 1 << 0
	scopeWriteFlag = 1 << 1
)

func NewPermissionsRepository(conn *Connection) *PermissionsRepository {
	return &PermissionsRepository{
		conn: conn,

		noPermissionErr: errs.NewNotFoundError("no permission"),
	}
}

func (r *PermissionsRepository) AddPermission(documentID types.ID, p *models.Permission) *errs.Error {
	const query = `
		INSERT INTO permission(user_id, document_id, scope) VALUES(?,?,?);
	`

	scope := getScopeFlag(p.Scope)

	_, err := r.conn.db.Exec(query,
		&p.UserID, &documentID, &scope,
	)
	if err != nil {
		return errs.New(err)
	}

	return nil
}

func (r *PermissionsRepository) CheckPermission(documentID types.ID, p *models.Permission) *errs.Error {
	const query = `
		SELECT EXISTS(
			SELECT * FROM permission
			WHERE
				user_id = ? AND
				document_id = ? AND
				scope = ?
		);
	`

	scope := getScopeFlag(p.Scope)

	rows, err := r.conn.db.Query(query,
		&p.UserID, &documentID, &scope,
	)
	if err != nil {
		return errs.New(err)
	}

	if !rows.Next() {
		return errs.New(errors.New("exist query returned 0 rows"))
	}

	var ok bool
	if err := rows.Scan(&ok); err != nil {
		return errs.New(err)
	}

	if !ok {
		return r.noPermissionErr
	}
	return nil
}

func (r *PermissionsRepository) UpdatePermission(documentID types.ID, p *models.Permission) *errs.Error {
	panic("implement me")
}

func (r *PermissionsRepository) DeletePermission(documentID types.ID, p *models.Permission) *errs.Error {
	panic("implement me")
}

func getScopeFlag(scope models.Scope) int {
	switch scope {
	case models.ScopeRead:
		return scopeReadFlag
	case models.ScopeWrite:
		return scopeReadFlag | scopeWriteFlag
	default:
		return 0
	}
}

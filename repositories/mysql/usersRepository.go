package mysql

import (
	"database/sql"

	"go.uber.org/dig"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/utils"
)

type UsersRepository struct {
	conn *Connection

	structLogger  logs.Logger
	reflectLogger logs.Logger

	userNotFoundErr *errs.Error
}

type UsersRepositoryParams struct {
	dig.In

	Connection *Connection

	LogsExporter *logs.Exporter
}

func NewUsersRepository(params UsersRepositoryParams) *UsersRepository {
	r := &UsersRepository{
		conn: params.Connection,

		userNotFoundErr: errs.NewNotFoundError("user not found"),
	}

	p := srvutils.LoggerParams{
		Value:    r,
		Exporter: params.LogsExporter,
	}

	r.structLogger = srvutils.MakeStructLogger(p)
	r.reflectLogger = srvutils.MakeReflectLogger(p)

	r.structLogger.Info("initialized")

	return r
}

func (r *UsersRepository) AddUser(u *models.User) *errs.Error {
	const query = `
		INSERT INTO user(id) VALUES(?);
	`

	_, err := r.db().Exec(query, u.ID)
	if err != nil {
		return r.wrapErr(err)
	}

	return nil
}

func (r *UsersRepository) CheckUserExistsByID(userID types.ID) *errs.Error {
	const query = `
		SELECT exists(
		    SELECT * FROM user u
		    WHERE u.id = ?
		);
	`

	row := r.db().QueryRow(query, userID)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return r.wrapErr(err)
	}
	if !exists {
		return r.userNotFoundErr
	}

	return nil
}

func (r *UsersRepository) DeleteUserByID(userID types.ID) *errs.Error {
	const query = `
		DELETE FROM user u
		WHERE u.id = ?;
	`

	res, err := r.db().Exec(query, userID)
	if err != nil {
		return r.wrapErr(err)
	}
	if n, _ := res.RowsAffected(); n != 1 {
		return r.userNotFoundErr
	}

	return nil
}

func (r *UsersRepository) db() *sql.DB {
	return r.conn.db
}

func (r *UsersRepository) wrapErr(err error) *errs.Error {
	r.reflectLogger.Error(err)
	return errs.New(err)
}

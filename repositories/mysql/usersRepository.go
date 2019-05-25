package mysql

import (
	"database/sql"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/models"
)

type UsersRepository struct {
	conn *Connection

	structLogger  logs.Logger
	reflectLogger logs.Logger

	userNotFoundErr *errs.Error
}

func NewUsersRepository(conn *Connection) *UsersRepository {
	structLogger := logs.NewStructLogger(
		logs.StructLoggerParams{
			Component: config.Component,
			Structure: "mysql.UsersRepository",
		},
	)

	structLogger.Info("initialization")

	return &UsersRepository{
		conn: conn,

		structLogger:  structLogger,
		reflectLogger: logs.NewReflectLogger(),

		userNotFoundErr: errs.NewNotFoundError("user not found"),
	}
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

func (r *UsersRepository) CheckExistsUserByID(userID string) *errs.Error {
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

func (r *UsersRepository) DeleteUserByID(userID string) *errs.Error {
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

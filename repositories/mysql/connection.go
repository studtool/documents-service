package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/utils"
)

type Connection struct {
	db *sql.DB

	structLogger  logs.Logger
	reflectLogger logs.Logger

	connStr string
}

func NewConnection() *Connection {
	conn := &Connection{
		connStr: fmt.Sprintf(
			"%s:%s@(%s:%s)/%s",
			config.StorageUser.Value(), config.StoragePassword.Value(),
			config.StorageHost.Value(), config.StoragePort.Value(),
			config.StorageDB.Value(),
		),
	}

	conn.structLogger = utils.MakeStructLogger(conn)
	conn.reflectLogger = utils.MakeReflectLogger(conn)

	conn.structLogger.Info("initialization")

	return conn
}

func (conn *Connection) Open() (err error) {
	conn.db, err = sql.Open("mysql", conn.connStr)
	return err
}

func (conn *Connection) Close() (err error) {
	return conn.db.Close()
}

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

	conn.structLogger = srvutils.MakeStructLogger(conn)
	conn.reflectLogger = srvutils.MakeReflectLogger(conn)

	conn.structLogger.Info("initialized")

	return conn
}

func (conn *Connection) Open() (err error) {
	conn.db, err = sql.Open("mysql", conn.connStr)
	if err == nil {
		conn.structLogger.Info("opened")
	}
	return
}

func (conn *Connection) Close() (err error) {
	err = conn.db.Close()
	if err == nil {
		conn.structLogger.Info("closed")
	}
	return
}

package mysql

import (
	"database/sql"
	"fmt"

	"go.uber.org/dig"

	// nolint:golint
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

type ConnectionParams struct {
	dig.In

	LogsExporter *logs.Exporter
}

func NewConnection(params ConnectionParams) *Connection {
	conn := &Connection{
		connStr: fmt.Sprintf(
			"%s:%s@(%s:%s)/%s",
			config.StorageUser.Value(), config.StoragePassword.Value(),
			config.StorageHost.Value(), config.StoragePort.Value(),
			config.StorageDB.Value(),
		),
	}

	p := srvutils.LoggerParams{
		Value:    conn,
		Exporter: params.LogsExporter,
	}

	conn.structLogger = srvutils.MakeStructLogger(p)
	conn.reflectLogger = srvutils.MakeReflectLogger(p)

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

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/studtool/documents-service/config"
)

type Connection struct {
	connStr string
	db      *sql.DB
}

func NewConnection() *Connection {
	return &Connection{
		connStr: fmt.Sprintf(
			"%s:%s@(%s:%s)/%s",
			config.StorageUser.Value(), config.StoragePassword.Value(),
			config.StorageHost.Value(), config.StoragePort.Value(),
			config.StorageDB.Value(),
		),
	}
}

func (conn *Connection) Open() (err error) {
	conn.db, err = sql.Open("mysql", conn.connStr)
	return err
}

func (conn *Connection) Close() (err error) {
	return conn.db.Close()
}

package logs

import (
	"fmt"
	"net"

	"github.com/firstrow/goautosocket"

	"github.com/studtool/logrustash"
)

type Exporter struct {
	params ExporterParams

	conn       net.Conn
	logrusHook *logrustash.Hook
}

type ExporterParams struct {
	StorageAddress   string
	ComponentName    string
	ComponentVersion string
}

func NewLogsExporter(params ExporterParams) *Exporter {
	return &Exporter{
		params: params,
	}
}

func (e *Exporter) OpenConnection() error {
	conn, err := gas.Dial("tcp", e.params.StorageAddress)
	if err != nil {
		return err
	}

	appName := fmt.Sprintf("%s:%s",
		e.params.ComponentName, e.params.ComponentVersion,
	)

	hook, err := logrustash.NewHookWithConn(conn, appName)
	if err != nil {
		return err
	}

	e.conn = conn
	e.logrusHook = hook

	return nil
}

func (e *Exporter) CloseConnection() error {
	return e.conn.Close()
}

func (e *Exporter) getHook() *logrustash.Hook {
	return e.logrusHook
}

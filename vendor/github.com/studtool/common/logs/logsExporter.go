package logs

import (
	"github.com/studtool/logrustash"
)

type Exporter struct {
	params     ExporterParams
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
	hook, err := logrustash.NewHook(
		"tcp", e.params.StorageAddress,
		e.params.ComponentName+":"+e.params.ComponentVersion,
	)
	if err != nil {
		return err
	}
	e.logrusHook = hook
	return nil
}

func (e *Exporter) CloseConnection() error {
	return nil //TODO
}

func (e *Exporter) getHook() *logrustash.Hook {
	return e.logrusHook
}

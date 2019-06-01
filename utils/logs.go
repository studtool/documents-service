package srvutils

import (
	"github.com/studtool/common/logs"
	"github.com/studtool/common/rft"

	"github.com/studtool/documents-service/config"
)

func MakeRawLogger(_ interface{}) logs.Logger {
	return logs.NewRawLogger()
}

type LoggerParams struct {
	Value    interface{}
	Exporter *logs.Exporter
}

func MakeStructLogger(params LoggerParams) logs.Logger {
	return logs.NewStructLogger(
		logs.StructLoggerParams{
			Exporter:          params.Exporter,
			ComponentName:     config.ComponentName,
			ComponentVersion:  config.ComponentVersion,
			StructWithPkgName: rft.StructName(params.Value),
		},
	)
}

func MakeReflectLogger(params LoggerParams) logs.Logger {
	return logs.NewReflectLogger(
		logs.ReflectLoggerParams{
			Exporter: params.Exporter,
		},
	)
}

func MakeRequestLogger(params LoggerParams) logs.Logger {
	return logs.NewRequestLogger(
		logs.RequestLoggerParams{
			Exporter:         params.Exporter,
			ComponentName:    config.ComponentName,
			ComponentVersion: config.ComponentVersion,
		},
	)
}

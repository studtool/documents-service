package utils

import (
	"github.com/studtool/common/logs"
	"github.com/studtool/common/utils"

	"github.com/studtool/documents-service/config"
)

func MakeRawLogger(_ interface{}) logs.Logger {
	return logs.NewRawLogger()
}

func MakeStructLogger(v interface{}) logs.Logger {
	return logs.NewStructLogger(
		logs.StructLoggerParams{
			Component: config.Component,
			Structure: utils.StructName(v),
		},
	)
}

func MakeReflectLogger(_ interface{}) logs.Logger {
	return logs.NewReflectLogger()
}

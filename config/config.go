package config

import (
	"github.com/studtool/common/config"

	"github.com/studtool/documents-service/beans"
)

var (
	_ = func() *config.FlagVar {
		f := config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			config.SetLogger(beans.Logger())
		}
		return f
	}()

	ServerPort = config.NewStringDefault("STUDTOOL_DOCUMENTS_SERVICE_PORT", "80")

	CorsAllowed         = config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_ALLOW_CORS", false)
	RequestsLogsEnabled = config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_REQUESTS", true)
)

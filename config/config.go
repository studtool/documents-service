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

	RepositoriesEnabled = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_REPOSITORIES_ENABLED", false)

	StorageHost     = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_HOST", "127.0.0.1")
	StoragePort     = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PORT", "5432")
	StorageDB       = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_NAME", "auth")
	StorageUser     = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_USER", "user")
	StoragePassword = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PASSWORD", "password")
	StorageSSL      = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_SSL_MODE", "disable")
)

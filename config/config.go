package config

import (
	"github.com/studtool/common/config"

	"github.com/studtool/documents-service/beans"
)

var (
	_ = func() *cconfig.FlagVar {
		f := cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			cconfig.SetLogger(beans.Logger())
		}
		return f
	}()

	ServerPort = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_SERVICE_PORT", "80")

	CorsAllowed         = cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_ALLOW_CORS", false)
	RequestsLogsEnabled = cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_REQUESTS", true)

	RepositoriesEnabled = cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_REPOSITORIES_ENABLED", false)

	DocsDir = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_DIRECTORY_PATH", "./documents")

	StorageHost     = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_HOST", "127.0.0.1")
	StoragePort     = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PORT", "5432")
	StorageDB       = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_NAME", "documents")
	StorageUser     = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_USER", "user")
	StoragePassword = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PASSWORD", "password")
	StorageSSL      = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_SSL_MODE", "disable")
)

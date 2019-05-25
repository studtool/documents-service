package config

import (
	"github.com/studtool/common/config"
	"github.com/studtool/common/logs"
	"strconv"
)

var (
	// Modified by the compiler
	Component = ""

	// Modified by the compiler
	repositoriesEnabled = "true"

	//Modified by the compiler
	servicesEnabled = "true"

	// Modified by the compiler
	queuesEnabled = "true"

	_ = func() *cconfig.FlagVar {
		f := cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			cconfig.SetLogger(logs.NewRawLogger())
		}
		return f
	}()

	ServerPort = cconfig.NewIntDefault("STUDTOOL_DOCUMENTS_SERVICE_PORT", 80)

	CorsAllowed         = cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_ALLOW_CORS", false)
	RequestsLogsEnabled = cconfig.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_REQUESTS", true)

	RepositoriesEnabled = parseBool(repositoriesEnabled)
	ServicesEnabled     = parseBool(servicesEnabled)
	QueuesEnabled       = parseBool(queuesEnabled)

	StorageHost     = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_HOST", "127.0.0.1")
	StoragePort     = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PORT", "3306")
	StorageDB       = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_NAME", "documents")
	StorageUser     = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_USER", "user")
	StoragePassword = cconfig.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PASSWORD", "password")
)

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}

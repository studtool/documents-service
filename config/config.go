package config

import (
	"strconv"
	"time"

	"github.com/studtool/common/config"
	"github.com/studtool/common/logs"
)

var (
	// TODO compile-time injection
	// nolint:golint,gochecknoglobals
	ComponentName = "documents-service"

	// TODO compile-time injection
	// nolint:golint,gochecknoglobals
	ComponentVersion = "v0.0.1"

	// Modified by the compiler
	// nolint:golint,gochecknoglobals
	repositoriesEnabled = "true"

	// Modified by the compiler
	// nolint:golint,gochecknoglobals
	servicesEnabled = "true"

	// Modified by the compiler
	// nolint:golint,gochecknoglobals
	queuesEnabled = "true"

	_ = func() *config.FlagVar {
		f := config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			config.SetLogger(logs.NewRawLogger())
		}
		return f
	}()

	// nolint:golint,gochecknoglobals
	ServerPort = config.NewIntDefault("STUDTOOL_DOCUMENTS_SERVICE_PORT", 80)

	// nolint:golint,gochecknoglobals
	CorsAllowed = config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_ALLOW_CORS", false)

	// nolint:golint,gochecknoglobals
	RequestsLogsEnabled = config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_REQUESTS", true)

	// nolint:golint,gochecknoglobals
	RepositoriesEnabled = parseBool(repositoriesEnabled)

	// nolint:golint,gochecknoglobals
	ServicesEnabled = parseBool(servicesEnabled)

	// nolint:golint,gochecknoglobals
	QueuesEnabled = parseBool(queuesEnabled)

	// nolint:golint,gochecknoglobals
	StorageHost = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_HOST", "127.0.0.1")
	// nolint:golint,gochecknoglobals
	StoragePort = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PORT", "3306")
	// nolint:golint,gochecknoglobals
	StorageDB = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_NAME", "documents")
	// nolint:golint,gochecknoglobals
	StorageUser = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_USER", "user")
	// nolint:golint,gochecknoglobals
	StoragePassword = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PASSWORD", "password")

	// nolint:golint,gochecknoglobals
	MqHost = config.NewStringDefault("STUDTOOL_MQ_HOST", "127.0.0.1")
	// nolint:golint,gochecknoglobals
	MqPort = config.NewIntDefault("STUDTOOL_MQ_PORT", 5672)
	// nolint:golint,gochecknoglobals
	MqUser = config.NewStringDefault("STUDTOOL_MQ_USER", "user")
	// nolint:golint,gochecknoglobals
	MqPassword = config.NewStringDefault("STUDTOOL_MQ_PASSWORD", "password")

	// nolint:golint,gochecknoglobals
	MqConnNumRet = config.NewIntDefault("STUDTOOL_MQ_CONNECTION_NUM_RETRIES", 10)
	// nolint:golint,gochecknoglobals
	MqConnRetItv = config.NewTimeDefault("STUDTOOL_MQ_CONNECTION_RETRY_INTERVAL", 2*time.Second)
)

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}

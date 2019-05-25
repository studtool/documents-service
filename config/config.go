package config

import (
	"strconv"
	"time"

	"github.com/studtool/common/config"
	"github.com/studtool/common/logs"
)

var (
	// TODO compile-time injection
	Component = "documents-service"

	// Modified by the compiler
	repositoriesEnabled = "true"

	//Modified by the compiler
	servicesEnabled = "true"

	// Modified by the compiler
	queuesEnabled = "true"

	_ = func() *config.FlagVar {
		f := config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			config.SetLogger(logs.NewRawLogger())
		}
		return f
	}()

	ServerPort = config.NewIntDefault("STUDTOOL_DOCUMENTS_SERVICE_PORT", 80)

	CorsAllowed         = config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_ALLOW_CORS", false)
	RequestsLogsEnabled = config.NewFlagDefault("STUDTOOL_DOCUMENTS_SERVICE_SHOULD_LOG_REQUESTS", true)

	RepositoriesEnabled = parseBool(repositoriesEnabled)
	ServicesEnabled     = parseBool(servicesEnabled)
	QueuesEnabled       = parseBool(queuesEnabled)

	StorageHost     = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_HOST", "127.0.0.1")
	StoragePort     = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PORT", "3306")
	StorageDB       = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_NAME", "documents")
	StorageUser     = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_USER", "user")
	StoragePassword = config.NewStringDefault("STUDTOOL_DOCUMENTS_INFO_STORAGE_PASSWORD", "password")

	MqHost     = config.NewStringDefault("STUDTOOL_MQ_HOST", "127.0.0.1")
	MqPort     = config.NewIntDefault("STUDTOOL_MQ_PORT", 5672)
	MqUser     = config.NewStringDefault("STUDTOOL_MQ_USER", "user")
	MqPassword = config.NewStringDefault("STUDTOOL_MQ_PASSWORD", "password")

	MqConnNumRet = config.NewIntDefault("STUDTOOL_MQ_CONNECTION_NUM_RETRIES", 10)
	MqConnRetItv = config.NewTimeDefault("STUDTOOL_MQ_CONNECTION_RETRY_INTERVAL", 2*time.Second)
)

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}

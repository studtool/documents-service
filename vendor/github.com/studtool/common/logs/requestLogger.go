package logs

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/studtool/common/utils"
)

type RequestLogger struct {
	host      string
	pid       int64
	component string

	logger *logrus.Logger
}

type RequestLoggerParams struct {
	Component string
}

func NewRequestLogger(params RequestLoggerParams) Logger {
	return &RequestLogger{
		host:      utils.GetHost(),
		pid:       utils.GetPid(),
		component: params.Component,

		logger: func() *logrus.Logger {
			log := logrus.StandardLogger()
			log.SetFormatter(&logrus.JSONFormatter{})
			return log
		}(),
	}
}

type RequestParams struct {
	Method string
	Path   string
	Status int
	Type   string
	UserID string
	Time   time.Duration
}

const (
	WrongMethodMessage = "method with format should not be called"
)

func (log *RequestLogger) Debug(args ...interface{}) {
	log.logger.WithFields(log.makeLogFields(args...)).Debug()
}

func (log *RequestLogger) Debugf(format string, args ...interface{}) {
	panic(WrongMethodMessage)
}

func (log *RequestLogger) Info(args ...interface{}) {
	log.logger.WithFields(log.makeLogFields(args...)).Info(args...)
}

func (log *RequestLogger) Infof(format string, args ...interface{}) {
	panic(WrongMethodMessage)
}

func (log *RequestLogger) Warning(args ...interface{}) {
	log.logger.WithFields(log.makeLogFields(args...)).Warn(args...)
}

func (log *RequestLogger) Warningf(format string, args ...interface{}) {
	panic(WrongMethodMessage)
}

func (log *RequestLogger) Error(args ...interface{}) {
	log.logger.WithFields(log.makeLogFields(args...)).Error(args...)
}

func (log *RequestLogger) Errorf(format string, args ...interface{}) {
	panic(WrongMethodMessage)
}

func (log *RequestLogger) Fatal(args ...interface{}) {
	log.logger.WithFields(log.makeLogFields(args...)).Fatal(args...)
}

func (log *RequestLogger) Fatalf(format string, args ...interface{}) {
	panic(WrongMethodMessage)
}

func (log *RequestLogger) makeLogFields(args ...interface{}) logrus.Fields {
	if len(args) != 1 {
		panic(args)
	}

	p, ok := args[0].(*RequestParams)
	if !ok {
		panic(args)
	}

	return logrus.Fields{
		"host":      log.host,
		"pid":       log.pid,
		"component": log.component,
		"method":    p.Method,
		"path":      p.Path,
		"status":    p.Status,
		"type":      p.Type,
		"userId":    p.UserID,
		"time":      p.Time,
	}
}

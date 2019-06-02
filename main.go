package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/utils/assertions"

	"github.com/studtool/documents-service/api"
	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/logic/fake"
	"github.com/studtool/documents-service/logic/impl"
	"github.com/studtool/documents-service/messages"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/repositories/fake"
	"github.com/studtool/documents-service/repositories/memory"
	"github.com/studtool/documents-service/repositories/mysql"
)

// nolint:gocyclo
func main() {
	c := dig.New()
	logger := logs.NewRawLogger()

	if config.LogsExportEnabled {
		assertions.AssertOk(c.Provide(func() *logs.Exporter {
			return logs.NewLogsExporter(logs.ExporterParams{
				StorageAddress:   config.LogsStorageAddress.Value(),
				ComponentName:    config.ComponentName,
				ComponentVersion: config.ComponentVersion,
			})
		}))

		assertions.AssertOk(c.Invoke(func(e *logs.Exporter) {
			if err := e.OpenConnection(); err != nil {
				logger.Fatal(err)
			}
		}))
		defer func() {
			assertions.AssertOk(c.Invoke(func(e *logs.Exporter) {
				if err := e.CloseConnection(); err != nil {
					logger.Fatal(err)
				}
			}))
		}()
	} else {
		assertions.AssertOk(c.Provide(func() *logs.Exporter {
			return nil //TODO
		}))
	}

	if config.RepositoriesEnabled {
		assertions.AssertOk(c.Provide(mysql.NewConnection))
		assertions.AssertOk(c.Invoke(func(conn *mysql.Connection) {
			if err := conn.Open(); err != nil {
				logger.Fatal(err)
			}
		}))
		defer func() {
			assertions.AssertOk(c.Invoke(func(conn *mysql.Connection) {
				if err := conn.Close(); err != nil {
					logger.Fatal(err)
				}
			}))
		}()

		assertions.AssertOk(c.Provide(
			mysql.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
		assertions.AssertOk(c.Provide(
			mysql.NewDocumentsInfoRepository,
			dig.As(new(repositories.DocumentsInfoRepository)),
		))
		assertions.AssertOk(c.Provide(
			memory.NewDocumentsContentRepository, //TODO
			dig.As(new(repositories.DocumentsContentRepository)),
		))
	} else {
		assertions.AssertOk(c.Provide(
			rfake.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
		assertions.AssertOk(c.Provide(
			rfake.NewDocumentsInfoRepository,
			dig.As(new(repositories.DocumentsInfoRepository)),
		))
		assertions.AssertOk(c.Provide( //TODO
			func() repositories.DocumentsContentRepository {
				return nil
			},
		))
	}

	if config.ServicesEnabled {
		assertions.AssertOk(c.Provide(
			impl.NewUsersService,
			dig.As(new(logic.UsersService)),
		))
		assertions.AssertOk(c.Provide(
			impl.NewDocumentsInfoService,
			dig.As(new(logic.DocumentsInfoService)),
		))
		assertions.AssertOk(c.Provide(
			impl.NewDocumentsContentService,
			dig.As(new(logic.DocumentsContentService)),
		))
	} else {
		assertions.AssertOk(c.Provide(
			sfake.NewUsersService,
			dig.As(new(logic.UsersService)),
		))
		assertions.AssertOk(c.Provide(
			sfake.NewDocumentsInfoService,
			dig.As(new(logic.DocumentsInfoService)),
		))
		assertions.AssertOk(c.Provide(
			sfake.NewDocumentsContentService,
			dig.As(new(logic.DocumentsContentService)),
		))
	}

	if config.QueuesEnabled {
		assertions.AssertOk(c.Provide(messages.NewMqClient))
		assertions.AssertOk(c.Invoke(func(c *messages.MqClient) {
			if err := c.OpenConnection(); err != nil {
				logger.Fatal(err)
			}
		}))
		defer func() {
			assertions.AssertOk(c.Invoke(func(c *messages.MqClient) {
				if err := c.CloseConnection(); err != nil {
					logger.Fatal(err)
				}
			}))
		}()

		assertions.AssertOk(c.Invoke(func(c *messages.MqClient) {
			if err := c.Run(); err != nil {
				logger.Fatal(err)
			}
		}))
	} else {
		assertions.AssertOk(c.Provide(func() *messages.MqClient {
			return nil //TODO
		}))
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM)

	assertions.AssertOk(c.Provide(api.NewServer))
	assertions.AssertOk(c.Invoke(func(srv *api.Server) {
		if err := srv.Run(); err != nil {
			logger.Fatal(err)
		}
	}))
	defer func() {
		assertions.AssertOk(c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				logger.Fatal(err)
			}
		}))
	}()

	<-ch
}

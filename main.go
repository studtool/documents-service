package main

import (
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/utils"

	"github.com/studtool/documents-service/api"
	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/logic/fake"
	"github.com/studtool/documents-service/logic/impl"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/repositories/fake"
	"github.com/studtool/documents-service/repositories/mysql"
)

func main() {
	c := dig.New()
	logger := logs.NewReflectLogger()

	if config.RepositoriesEnabled {
		utils.AssertOk(c.Provide(mysql.NewConnection))
		utils.AssertOk(c.Invoke(func(conn *mysql.Connection) {
			if err := conn.Open(); err != nil {
				logger.Fatal(err)
			}
		}))
		defer func() {
			utils.AssertOk(c.Invoke(func(conn *mysql.Connection) {
				if err := conn.Close(); err != nil {
					logger.Fatal(err)
				}
			}))
		}()

		utils.AssertOk(c.Provide(
			mysql.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
		utils.AssertOk(c.Provide(
			mysql.NewDocumentsInfoRepository,
			dig.As(new(repositories.DocumentsInfoRepository)),
		))
	} else {
		utils.AssertOk(c.Provide(
			rfake.NewUsersRepository,
			dig.As(new(repositories.UsersRepository)),
		))
		utils.AssertOk(c.Provide(
			rfake.NewDocumentsInfoRepository,
			dig.As(new(repositories.DocumentsInfoRepository)),
		))
	}

	if config.ServicesEnabled {
		utils.AssertOk(c.Provide(
			impl.NewUsersService,
			dig.As(new(logic.UsersService)),
		))
		utils.AssertOk(c.Provide(
			impl.NewDocumentsInfoService,
			dig.As(new(logic.DocumentsInfoService)),
		))
	} else {
		utils.AssertOk(c.Provide(
			sfake.NewUsersService,
			dig.As(new(logic.UsersService)),
		))
		utils.AssertOk(c.Provide(
			sfake.NewDocumentsInfoService,
			dig.As(new(logic.DocumentsInfoService)),
		))
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill)
	signal.Notify(ch, os.Interrupt)

	utils.AssertOk(c.Provide(api.NewServer))
	utils.AssertOk(c.Invoke(func(srv *api.Server) {
		if err := srv.Run(); err != nil {
			logger.Fatal(err)
		}
	}))
	defer func() {
		utils.AssertOk(c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				logger.Fatal(err)
			}
		}))
	}()

	<-ch
}

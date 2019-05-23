package main

import (
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/documents-service/api"
	"github.com/studtool/documents-service/beans"
	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/repositories/fs"
	"github.com/studtool/documents-service/repositories/mysql"
)

func main() {
	c := dig.New()

	if config.RepositoriesEnabled {
		utils.AssertOk(c.Provide(
			fs.NewDocumentsRepository,
			dig.As(new(repositories.DocumentsRepository)),
		))

		utils.AssertOk(c.Provide(mysql.NewConnection))
		utils.AssertOk(c.Invoke(func(conn *mysql.Connection) {
			if err := conn.Open(); err != nil {
				beans.Logger().Fatal(err.Error())
			} else {
				beans.Logger().Info("storage: connection open")
			}
		}))
		defer func() {
			utils.AssertOk(c.Invoke(func(conn *mysql.Connection) {
				if err := conn.Close(); err != nil {
					beans.Logger().Fatal(err)
				} else {
					beans.Logger().Info("storage: connection closed")
				}
			}))
		}()

		utils.AssertOk(c.Provide(
			mysql.NewDocumentsInfoRepository,
			dig.As(new(repositories.DocumentsInfoRepository)),
		))
		utils.AssertOk(c.Provide(
			mysql.NewPermissionsRepository,
			dig.As(new(repositories.PermissionsRepository)),
		))
	} else {
		utils.AssertOk(c.Provide(
			func() repositories.DocumentsRepository {
				return nil
			},
		))
		utils.AssertOk(c.Provide(
			func() repositories.DocumentsInfoRepository {
				return nil
			},
		))
		utils.AssertOk(c.Provide(
			func() repositories.PermissionsRepository {
				return nil
			},
		))
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill)
	signal.Notify(ch, os.Interrupt)

	utils.AssertOk(c.Provide(api.NewServer))
	utils.AssertOk(c.Invoke(func(srv *api.Server) {
		if err := srv.Run(); err != nil {
			beans.Logger().Fatal(err)
		}
	}))
	defer func() {
		utils.AssertOk(c.Invoke(func(srv *api.Server) {
			if err := srv.Shutdown(); err != nil {
				beans.Logger().Fatal(err)
			}
		}))
	}()

	<-ch
}

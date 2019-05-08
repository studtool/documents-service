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
	"github.com/studtool/documents-service/repositories/postgres"
)

func main() {
	c := dig.New()

	if config.RepositoriesEnabled.Value() {
		utils.AssertOk(c.Provide(
			fs.NewDocumentsRepository,
			dig.As(new(repositories.DocumentsRepository)),
		))

		utils.AssertOk(c.Provide(postgres.NewConnection))
		utils.AssertOk(c.Invoke(func(conn *postgres.Connection) {
			if err := conn.Open(); err != nil {
				beans.Logger().Fatal(err.Error())
			} else {
				beans.Logger().Info("storage: connection open")
			}
		}))
		defer func() {
			utils.AssertOk(c.Invoke(func(conn *postgres.Connection) {
				if err := conn.Close(); err != nil {
					beans.Logger().Fatal(err)
				} else {
					beans.Logger().Info("storage: connection closed")
				}
			}))
		}()

		utils.AssertOk(c.Provide(
			postgres.NewDocumentsInfoRepository,
			dig.As(new(repositories.DocumentsInfoRepository)),
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
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill)
	signal.Notify(ch, os.Interrupt)

	utils.AssertOk(c.Provide(api.NewServer))
	utils.AssertOk(c.Invoke(func(srv *api.Server) {
		go func() {
			if err := srv.Run(); err != nil {
				beans.Logger().Fatal(err)
				ch <- os.Interrupt
			}
		}()
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

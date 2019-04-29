package main

import (
	"os"
	"os/signal"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/documents-service/api"
	"github.com/studtool/documents-service/beans"
)

func main() {
	c := dig.New()

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

package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/studtool/common/logs"
)

type ServerConfig struct {
	Host string
	Port int
}

type Server struct {
	server *http.Server
	logger logs.Logger
}

func NewServer(c ServerConfig) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
		},
		logger: logs.NewStructLogger(
			logs.StructLoggerParams{
				Component: "common",
				Structure: "rest.Server",
			},
		),
	}
}

func (srv *Server) MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func (srv *Server) SetHandler(h http.Handler) {
	srv.server.Handler = h
}

func (srv *Server) Run() error {
	srv.logger.Infof("started on %s", srv.server.Addr)
	return srv.server.ListenAndServe()
}

func (srv *Server) Shutdown() error {
	srv.logger.Infof("stopped")
	return srv.server.Shutdown(context.TODO())
}

package rest

import (
	"context"
	"fmt"
	"net/http"

	//nolint:golint
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/studtool/common/config"
	"github.com/studtool/common/logs"
)

type ServerConfig struct {
	Host string
	Port int
}

type Server struct {
	server *http.Server

	structLogger  logs.Logger
	requestLogger logs.Logger

	apiClassifier APIClassifier
}

func NewServer(c ServerConfig) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
		},

		structLogger: logs.NewStructLogger(
			logs.StructLoggerParams{
				Component: config.Component,
				Structure: "rest.Server",
			},
		),

		requestLogger: logs.NewRequestLogger(
			logs.RequestLoggerParams{
				Component: config.Component,
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

// Optimization to seek public/protected/private/internal faster
type APIClassifier func(path string) string

func (srv *Server) SetAPIClassifier(c APIClassifier) {
	srv.apiClassifier = c
}

func (srv *Server) Run() error {
	srv.structLogger.Infof("started on %s", srv.server.Addr)
	go func() {
		if err := srv.server.ListenAndServe(); err != nil {
			srv.structLogger.Fatal(err)
		}
	}()
	return nil
}

func (srv *Server) Shutdown() error {
	srv.structLogger.Infof("stopped")
	return srv.server.Shutdown(context.TODO())
}

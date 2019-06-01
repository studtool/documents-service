package rest

import (
	"context"
	"fmt"
	"net/http"

	//nolint:golint
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

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

type ServerParams struct {
	Host string
	Port int

	ComponentName    string
	ComponentVersion string
}

func NewServer(params ServerParams) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(
				"%s:%d", params.Host, params.Port,
			),
		},

		structLogger: logs.NewStructLogger(
			logs.StructLoggerParams{
				ComponentName:     params.ComponentName,
				StructWithPkgName: "rest.Server",
			},
		),

		requestLogger: logs.NewRequestLogger(
			logs.RequestLoggerParams{
				Component: params.ComponentName,
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

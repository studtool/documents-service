package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/metrics"
)

type Server struct {
	server *http.Server

	structLogger  logs.Logger
	reflectLogger logs.Logger
	requestLogger logs.Logger

	apiClassifier APIClassifier

	metrics serverMetrics
}

type serverMetrics struct {
	rpsCounter *metrics.Counter
}

type ServerParams struct {
	Address string
	Handler http.Handler

	StructLogger  logs.Logger
	ReflectLogger logs.Logger
	RequestLogger logs.Logger

	APIClassifier APIClassifier
}

func NewServer(params ServerParams) *Server {
	return &Server{
		server: &http.Server{
			Addr:    params.Address,
			Handler: params.Handler,
		},

		structLogger:  params.StructLogger,
		reflectLogger: params.ReflectLogger,
		requestLogger: params.RequestLogger,

		apiClassifier: params.APIClassifier,

		metrics: serverMetrics{
			rpsCounter: metrics.NewCounter(metrics.CounterParams{
				Name:          "http_requests_per_second",
				ClearInterval: time.Second,
			}),
		},
	}
}

func (srv *Server) Run() error {
	srv.structLogger.Infof("started on %s", srv.server.Addr)
	go func() {
		if err := srv.server.ListenAndServe(); err != nil {
			srv.structLogger.Fatal(err)
		}
	}()

	srv.metrics.rpsCounter.Run()

	return nil
}

func (srv *Server) Shutdown() error {
	srv.structLogger.Infof("stopped")
	return srv.server.Shutdown(context.TODO())
}

package api

import (
	"github.com/studtool/documents-service/repositories"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/dig"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/logs"
	"github.com/studtool/common/rest"

	"github.com/studtool/documents-service/config"
)

type Server struct {
	server                  *rest.Server
	logger                  logs.Logger
	documentsInfoRepository repositories.DocumentsInfoRepository
}

type ServerParams struct {
	dig.In
	DocumentsInfoRepository repositories.DocumentsInfoRepository
}

func NewServer(params ServerParams) *Server {
	srv := &Server{
		server: rest.NewServer(
			rest.ServerConfig{
				Host: consts.EmptyString,
				Port: config.ServerPort.Value(),
			},
		),
		logger: logs.NewStructLogger(
			logs.StructLoggerParams{
				Component: "documents-service",
				Structure: "api.Server",
			},
		),
		documentsInfoRepository: params.DocumentsInfoRepository,
	}

	mx := mux.NewRouter()
	mx.Handle(`/api/protected/documents`, handlers.MethodHandler{
		http.MethodPost:   srv.server.WithAuth(http.HandlerFunc(srv.addDocument)),
		http.MethodGet:    srv.server.WithAuth(http.HandlerFunc(srv.getDocuments)),
		http.MethodDelete: srv.server.WithAuth(http.HandlerFunc(srv.deleteDocuments)),
	})
	mx.Handle(`/api/protected/documents/{document_id}`, handlers.MethodHandler{
		http.MethodDelete: srv.server.WithAuth(http.HandlerFunc(srv.deleteDocument)),
	})
	mx.Handle(`/api/protected/documents/{document_id}/info`, handlers.MethodHandler{
		http.MethodGet:   srv.server.WithAuth(http.HandlerFunc(srv.getDocumentInfo)),
		http.MethodPatch: srv.server.WithAuth(http.HandlerFunc(srv.updateDocumentInfo)),
	})
	mx.Handle(`/api/protected/documents/{document_id}/content`, handlers.MethodHandler{
		http.MethodGet:   srv.server.WithAuth(http.HandlerFunc(srv.getDocumentContent)),
		http.MethodPatch: srv.server.WithAuth(http.HandlerFunc(srv.updateDocumentContent)),
	})
	mx.Handle(`/metrics`, srv.server.MetricsHandler())

	h := srv.server.WithRecover(mx)
	if config.RequestsLogsEnabled.Value() {
		h = srv.server.WithLogs(h)
	}
	if config.CorsAllowed.Value() {
		h = srv.server.WithCORS(h, rest.CORS{
			Origins: []string{"*"},
			Methods: []string{
				http.MethodGet, http.MethodHead,
				http.MethodPost, http.MethodPatch,
				http.MethodDelete, http.MethodOptions,
			},
			Headers: []string{
				"User-Agent", "Authorization",
				"Content-Type", "Content-Length",
			},
			Credentials: false,
		})
	}

	srv.server.SetHandler(h)

	return srv
}

func (srv *Server) Run() error {
	err := srv.server.Run()
	if err == nil {
		srv.logger.Info("started")
	}
	return err
}

func (srv *Server) Shutdown() error {
	srv.logger.Info("shutdown")
	return srv.server.Shutdown()
}

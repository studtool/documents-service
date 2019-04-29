package api

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/rest"

	"github.com/studtool/documents-service/beans"
	"github.com/studtool/documents-service/config"
)

type Server struct {
	server *rest.Server
}

func NewServer() *Server {
	srv := &Server{
		server: rest.NewServer(
			rest.ServerConfig{
				Host: consts.EmptyString,
				Port: config.ServerPort.Value(),
			},
		),
	}

	mx := mux.NewRouter()
	mx.Handle(`/api/documents`, handlers.MethodHandler{
		http.MethodPost:   http.HandlerFunc(srv.addDocument),
		http.MethodGet:    http.HandlerFunc(srv.getDocuments),
		http.MethodDelete: http.HandlerFunc(srv.deleteDocuments),
	})
	mx.Handle(`/api/documents/{document_id}`, handlers.MethodHandler{
		http.MethodDelete: http.HandlerFunc(srv.deleteDocument),
	})
	mx.Handle(`/api/documents/{document_id}/info`, handlers.MethodHandler{
		http.MethodGet:   http.HandlerFunc(srv.getDocumentInfo),
		http.MethodPatch: http.HandlerFunc(srv.updateDocumentInfo),
	})
	mx.Handle(`/api/documents/{document_id}/content`, handlers.MethodHandler{
		http.MethodGet:   http.HandlerFunc(srv.getDocumentContent),
		http.MethodPatch: http.HandlerFunc(srv.updateDocumentContent),
	})

	srv.server.SetLogger(beans.Logger())

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
	return srv.server.Run()
}

func (srv *Server) Shutdown() error {
	return srv.server.Shutdown()
}

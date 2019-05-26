package api

import (
	"net/http"
	"strings"

	"github.com/go-http-utils/headers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/dig"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/logs"
	"github.com/studtool/common/rest"

	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type Server struct {
	server *rest.Server

	structLogger  logs.Logger
	reflectLogger logs.Logger

	documentsInfoService  logic.DocumentsInfoService
	permissionsRepository repositories.PermissionsRepository
}

type ServerParams struct {
	dig.In

	DocumentsInfoService logic.DocumentsInfoService
}

func NewServer(params ServerParams) *Server {
	srv := &Server{
		server: rest.NewServer(
			rest.ServerConfig{
				Host: consts.EmptyString,
				Port: config.ServerPort.Value(),
			},
		),

		documentsInfoService: params.DocumentsInfoService,
	}

	srv.structLogger = srvutils.MakeStructLogger(srv)
	srv.reflectLogger = srvutils.MakeReflectLogger(srv)

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
				headers.Authorization, headers.UserAgent,
				headers.ContentType, headers.ContentLength,
				headers.ContentEncoding, headers.ContentLanguage,
			},
			Credentials: false,
		})
	}

	srv.server.SetHandler(h)
	srv.server.SetAPIClassifier(srv.makeApiClassifier())

	srv.structLogger.Info("initialized")

	return srv
}

func (srv *Server) Run() error {
	err := srv.server.Run()
	if err == nil {
		srv.structLogger.Info("start")
	}
	return err
}

func (srv *Server) Shutdown() error {
	srv.structLogger.Info("shutdown")
	return srv.server.Shutdown()
}

func (srv *Server) makeApiClassifier() rest.APIClassifier {
	apiTypeIndex := len("/api/")

	return rest.APIClassifier(func(path string) string {
		path = path[apiTypeIndex:]

		if strings.HasPrefix(path, rest.APITypePublic) {
			return rest.APITypePublic
		}
		if strings.HasPrefix(path, rest.APITypeProtected) {
			return rest.APITypeProtected
		}
		if strings.HasPrefix(path, rest.APITypePublic) {
			return rest.APITypePrivate
		}
		if strings.HasPrefix(path, rest.APITypeInternal) {
			return rest.APITypeInternal
		}

		return consts.EmptyString
	})
}

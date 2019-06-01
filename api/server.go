package api

import (
	"fmt"
	"net/http"

	"go.uber.org/dig"

	"github.com/go-http-utils/headers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/rest"

	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/utils"
)

type Server struct {
	rest.Server

	structLogger  logs.Logger
	reflectLogger logs.Logger

	documentsInfoService    logic.DocumentsInfoService
	documentsContentService logic.DocumentsContentService
}

type ServerParams struct {
	dig.In

	DocumentsInfoService    logic.DocumentsInfoService
	DocumentsContentService logic.DocumentsContentService
}

func NewServer(params ServerParams) *Server {
	srv := &Server{
		documentsInfoService:    params.DocumentsInfoService,
		documentsContentService: params.DocumentsContentService,
	}

	srv.structLogger = srvutils.MakeStructLogger(srv)
	srv.reflectLogger = srvutils.MakeReflectLogger(srv)

	v := rest.ParseAPIVersion(config.ComponentVersion)
	srvPath := rest.MakeAPIPath(v, rest.APITypeProtected, "/documents")

	mx := mux.NewRouter()

	mx.Handle(srvPath, handlers.MethodHandler{
		http.MethodPost:   srv.WithAuth(http.HandlerFunc(srv.addDocument)),
		http.MethodGet:    srv.WithAuth(http.HandlerFunc(srv.getDocumentsInfo)),
		http.MethodDelete: srv.WithAuth(http.HandlerFunc(srv.deleteDocuments)),
	})
	mx.Handle(srvPath+"/{document_id}", handlers.MethodHandler{
		http.MethodDelete: srv.WithAuth(http.HandlerFunc(srv.deleteDocument)),
	})
	mx.Handle(srvPath+"/{document_id}/info", handlers.MethodHandler{
		http.MethodGet:   srv.WithAuth(http.HandlerFunc(srv.getDocumentInfo)),
		http.MethodPatch: srv.WithAuth(http.HandlerFunc(srv.updateDocumentInfo)),
	})
	mx.Handle(srvPath+"/{document_id}/content", handlers.MethodHandler{
		http.MethodGet:   srv.WithAuth(http.HandlerFunc(srv.getDocumentContent)),
		http.MethodPatch: srv.WithAuth(http.HandlerFunc(srv.updateDocumentContent)),
	})
	mx.Handle(`/pprof`, rest.GetProfilerHandler())
	mx.Handle(`/metrics`, rest.GetMetricsHandler())

	reqHandler := srv.WithRecover(mx)
	if config.RequestsLogsEnabled.Value() {
		reqHandler = srv.WithLogs(reqHandler)
	}
	if config.CorsAllowed.Value() {
		reqHandler = srv.WithCORS(reqHandler, rest.CORS{
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

	srv.structLogger = srvutils.MakeStructLogger(srv)
	srv.reflectLogger = srvutils.MakeReflectLogger(srv)

	srv.Server = *rest.NewServer(
		rest.ServerParams{
			Address: fmt.Sprintf(":%d", config.ServerPort.Value()),
			Handler: reqHandler,

			StructLogger:  srv.structLogger,
			ReflectLogger: srv.reflectLogger,
			RequestLogger: srvutils.MakeRequestLogger(srv),

			APIClassifier: rest.NewPathAPIClassifier(),
		},
	)
	srv.structLogger.Info("initialized")

	return srv
}

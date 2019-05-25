package api

import (
	"net/http"

	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/models"
)

func (srv *Server) addDocument(w http.ResponseWriter, r *http.Request) {
	documentInfo := &models.DocumentInfo{}
	if err := srv.server.ParseBodyJSON(documentInfo, r); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	params := logic.AddDocumentInfoParams{
		UserID:   types.ID(srv.server.ParseUserID(r)),
		Document: documentInfo,
	}
	if err := srv.documentsInfoService.AddDocumentInfo(params); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOkJSON(w, documentInfo)
}

func (srv *Server) getDocuments(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

func (srv *Server) deleteDocuments(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

func (srv *Server) deleteDocument(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

func (srv *Server) getDocumentInfo(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

func (srv *Server) updateDocumentInfo(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

func (srv *Server) getDocumentContent(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

func (srv *Server) updateDocumentContent(w http.ResponseWriter, r *http.Request) {
	srv.server.WriteNotImplemented(w) //TODO
}

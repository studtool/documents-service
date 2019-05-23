package api

import (
	"github.com/studtool/common/errs"
	"net/http"

	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

func (srv *Server) addDocument(w http.ResponseWriter, r *http.Request) {
	documentInfo := &models.DocumentInfoFull{}
	if err := srv.server.ParseBodyJSON(&documentInfo.DocumentInfo, r); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	documentInfo.OwnerID = types.ID(srv.server.ParseUserID(r))
	if err := srv.documentsInfoRepository.AddDocumentInfo(documentInfo); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOkJSON(w, &documentInfo.DocumentInfo)
}

func (srv *Server) getDocuments(w http.ResponseWriter, r *http.Request) {
	ownerID, err := srv.parseOwnerID(r)
	if err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	if ownerID != srv.server.ParseUserID(r) {
		srv.server.WriteErrJSON(w, errs.NewNotAuthorizedError("not authorized")) //TODO
		return
	}

	subject, err := srv.parseSubject(r)
	if err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	page := srv.parsePage(r)

	documents, err := srv.documentsInfoRepository.GetDocumentsInfo(
		ownerID, subject, page,
	)
	if err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOkJSON(w, &documents)
}

func (srv *Server) deleteDocuments(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Server) deleteDocument(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Server) getDocumentInfo(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Server) updateDocumentInfo(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Server) getDocumentContent(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Server) updateDocumentContent(w http.ResponseWriter, r *http.Request) {
	//TODO
}

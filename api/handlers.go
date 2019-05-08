package api

import (
	"net/http"

	"github.com/studtool/documents-service/models"
)

func (srv *Server) addDocument(w http.ResponseWriter, r *http.Request) {
	documentInfo := &models.DocumentInfoFull{}
	if err := srv.server.ParseBodyJSON(&documentInfo.DocumentInfo, r); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	if err := srv.documentsInfoRepository.SaveDocumentInfo(documentInfo); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	if err := srv.documentsRepository.AddDocument(documentInfo.Id, []byte{}); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	srv.server.WriteOk(w)
}

func (srv *Server) getDocuments(w http.ResponseWriter, r *http.Request) {
	//TODO
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

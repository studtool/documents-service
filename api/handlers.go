package api

import (
	"net/http"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/models"
)

func (srv *Server) addDocument(w http.ResponseWriter, r *http.Request) {
	documentInfo := &models.DocumentInfo{}
	if err := srv.ParseBodyJSON(documentInfo, r); err != nil {
		srv.WriteErrJSON(w, err)
		return
	}

	params := logic.AddDocumentInfoParams{
		UserID:       srv.parseHeaderUserID(r),
		DocumentInfo: documentInfo,
	}
	if err := srv.documentsInfoService.AddDocumentInfo(params); err != nil {
		srv.WriteErrJSON(w, err)
		return
	}

	srv.WriteOkJSON(w, documentInfo)
}

func (srv *Server) getDocumentsInfo(w http.ResponseWriter, r *http.Request) {
	var documentsInfo models.DocumentsInfo

	params := logic.GetDocumentsInfoParams{
		UserID:        srv.parseHeaderUserID(r),
		OwnerID:       srv.parseParamOwnerID(r),
		Subject:       srv.parseParamSubject(r),
		Page:          srv.parseParamsPage(r),
		DocumentsInfo: &documentsInfo,
	}
	if err := srv.documentsInfoService.GetDocumentsInfo(params); err != nil {
		srv.WriteErrJSON(w, err)
		return
	}

	srv.WriteOkJSON(w, documentsInfo)
}

func (srv *Server) getDocumentInfo(w http.ResponseWriter, r *http.Request) {
	srv.WriteNotImplemented(w) //TODO
}

func (srv *Server) deleteDocuments(w http.ResponseWriter, r *http.Request) {
	srv.WriteNotImplemented(w) //TODO
}

func (srv *Server) deleteDocument(w http.ResponseWriter, r *http.Request) {
	srv.WriteNotImplemented(w) //TODO
}

func (srv *Server) updateDocumentInfo(w http.ResponseWriter, r *http.Request) {
	srv.WriteNotImplemented(w) //TODO
}

func (srv *Server) getDocumentContent(w http.ResponseWriter, r *http.Request) {
	var content models.DocumentContent

	params := logic.GetDocumentContentParams{
		UserID:          srv.parseHeaderUserID(r),
		DocumentID:      srv.parsePathDocumentID(r),
		DocumentContent: &content,
	}
	if err := srv.documentsContentService.GetDocumentContent(&params); err != nil {
		srv.WriteErrJSON(w, err)
		return
	}

	srv.WriteOkRaw(w, content)
}

func (srv *Server) updateDocumentContent(w http.ResponseWriter, r *http.Request) {
	content, err := srv.parseBodyDocumentContent(r)
	if err != nil {
		srv.WriteErrJSON(w, err)
		return
	}

	params := logic.UpdateDocumentContentParams{
		UserID:          srv.parseHeaderUserID(r),
		DocumentID:      srv.parsePathDocumentID(r),
		DocumentContent: &content,
	}
	if err := srv.documentsContentService.UpdateDocumentContent(&params); err != nil {
		srv.WriteErrJSON(w, err)
		return
	}

	srv.WriteOk(w)
}

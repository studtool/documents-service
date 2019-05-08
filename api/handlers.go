package api

import (
	"github.com/google/uuid"
	"math/rand"
	"net/http"

	"github.com/studtool/documents-service/models"
)

func (srv *Server) addDocument(w http.ResponseWriter, r *http.Request) {
	documentInfo := &models.DocumentInfoFull{}
	if err := srv.server.ParseBodyJSON(&documentInfo.DocumentInfo, r); err != nil {
		srv.server.WriteErrJSON(w, err)
		return
	}

	if err := srv.documentsInfoRepository.AddDocumentInfo(documentInfo); err != nil {
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
	genId := func() string {
		v, _ := uuid.NewRandom()
		return v.String()
	}

	documents := make([]models.DocumentInfo, 0)
	for i := 0; i < 50; i++ {
		id := genId()

		m := models.DocumentInfo{
			Id:      id,
			Title:   "Title" + id,
			OwnerId: genId(),
			Subject: "The subject",
			Meta: models.DocumentMeta{
				Size: rand.Int63(),
			},
		}

		documents = append(documents, m)
	}

	srv.server.WriteBodyJSON(w, http.StatusOK, models.DocumentsInfo(documents))
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

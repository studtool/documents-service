package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
)

const (
	ownerIDParamName   = "owner_id"
	subjectParamName   = "subject"
	pageIndexParamName = "page"
	pageSizeParamName  = "size"
)

const (
	defaultPageIndex = 0
	defaultPageSize  = 10
)

const (
	documentIDVarName = "document_id"
)

func (srv *Server) parseHeaderUserID(r *http.Request) types.ID {
	return types.ID(srv.server.ParseUserID(r))
}

func (srv *Server) parsePathDocumentID(r *http.Request) types.ID {
	return types.ID(mux.Vars(r)[documentIDVarName])
}

func (srv *Server) parseParamOwnerID(r *http.Request) types.ID {
	return types.ID(r.URL.Query().Get(ownerIDParamName))
}

func (srv *Server) parseParamSubject(r *http.Request) string {
	return r.URL.Query().Get(subjectParamName)
}

func (srv *Server) parseParamsPage(r *http.Request) repositories.Page {
	return repositories.Page{
		Index: srv.parsePageIndex(r),
		Size:  srv.parsePageSize(r),
	}
}

func (srv *Server) parseBodyDocumentContent(r *http.Request) (models.DocumentContent, *errs.Error) {
	content, err := srv.server.GetRawBody(r)
	return models.DocumentContent(content), err
}

func (srv *Server) parsePageIndex(r *http.Request) int32 { //TODO remove logic
	indexStr := r.URL.Query().Get(pageIndexParamName)
	if indexStr == consts.EmptyString {
		return defaultPageIndex
	} else {
		if p, err := strconv.ParseInt(indexStr, 10, 32); err == nil {
			if p >= 0 {
				return int32(p)
			} else {
				return defaultPageIndex
			}
		} else {
			return defaultPageIndex
		}
	}
}

func (srv *Server) parsePageSize(r *http.Request) int32 { //TODO remove logic
	indexStr := r.URL.Query().Get(pageSizeParamName)
	if indexStr == consts.EmptyString {
		return defaultPageSize
	} else {
		if p, err := strconv.ParseInt(indexStr, 10, 32); err == nil {
			if p >= 0 {
				return int32(p)
			} else {
				return defaultPageSize
			}
		} else {
			return defaultPageSize
		}
	}
}

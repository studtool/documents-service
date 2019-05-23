package api

import (
	"net/http"
	"strconv"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"

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

func (srv *Server) parseOwnerID(r *http.Request) (string, *errs.Error) {
	id := r.URL.Query().Get(ownerIDParamName)
	if id == consts.EmptyString {
		return consts.EmptyString, errs.NewBadFormatError("no owner_id") //TODO
	}
	return id, nil
}

func (srv *Server) parseSubject(r *http.Request) (*string, *errs.Error) { //TODO no error, empty subject
	s := new(string)
	*s = r.URL.Query().Get(subjectParamName)
	if *s == consts.EmptyString {
		return nil, errs.NewBadFormatError("no subject") //TODO
	}
	return s, nil
}

func (srv *Server) parsePage(r *http.Request) repositories.Page {
	return repositories.Page{
		Index: srv.parsePageIndex(r),
		Size:  srv.parsePageSize(r),
	}
}

func (srv *Server) parsePageIndex(r *http.Request) int32 {
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

func (srv *Server) parsePageSize(r *http.Request) int32 {
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

package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type DocumentsContentService struct {
	documentsInfoRepository    repositories.DocumentsInfoRepository
	documentsContentRepository repositories.DocumentsContentRepository

	structLogger  logs.Logger
	reflectLogger logs.Logger
}

type DocumentsContentServiceParams struct {
	dig.In

	DocumentsInfoRepository    repositories.DocumentsInfoRepository
	DocumentsContentRepository repositories.DocumentsContentRepository
}

func NewDocumentsContentService(params DocumentsContentServiceParams) *DocumentsContentService {
	r := &DocumentsContentService{
		documentsInfoRepository:    params.DocumentsInfoRepository,
		documentsContentRepository: params.DocumentsContentRepository,
	}

	r.structLogger = srvutils.MakeStructLogger(r)
	r.reflectLogger = srvutils.MakeReflectLogger(r)

	r.structLogger.Info("initialized")

	return r
}

func (s *DocumentsContentService) GetDocumentContent(params *logic.GetDocumentContentParams) *errs.Error {
	panic("implement me")
}

func (s *DocumentsContentService) UpdateDocumentContent(params *logic.UpdateDocumentContentParams) *errs.Error {
	panic("implement me")
}

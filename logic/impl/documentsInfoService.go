package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/repositories"
)

type DocumentsInfoService struct {
	logger logs.Logger

	usersRepository         repositories.UsersRepository
	documentsInfoRepository repositories.DocumentsInfoRepository
}

type DocumentsInfoServiceParams struct {
	dig.In

	UsersRepository         repositories.UsersRepository
	DocumentsInfoRepository repositories.DocumentsInfoRepository
}

func NewDocumentsInfoService(params DocumentsInfoServiceParams) *DocumentsInfoService {
	return &DocumentsInfoService{
		logger: logs.NewStructLogger(logs.StructLoggerParams{
			Component: config.Component,
			Structure: "impl.DocumentsInfoService",
		}),

		usersRepository:         params.UsersRepository,
		documentsInfoRepository: params.DocumentsInfoRepository,
	}
}

func (s *DocumentsInfoService) AddDocumentInfo(params logic.AddDocumentInfoParams) *errs.Error {
	panic("implement me")
}

func (s *DocumentsInfoService) GetDocumentInfo(params logic.GetDocumentInfoParams) *errs.Error {
	panic("implement me")
}

func (s *DocumentsInfoService) UpdateDocumentTitle(params logic.UpdateDocumentTitleParams) *errs.Error {
	panic("implement me")
}

func (s *DocumentsInfoService) UpdateDocumentSubject(params logic.UpdateDocumentSubjectParams) *errs.Error {
	panic("implement me")
}

func (s *DocumentsInfoService) DeleteDocument(params logic.DeleteDocumentParams) *errs.Error {
	panic("implement me")
}

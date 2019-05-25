package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type DocumentsInfoService struct {
	structLogger  logs.Logger
	reflectLogger logs.Logger

	usersRepository         repositories.UsersRepository
	documentsInfoRepository repositories.DocumentsInfoRepository
}

type DocumentsInfoServiceParams struct {
	dig.In

	UsersRepository         repositories.UsersRepository
	DocumentsInfoRepository repositories.DocumentsInfoRepository
}

func NewDocumentsInfoService(params DocumentsInfoServiceParams) *DocumentsInfoService {
	s := &DocumentsInfoService{
		usersRepository:         params.UsersRepository,
		documentsInfoRepository: params.DocumentsInfoRepository,
	}

	s.structLogger = utils.MakeStructLogger(s)
	s.reflectLogger = utils.MakeReflectLogger(s)

	s.structLogger.Info("initialized")

	return s
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

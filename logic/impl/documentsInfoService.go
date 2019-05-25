package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type DocumentsInfoService struct {
	usersRepository         repositories.UsersRepository
	documentsInfoRepository repositories.DocumentsInfoRepository

	structLogger  logs.Logger
	reflectLogger logs.Logger

	readPermissionErr  *errs.Error
	writePermissionErr *errs.Error
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

		// TODO these errors should be 404, but for dev they are 403
		readPermissionErr:  errs.NewPermissionDeniedError("permission to read document denied"),
		writePermissionErr: errs.NewPermissionDeniedError("permission to write document denied"),
	}

	s.structLogger = srvutils.MakeStructLogger(s)
	s.reflectLogger = srvutils.MakeReflectLogger(s)

	s.structLogger.Info("initialized")

	return s
}

func (s *DocumentsInfoService) AddDocumentInfo(params logic.AddDocumentInfoParams) *errs.Error {
	docInfo := params.DocumentInfo

	docInfo.OwnerID = params.UserID
	docInfo.DocumentID = srvutils.MakeID()

	err := s.usersRepository.CheckUserExistsByID(params.UserID)
	if err != nil {
		if err.Type == errs.NotFound {
			s.structLogger.Error("user [id = %s] should exist", params.UserID)
		}
		return err
	}

	err = s.documentsInfoRepository.AddDocumentInfo(docInfo)
	if err == nil {
		s.structLogger.Infof("document [id = %s] added", docInfo.DocumentID)
	}

	return nil
}

func (s *DocumentsInfoService) GetDocumentInfo(params logic.GetDocumentInfoParams) *errs.Error {
	docInfo := params.DocumentInfo

	if params.UserID != docInfo.OwnerID {
		return s.readPermissionErr
	}

	return s.documentsInfoRepository.GetDocumentInfoByID(docInfo)
}

func (s *DocumentsInfoService) GetDocumentsInfo(params logic.GetDocumentsInfoParams) *errs.Error {
	docsInfo := params.DocumentsInfo

	if params.UserID != params.OwnerID {
		return s.readPermissionErr
	}

	var err *errs.Error
	if params.Subject == consts.EmptyString {
		*docsInfo, err = s.documentsInfoRepository.
			GetDocumentsInfoByOwnerID(params.OwnerID, params.Page)
	} else {
		*docsInfo, err = s.documentsInfoRepository.
			GetDocumentsInfoByOwnerIDAndSubject(params.OwnerID, params.Subject, params.Page)
	}

	return err
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

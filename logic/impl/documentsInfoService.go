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

	LogsExporter *logs.Exporter
}

func NewDocumentsInfoService(params DocumentsInfoServiceParams) *DocumentsInfoService {
	s := &DocumentsInfoService{
		usersRepository:         params.UsersRepository,
		documentsInfoRepository: params.DocumentsInfoRepository,

		readPermissionErr:  errs.NewPermissionDeniedError("permission to read document denied"),
		writePermissionErr: errs.NewPermissionDeniedError("permission to write document denied"),
	}

	p := srvutils.LoggerParams{
		Value:    s,
		Exporter: params.LogsExporter,
	}

	s.structLogger = srvutils.MakeStructLogger(p)
	s.reflectLogger = srvutils.MakeReflectLogger(p)

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
			s.structLogger.Errorf("user [id = '%s'] should exist", params.UserID)
		}
		return err
	}

	err = s.documentsInfoRepository.AddDocumentInfo(docInfo)
	if err == nil {
		s.structLogger.Infof("document [id = '%s'] added", docInfo.DocumentID)
	}

	return nil
}

func (s *DocumentsInfoService) GetDocumentInfo(params logic.GetDocumentInfoParams) *errs.Error {
	docInfo := params.DocumentInfo

	if params.UserID != docInfo.OwnerID {
		return s.readPermissionErr
	}

	err := s.documentsInfoRepository.GetDocumentInfoByID(docInfo)
	if err != nil {
		if err.Type == errs.NotFound {
			s.structLogger.Warningf("document [id = '%s'] not found", params.DocumentInfo.DocumentID)
		}
	}

	return err
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
	if err != nil {
		if err.Type == errs.NotFound {
			s.structLogger.Warningf("documents [owner_id = '%s'] not found", params.UserID)
		}
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

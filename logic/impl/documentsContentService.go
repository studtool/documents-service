package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type DocumentsContentService struct {
	documentsInfoRepository    repositories.DocumentsInfoRepository
	documentsContentRepository repositories.DocumentsContentRepository

	documentNotFoundErr     *errs.Error
	documentAccessDeniedErr *errs.Error

	structLogger  logs.Logger
	reflectLogger logs.Logger
}

type DocumentsContentServiceParams struct {
	dig.In

	DocumentsInfoRepository    repositories.DocumentsInfoRepository
	DocumentsContentRepository repositories.DocumentsContentRepository

	LogsExporter *logs.Exporter
}

func NewDocumentsContentService(params DocumentsContentServiceParams) *DocumentsContentService {
	r := &DocumentsContentService{
		documentsInfoRepository:    params.DocumentsInfoRepository,
		documentsContentRepository: params.DocumentsContentRepository,
		documentNotFoundErr:        errs.NewNotFoundError("document not found"),
		documentAccessDeniedErr:    errs.NewPermissionDeniedError("document access denied"),
	}

	p := srvutils.LoggerParams{
		Value:    r,
		Exporter: params.LogsExporter,
	}

	r.structLogger = srvutils.MakeStructLogger(p)
	r.reflectLogger = srvutils.MakeReflectLogger(p)

	r.structLogger.Info("initialized")

	return r
}

func (s *DocumentsContentService) GetDocumentContent(params *logic.GetDocumentContentParams) *errs.Error {
	documentInfo := models.DocumentInfo{
		DocumentID: params.DocumentID,
	}
	if err := s.documentsInfoRepository.GetDocumentInfoByID(&documentInfo); err != nil {
		if err.Type == errs.NotFound {
			s.structLogger.Warningf("document [id = '%s'] not found", documentInfo.DocumentID)
		}
		return err
	}

	if params.UserID != documentInfo.OwnerID {
		s.structLogger.Warningf(
			"document [id = '%s'] access forbidden [user_id = '%s'; scope = 'read']",
			documentInfo.DocumentID, params.UserID,
		)
		return s.documentAccessDeniedErr
	}

	err := s.documentsContentRepository.
		GetDocumentContent(params.DocumentID, params.DocumentContent)
	if err != nil {
		if err.Type == errs.NotFound {
			*params.DocumentContent = make(models.DocumentContent, 0) //TODO
		} else {
			return err
		}
	}

	return nil
}

func (s *DocumentsContentService) UpdateDocumentContent(params *logic.UpdateDocumentContentParams) *errs.Error {
	documentInfo := models.DocumentInfo{
		DocumentID: params.DocumentID,
	}
	if err := s.documentsInfoRepository.GetDocumentInfoByID(&documentInfo); err != nil {
		if err.Type == errs.NotFound {
			s.structLogger.Warningf("document [id = '%s'] not found", documentInfo.DocumentID)
		}
		return err
	}

	if params.UserID != documentInfo.OwnerID {
		s.structLogger.Warningf(
			"document [id = '%s'] access forbidden [user_id = '%s'; scope = 'write']",
			params.DocumentID, params.UserID,
		)
		return s.documentAccessDeniedErr
	}

	return s.documentsContentRepository.
		UpdateDocumentContent(params.DocumentID, params.DocumentContent)
}

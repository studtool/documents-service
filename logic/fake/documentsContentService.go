package sfake

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/utils"
)

type DocumentsContentService struct {
	structLogger logs.Logger
}

func NewDocumentsContentService() *DocumentsContentService {
	s := &DocumentsContentService{}

	s.structLogger = srvutils.MakeStructLogger(s)
	s.structLogger.Warning("initialized")

	return s
}

const (
	FakeDocumentContent = "THIS IS THE DOCUMENT"
)

func (s *DocumentsContentService) GetDocumentContent(params *logic.GetDocumentContentParams) *errs.Error {
	*params.DocumentContent = []byte(FakeDocumentContent)

	s.structLogger.Warningf(
		"GetDocumentContent() called [user_id = '%s'; document_id = '%s'", params.UserID, params.DocumentID,
	)

	return nil
}

func (s *DocumentsContentService) UpdateDocumentContent(params *logic.UpdateDocumentContentParams) *errs.Error {
	s.structLogger.Warningf(
		"UpdateDocumentContent() called [user_id = '%s'; document_id = '%s'", params.UserID, params.DocumentID,
	)

	return nil
}

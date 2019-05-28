package impl

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/logic"
)

type DocumentsContentService struct{}

func NewDocumentsContentService() *DocumentsContentService {
	return &DocumentsContentService{}
}

func (*DocumentsContentService) GetDocumentContent(params *logic.GetDocumentContentParams) *errs.Error {
	panic("implement me")
}

func (*DocumentsContentService) UpdateDocumentContent(params *logic.UpdateDocumentContentParams) *errs.Error {
	panic("implement me")
}

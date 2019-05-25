package rfake

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
)

type DocumentsInfoRepository struct{}

func NewDocumentsInfoRepository() *DocumentsInfoRepository {
	return &DocumentsInfoRepository{}
}

func (r *DocumentsInfoRepository) AddDocumentInfo(info *models.DocumentInfo) *errs.Error {
	return errs.NewNotImplementedError(r.AddDocumentInfo)
}

func (r *DocumentsInfoRepository) GetDocumentInfoByID(info *models.DocumentInfo) *errs.Error {
	return errs.NewNotImplementedError(r.GetDocumentInfoByID)
}

func (r *DocumentsInfoRepository) GetDocumentsInfoByOwnerID(
	ownerID types.ID, page repositories.Page) (models.DocumentsInfo, *errs.Error) {

	return nil, errs.NewNotImplementedError(r.GetDocumentsInfoByOwnerID)
}

func (r *DocumentsInfoRepository) GetDocumentsInfoByOwnerIDAndSubject(
	ownerID types.ID, subject string, page repositories.Page) (models.DocumentsInfo, *errs.Error) {

	return nil, errs.NewNotImplementedError(r.GetDocumentsInfoByOwnerIDAndSubject)
}

func (r *DocumentsInfoRepository) UpdateDocumentTitleByID(update *models.DocumentTitleUpdate) *errs.Error {
	return errs.NewNotImplementedError(r.UpdateDocumentTitleByID)
}

func (r *DocumentsInfoRepository) UpdateDocumentSubjectByID(update *models.DocumentSubjectUpdate) *errs.Error {
	return errs.NewNotImplementedError(r.UpdateDocumentSubjectByID)
}

func (r *DocumentsInfoRepository) DeleteDocumentByID(documentID types.ID) *errs.Error {
	return errs.NewNotImplementedError(r.DeleteDocumentByID)
}

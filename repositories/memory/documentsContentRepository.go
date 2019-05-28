package memory

import (
	"sync"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/utils"
)

type DocumentsContentRepository struct {
	rwMutex        *sync.RWMutex
	documents      documentsContentMap
	docNotFoundErr *errs.Error
	structLogger   logs.Logger
}

func NewDocumentsContentRepository() *DocumentsContentRepository {
	r := &DocumentsContentRepository{
		rwMutex:        &sync.RWMutex{},
		documents:      make(documentsContentMap),
		docNotFoundErr: errs.NewNotFoundError("document not found"),
	}

	r.structLogger = srvutils.MakeStructLogger(r)
	r.structLogger.Info("initialized")

	return r
}

func (r *DocumentsContentRepository) GetDocumentContent(documentID types.ID, content *models.DocumentContent) *errs.Error {
	rContent := make(models.DocumentContent, len(*content))
	copy(rContent, *content)

	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	r.documents[documentID] = rContent

	return nil
}

func (r *DocumentsContentRepository) UpdateDocumentContent(documentID types.ID, content *models.DocumentContent) *errs.Error {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()

	rContent, ok := r.documents[documentID]
	if !ok {
		return r.docNotFoundErr
	}

	*content = make(models.DocumentContent, len(rContent))
	copy(*content, rContent)

	return nil
}

type documentsContentMap map[types.ID]models.DocumentContent

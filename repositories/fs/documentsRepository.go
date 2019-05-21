package fs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/beans"
)

const (
	rwPermission = 0644
)

type DocumentsRepository struct {
	docsDir string
}

func NewDocumentsRepository() *DocumentsRepository {
	return &DocumentsRepository{}
}

func (r *DocumentsRepository) AddDocument(documentId string, data []byte) *errs.Error {
	f, err := os.Create(r.makeFilePath(documentId))
	if err != nil {
		return errs.New(err)
	}
	defer r.closeFileWithCheck(f)
	return nil
}

func (r *DocumentsRepository) GetDocument(documentId string) ([]byte, *errs.Error) {
	f, err := os.Open(r.makeFilePath(documentId))
	if err != nil {
		return nil, errs.NewNotFoundError(err.Error()) //TODO
	}
	defer r.closeFileWithCheck(f)

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errs.New(err)
	}

	return data, nil
}

func (r *DocumentsRepository) UpdateDocument(documentId string, data []byte) *errs.Error {
	f, err := os.OpenFile(r.makeFilePath(documentId), os.O_RDWR, rwPermission)
	if err != nil {
		return errs.NewNotFoundError(err.Error()) //TODO
	}
	defer r.closeFileWithCheck(f)

	err = f.Truncate(int64(len(data)))
	if err != nil {
		return errs.New(err)
	}

	if _, err := f.Write(data); err != nil {
		return errs.New(err)
	}

	return nil
}

func (r *DocumentsRepository) DeleteDocument(documentId string) *errs.Error {
	if err := os.Remove(r.makeFilePath(documentId)); err != nil {
		return errs.NewNotFoundError(err.Error()) //TODO
	}
	return nil
}

func (r *DocumentsRepository) makeFilePath(documentId string) string {
	return fmt.Sprintf("%s%v%s", r.docsDir, os.PathSeparator, documentId)
}

func (r *DocumentsRepository) closeFileWithCheck(f *os.File) {
	err := f.Close()
	if err != nil {
		beans.Logger().Error(err)
	}
}

package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type DocumentInfo struct {
	DocumentID types.ID `json:"documentId"`
	Title      string   `json:"title"`
	OwnerID    types.ID `json:"ownerId"`
	Subject    string   `json:"subject"`
}

//easyjson:json
type DocumentTitleUpdate struct {
	DocumentID types.ID `json:"-"`
	NewTitle   string   `json:"newTitle"`
}

//easyjson:json
type DocumentSubjectUpdate struct {
	DocumentID types.ID `json:"-"`
	NewSubject string   `json:"newSubject"`
}

//easyjson:json
type DocumentsInfo []DocumentInfo

//easyjson:json
type UpdateInfo struct {
	UserID    types.ID       `json:"userId"`
	Timestamp types.DateTime `json:"timestamp"`
}

package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type DocumentInfo struct {
	DocumentID types.ID `msg:"documentId" json:"documentId"`
	Title      string   `msg:"title"      json:"title"`
	OwnerID    types.ID `msg:"ownerId"    json:"ownerId"`
	Subject    string   `msg:"subject"    json:"subject"`
}

//easyjson:json
type DocumentTitleUpdate struct {
	DocumentID types.ID `msg:"-"        json:"-"`
	NewTitle   string   `msg:"newTitle" json:"newTitle"`
}

//easyjson:json
type DocumentSubjectUpdate struct {
	DocumentID types.ID `msg:"-"          json:"-"`
	NewSubject string   `msg:"newSubject" json:"newSubject"`
}

//easyjson:json
type DocumentsInfo []DocumentInfo

//easyjson:json
type UpdateInfo struct {
	UserID    types.ID       `msg:"userId"    json:"userId"`
	Timestamp types.DateTime `msg:"timestamp" json:"timestamp"`
}

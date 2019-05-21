package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type DocumentInfo struct {
	Id      string       `json:"id"`
	Title   string       `json:"title"`
	OwnerId string       `json:"ownerId"`
	Subject string       `json:"subject"`
	Meta    DocumentMeta `json:"meta"`
}

//easyjson:json
type DocumentInfoFull struct {
	DocumentInfo

	MembersInfo   []MemberInfo `json:"membersInfo"`
	UpdateHistory []UpdateInfo `json:"updateHistory"`
}

//easyjson:json
type DocumentsInfo []DocumentInfo

//easyjson:json
type DocumentMeta struct {
	Size int64 `json:"size"`
}

//easyjson:json
type UpdateInfo struct {
	UserId    string         `json:"userId"`
	Timestamp types.DateTime `json:"timestamp"`
}

const (
	ReadPrivilege  = "read"
	WritePrivilege = "write"
	SharePrivilege = "share"
)

//easyjson:json
type MemberInfo struct {
	UserId    string `json:"userId"`
	Privilege string `json:"privilege"`
}

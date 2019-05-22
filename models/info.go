package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type DocumentInfo struct {
	ID      types.ID `json:"-"`
	Title   string   `json:"title"`
	OwnerID types.ID `json:"ownerId"`
	Subject string   `json:"subject"`
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
type UpdateInfo struct {
	UserID    types.ID       `json:"userId"`
	Timestamp types.DateTime `json:"timestamp"`
}

type Privilege string

const (
	PrivilegeRead  = "read"
	PrivilegeWrite = "write"
	PrivilegeShare = "share"
)

//easyjson:json
type MemberInfo struct {
	UserID    types.ID  `json:"userId"`
	Privilege Privilege `json:"privilege"`
}

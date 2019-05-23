package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type DocumentInfo struct {
	ID      types.ID `json:"id"`
	Title   string   `json:"title"`
	OwnerID types.ID `json:"ownerId"`
	Subject string   `json:"subject"`
}

//easyjson:json
type DocumentInfoFull struct {
	DocumentInfo
	MembersPermissions []Permission `json:"membersInfo"`
	UpdateHistory      []UpdateInfo `json:"updateHistory"`
}

//easyjson:json
type DocumentsInfo []DocumentInfo

//easyjson:json
type UpdateInfo struct {
	UserID    types.ID       `json:"userId"`
	Timestamp types.DateTime `json:"timestamp"`
}

type Scope string

const (
	ScopeRead  = "read"
	ScopeWrite = "write"
)

//easyjson:json
type Permission struct {
	UserID types.ID `json:"userId"`
	Scope  Scope    `json:"scope"`
}

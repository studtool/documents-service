package models

import (
	"github.com/studtool/common/types"
)

type Scope string

const (
	ScopeRead  = "read"
	ScopeWrite = "write"
	ScopeShare = "share"
)

//easyjson:json
type Permission struct {
	UserID types.ID `json:"userId"`
	Scope  Scope    `json:"scope"`
}

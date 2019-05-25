package models

//go:generate msgp
//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type User struct {
	ID types.ID `msg:"id" json:"id"`
}

package models

//go:generate easyjson

import (
	"github.com/studtool/common/types"
)

//easyjson:json
type DocumentInfo struct {
	Id              string         `json:"id"`
	Title           string         `json:"title"`
	OwnerId         string         `json:"ownerId"`
	Subject         string         `json:"subject"`
	DocumentSize    int64          `json:"documentSize"`
	UpdateTimestamp types.DateTime `json:"updateTimestamp"`
	UpdateUserId    string         `json:"updateUserId"`
}

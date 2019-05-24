package models

//go:generate msgp

import (
	"github.com/studtool/common/types"
)

type LongInt int64
type Content byte

const (
	ContentText  = 1
	ContentImage = 2
)

type DocumentBlock struct {
	DocumentID types.ID `msg:"-"`
	Position   LongInt  `msg:"position"`
	Type       Content  `msg:"type"`
	Data       []byte   `msg:"data"`
}

type DocumentBlocks []DocumentBlock

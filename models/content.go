package models

//go:generate msgp

type ID string
type LongInt int64
type Content byte

const (
	ContentText  = 1
	ContentImage = 2
)

type DocumentBlock struct {
	DocumentID ID      `msg:"-"`
	Position   LongInt `msg:"position"`
	Type       Content `msg:"type"`
	Data       []byte  `msg:"data"`
}

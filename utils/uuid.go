package srvutils

import (
	"github.com/studtool/common/types"
)

func MakeID() types.ID {
	id, err := types.MakeID()
	if err != nil {
		panic(err)
	}
	return id
}

package utils

import (
	"github.com/studtool/common/types"
	"github.com/studtool/common/utils"
)

func MakeID() types.ID {
	id, err := utils.MakeID()
	if err != nil {
		panic(err)
	}
	return id
}

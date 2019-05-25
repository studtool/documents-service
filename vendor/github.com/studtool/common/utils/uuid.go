package utils

import (
	"github.com/google/uuid"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"
)

func MakeID() (types.ID, *errs.Error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return consts.EmptyString, errs.New(err)
	}
	return types.ID(id.String()), nil
}

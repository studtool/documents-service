package repositories

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type UsersRepository interface {
	AddUser(u *models.User) *errs.Error
	CheckUserExistsByID(userID types.ID) *errs.Error
	DeleteUserByID(userID types.ID) *errs.Error
}

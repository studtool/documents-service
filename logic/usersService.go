package logic

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
)

type UsersService interface {
	AddUser(u *models.User) *errs.Error
	CheckUserExists(u *models.User) *errs.Error
	DeleteUser(u *models.User) *errs.Error
}

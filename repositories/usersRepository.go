package repositories

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
)

type UsersRepository interface {
	AddUser(u *models.User) *errs.Error
	CheckExistsUserByID(id string) *errs.Error
	DeleteUserByID(id string) *errs.Error
}

package rfake

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) AddUser(u *models.User) *errs.Error {
	return errs.NewNotImplementedError(r.AddUser)
}

func (r *UsersRepository) CheckExistsUserByID(id string) *errs.Error {
	return errs.NewNotImplementedError(r.CheckExistsUserByID)
}

func (r *UsersRepository) DeleteUserByID(id string) *errs.Error {
	return errs.NewNotImplementedError(r.DeleteUserByID)
}

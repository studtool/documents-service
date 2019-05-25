package rfake

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) AddUser(u *models.User) *errs.Error {
	return errs.NewNotImplementedError(r.AddUser)
}

func (r *UsersRepository) CheckUserExistsByID(userID types.ID) *errs.Error {
	return errs.NewNotImplementedError(r.CheckUserExistsByID)
}

func (r *UsersRepository) DeleteUserByID(userID string) *errs.Error {
	return errs.NewNotImplementedError(r.DeleteUserByID)
}

package sfake

import (
	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
)

type UsersService struct{}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) AddUser(u *models.User) *errs.Error {
	return errs.NewNotImplementedError(s.AddUser)
}

func (s *UsersService) CheckUserExists(u *models.User) *errs.Error {
	return errs.NewNotImplementedError(s.CheckUserExists)
}

func (s *UsersService) DeleteUser(u *models.User) *errs.Error {
	return errs.NewNotImplementedError(s.DeleteUser)
}

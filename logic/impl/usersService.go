package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/errs"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
)

type UsersService struct {
	usersRepository repositories.UsersRepository
}

type UsersServiceParams struct {
	dig.In
	UsersRepository repositories.UsersRepository
}

func NewUsersService(params UsersServiceParams) *UsersService {
	return &UsersService{
		usersRepository: params.UsersRepository,
	}
}

func (s *UsersService) AddUser(u *models.User) *errs.Error {
	return s.usersRepository.AddUser(u)
}

func (s *UsersService) CheckUserExists(u *models.User) *errs.Error {
	return s.usersRepository.CheckExistsUserByID(u.ID)
}

func (s *UsersService) DeleteUser(u *models.User) *errs.Error {
	return s.usersRepository.DeleteUserByID(u.ID)
}

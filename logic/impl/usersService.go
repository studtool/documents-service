package impl

import (
	"go.uber.org/dig"

	"github.com/studtool/common/errs"
	"github.com/studtool/common/logs"

	"github.com/studtool/documents-service/models"
	"github.com/studtool/documents-service/repositories"
	"github.com/studtool/documents-service/utils"
)

type UsersService struct {
	structLogger  logs.Logger
	reflectLogger logs.Logger

	usersRepository repositories.UsersRepository
}

type UsersServiceParams struct {
	dig.In

	UsersRepository repositories.UsersRepository

	LogsExporter *logs.Exporter
}

func NewUsersService(params UsersServiceParams) *UsersService {
	s := &UsersService{
		usersRepository: params.UsersRepository,
	}

	p := srvutils.LoggerParams{
		Value:    s,
		Exporter: params.LogsExporter,
	}

	s.structLogger = srvutils.MakeStructLogger(p)
	s.reflectLogger = srvutils.MakeReflectLogger(p)

	s.structLogger.Info("initialized")

	return s
}

func (s *UsersService) AddUser(u *models.User) *errs.Error {
	err := s.usersRepository.AddUser(u)
	if err == nil {
		s.structLogger.Infof("user [id = '%s'] added", u.ID)
	} else {
		s.structLogger.Errorf("user [id = '%s'] not added", u.ID)
	}
	return err
}

func (s *UsersService) CheckUserExists(u *models.User) *errs.Error {
	return s.usersRepository.CheckUserExistsByID(u.ID)
}

func (s *UsersService) DeleteUser(u *models.User) *errs.Error {
	err := s.usersRepository.DeleteUserByID(u.ID)
	if err == nil {
		s.structLogger.Infof("user [id = '%s'] deleted", u.ID)
	} else {
		s.structLogger.Errorf("user [id = '%s'] not deleted", u.ID)
	}
	return s.usersRepository.DeleteUserByID(u.ID)
}

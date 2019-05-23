package repositories

import (
	"github.com/studtool/common/errs"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

type PermissionsRepository interface {
	AddPermission(documentID types.ID, p *models.Permission) *errs.Error
	CheckPermission(documentID types.ID, p *models.Permission) *errs.Error
	UpdatePermission(documentID types.ID, p *models.Permission) *errs.Error
	DeletePermission(documentID types.ID, p *models.Permission) *errs.Error
}

package repository

import (
	"github.com/qomaindo-dev/go-bookstore/app/models"
	"github.com/qomaindo-dev/go-bookstore/graph/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(roleInput *model.RoleInput) (*models.Role, error)
	UpdateRole(roleInput *model.RoleInput, id int) error
	DeleteRole(id int) error
	GetOneRole(id int) (*models.Role, error)
	GetAllRoles() ([]*model.Role, error)
}

type RoleService struct {
	Db *gorm.DB
}

var _ RoleRepository = &RoleService{}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		Db: db,
	}
}

func (rs *RoleService) CreateRole(roleInput *model.RoleInput) (*models.Role, error) {
	roleData := &models.Role{
		Name:        roleInput.Name,
		Description: roleInput.Description,
		IsDeleted:   roleInput.IsDeleted,
	}

	err := rs.Db.Create(&roleData).Error

	return roleData, err
}

func (rs *RoleService) UpdateRole(roleInput *model.RoleInput, id int) error {
	role := &models.Role{
		ID:          id,
		Name:        roleInput.Name,
		Description: roleInput.Description,
		IsDeleted:   roleInput.IsDeleted,
	}
	err := rs.Db.Model(&role).Where("id = ?", id).Updates(role).Error

	return err
}

func (rs *RoleService) DeleteRole(id int) error {
	role := &models.Role{}
	err := rs.Db.Where(role, id).Error

	return err
}

func (rs *RoleService) GetOneRole(id int) (*models.Role, error) {
	role := &models.Role{}
	err := rs.Db.Where("id = ?", id).First(role).Error

	return role, err
}

func (rs *RoleService) GetAllRoles() ([]*model.Role, error) {
	roles := []*model.Role{}

	err := rs.Db.Find(&roles).Error

	return roles, err
}

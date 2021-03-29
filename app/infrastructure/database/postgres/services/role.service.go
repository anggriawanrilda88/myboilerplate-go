package services

import (
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"gorm.io/gorm"
)

// RoleService interface
type RoleService interface {
	Transaction() (response *gorm.DB)
	FindOne(Body *models.Role, RoleID uint) (response *gorm.DB)
}

// NewRoleService Instantiate Model of Role
func NewRoleService() RoleService {
	return &roleService{}
}

type roleService struct {
}

// Transaction function
func (fn *roleService) Transaction() (tx *gorm.DB) {
	tx = database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return tx
}

// FindOne function to get all role list
func (fn *roleService) FindOne(Role *models.Role, RoleID uint) (response *gorm.DB) {
	response = database.DB.Find(&Role, RoleID)
	return response
}

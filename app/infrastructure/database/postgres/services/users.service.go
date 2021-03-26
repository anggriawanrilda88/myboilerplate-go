package services

import (
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"gorm.io/gorm"
)

// UsersService interface
type UsersService interface {
	Transaction() (response *gorm.DB)
	Create(Body *models.User) (response *gorm.DB)

	// // create with transaction
	// Create(tx *gorm.DB, Body *models.User) (response *gorm.DB)
}

// NewUsersService Instantiate Model of Users
func NewUsersService() UsersService {
	return &usersService{}
}

type usersService struct {
}

// Transaction function to get all users list with transaction
func (fn *usersService) Transaction() (tx *gorm.DB) {
	tx = database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return tx
}

// Create function to get all users list
func (fn *usersService) Create(Body *models.User) (response *gorm.DB) {
	response = database.DB.Create(&Body)
	return response
}

// // Create function to get all users list with transaction
// func (fn *usersService) Create(tx *gorm.DB, Body *models.User) (response *gorm.DB) {
// 	response = tx.Create(&Body)
// 	return response
// }

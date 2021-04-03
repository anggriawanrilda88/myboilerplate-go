package services

import (
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UsersService interface
type UsersService interface {
	Transaction() (response *gorm.DB)
	Create(Body *models.User) (response *gorm.DB)
	FindOne(Users *models.User, data interface{}) (response *gorm.DB)
	Find(ctx *fiber.Ctx, Users []models.User) (data []models.User, err error)
	GetVersionCount() (count uint, err error)

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

// FindOne function to get users list
func (fn *usersService) FindOne(Users *models.User, data interface{}) (response *gorm.DB) {
	response = database.DB.Find(&Users, data)
	return response
}

// FindOne function to get users list
func (fn *usersService) GetVersionCount() (count uint, err error) {
	response := database.DB.Raw("select sum(version) as version from users").Scan(&count)
	if response.Error != nil {
		return
	}

	return
}

// FindOne function to get all users list
func (fn *usersService) Find(ctx *fiber.Ctx, Users []models.User) (data []models.User, err error) {
	response := database.DB.Raw(`select  
					A.id,
					A.created_at,
					A.updated_at,
					A.deleted_at,
					A.name,
					A.password,
					A.email,
					A.role_id,
					A.version,
					row_to_json(B.*)::jsonb as "role"
				from users as A
				join roles as B on A.role_id = B.id 
				group by A.id, B.id`).Scan(&Users)
	if response.Error != nil {
		err = response.Error
		return
	}

	return Users, nil
}

// // Create function to get all users list with transaction
// func (fn *usersService) Create(tx *gorm.DB, Body *models.User) (response *gorm.DB) {
// 	response = tx.Create(&Body)
// 	return response
// }

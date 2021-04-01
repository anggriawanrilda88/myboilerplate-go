package services

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UsersService interface
type UsersService interface {
	Transaction() (response *gorm.DB)
	Create(Body *models.User) (response *gorm.DB)
	FindOne(Users *models.User, data interface{}) (response *gorm.DB)
	Find(ctx *fiber.Ctx, Users []models.User) (data []models.User, err error)

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

// FindOne function to get all users list
func (fn *usersService) Find(ctx *fiber.Ctx, Users []models.User) (data []models.User, err error) {
	var version int
	database.DB.Raw("select sum(version) as version from users").Scan(&version)
	strVersion := strconv.Itoa(version)
	cache := helper.GetCache(ctx, strVersion)
	if cache.Err() != nil {
		if cache.Err().Error() == "redis: nil" {
			response := database.DB.Find(&Users)
			if response.Error != nil {
				err = response.Error
				return
			}
			err = helper.SetCache(ctx, strVersion, Users)
			if err != nil {
				return
			}

			data = Users
			log.Println("dari postgres loo")
			return
		}
		return
	}

	byte, err := cache.Bytes()
	err = json.Unmarshal(byte, &Users)
	data = Users
	log.Println("dari redis loo")

	return
}

// // Create function to get all users list with transaction
// func (fn *usersService) Create(tx *gorm.DB, Body *models.User) (response *gorm.DB) {
// 	response = tx.Create(&Body)
// 	return response
// }

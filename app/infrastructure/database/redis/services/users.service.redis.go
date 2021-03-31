package services

import (
	"encoding/json"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// UsersServiceRedis interface
type UsersServiceRedis interface {
	Create(ctx *fiber.Ctx, Body *models.User) (response *redis.StatusCmd)
	FindOne(ctx *fiber.Ctx) (response *redis.StringCmd)

	// // create with transaction
	// Create(tx *gorm.DB, Body *models.User) (response *gorm.DB)
}

// NewUsersServiceRedis Instantiate Model of Users
func NewUsersServiceRedis() UsersServiceRedis {
	return &usersServiceRedis{}
}

type usersServiceRedis struct {
}

// Create function to get all users list
func (fn *usersServiceRedis) Create(ctx *fiber.Ctx, Body *models.User) (response *redis.StatusCmd) {
	haha, _ := json.Marshal(fiber.Map{"1234.users": Body})
	response = database.Redis.Set(ctx.Context(), "keys", haha, 0)
	return
}

// FindOne function to get all users list
func (fn *usersServiceRedis) FindOne(ctx *fiber.Ctx) (response *redis.StringCmd) {
	response = database.Redis.Get(ctx.Context(), "keys")
	return
}

// // Create function to get all users list with transaction
// func (fn *usersServiceRedis) Create(tx *gorm.DB, Body *models.User) (response *gorm.DB) {
// 	response = tx.Create(&Body)
// 	return response
// }

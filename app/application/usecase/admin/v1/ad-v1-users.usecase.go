package usecases

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/services"
	redisService "github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/redis/services"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/helper"
	"github.com/gofiber/fiber/v2"
)

// UsersUseCase interface
type UsersUseCase interface {
	Create(ctx *fiber.Ctx, Body *models.User) (err error)
	Find(ctx *fiber.Ctx, Users []models.User) (data interface{}, count uint, err error)
	FindOne(ctx *fiber.Ctx, Users *models.User) (data interface{}, err error)
}

// NewUsersUseCase Instantiate the UseCase
func NewUsersUseCase() UsersUseCase {
	return &usersUseCase{
		service:           services.NewUsersService(),
		serviceRole:       services.NewRoleService(),
		serviceRedisUsers: redisService.NewUsersServiceRedis(),
	}
}

type usersUseCase struct {
	service           services.UsersService
	serviceRole       services.RoleService
	serviceRedisUsers redisService.UsersServiceRedis
}

// Create usecase Users
func (fn *usersUseCase) Create(ctx *fiber.Ctx, Body *models.User) (err error) {
	Body.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	Body.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	Body.Version = 1
	if response := fn.service.Create(Body); response.Error != nil {
		return response.Error
	}

	// Match role to user
	if Body.RoleID != 0 {
		Role := new(models.Role)
		res := fn.serviceRole.FindOne(Role, Body.RoleID)
		if res.Error != nil {
			return res.Error
		}

		Body.Role, err = json.Marshal(Role)
		if err != nil {
			return err
		}
	}

	return err
}

// Find usecase Users
func (fn *usersUseCase) Find(ctx *fiber.Ctx, Users []models.User) (data interface{}, count uint, err error) {
	version, count, err := fn.service.GetVersionCount()
	if err != nil {
		return
	}

	// get list with redis cache
	strVersion := strconv.Itoa(int(version))
	cache := helper.GetCache(ctx, strVersion)
	if cache.Err() != nil {
		if cache.Err().Error() == "redis: nil" {
			responseUsers, err2 := fn.service.Find(ctx, Users)
			if err2 != nil {
				return nil, 0, err2
			}

			err2 = helper.SetCache(ctx, strVersion, responseUsers)
			if err2 != nil {
				return nil, 0, err2
			}

			log.Println("dari postgres loo")
			return responseUsers, count, nil
		}
		return
	}

	byte, err := cache.Bytes()
	err = jsoniter.Unmarshal(byte, &Users)
	if err != nil {
		return nil, 0, err
	}

	log.Println("dari redis loo")
	return Users, count, nil
}

// Find usecase Users
func (fn *usersUseCase) FindOne(ctx *fiber.Ctx, Users *models.User) (data interface{}, err error) {
	id := ctx.Params("id")
	if response := fn.service.FindOne(Users, id); response.Error != nil {
		err = response.Error
		return
	}

	// Match role to user
	if Users.RoleID != 0 {
		Role := new(models.Role)
		res := fn.serviceRole.FindOne(Role, Users.RoleID)
		if res.Error != nil {
			return nil, res.Error
		}

		Users.Role, err = json.Marshal(Role)
		if err != nil {
			return nil, err
		}
	}

	return Users, nil
}

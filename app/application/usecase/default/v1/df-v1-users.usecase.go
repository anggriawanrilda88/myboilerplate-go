package usecases

import (
	"time"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/services"
	"github.com/gofiber/fiber/v2"
)

// UsersUseCase interface
type UsersUseCase interface {
	Create(ctx *fiber.Ctx, Body *models.User) (err error)

	// // Create usecase Users with transaction
	// Create(ctx *fiber.Ctx, Body *models.User) (err error)
}

// NewUsersUseCase Instantiate the UseCase
func NewUsersUseCase() UsersUseCase {
	return &usersUseCase{
		service:     services.NewUsersService(),
		serviceRole: services.NewRoleService(),
	}
}

type usersUseCase struct {
	service     services.UsersService
	serviceRole services.RoleService
}

// Create usecase Users
func (fn *usersUseCase) Create(ctx *fiber.Ctx, Body *models.User) (err error) {
	Body.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	Body.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if response := fn.service.Create(Body); response.Error != nil {
		err = response.Error
	}

	// Match role to user
	if Body.RoleID != "" {
		Role := new(models.Role)
		_ = fn.serviceRole.FindOne(Role, Body.RoleID)
		Body.Role = *Role
	}

	err = ctx.JSON(Body)
	return err
}

// // Create usecase Users with transaction
// func (fn *usersUseCase) Create(ctx *fiber.Ctx, Body *models.User) (err error) {
// 	// begin transaction
// 	tx := fn.service.Transaction(DB)
// 	if err := tx.Error; err != nil {
// 		return err
// 	}

// 	Body.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
// 	Body.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
// 	if err = fn.service.Create(tx, Body).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// Match role to user
// 	if Body.RoleID != "" {
// 		Role := new(models.Role)
// 		err = tx.Find(&Role, Body.RoleID).Error
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 		Body.Role = *Role
// 	}

// 	// create json from response db
// 	err = ctx.JSON(Body)

// 	// commit if not error
// 	if err == nil {
// 		tx.Commit()
// 	}
// 	return err
// }

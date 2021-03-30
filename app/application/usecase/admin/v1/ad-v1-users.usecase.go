package usecases

import (
	"time"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/services"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/helper"
	configuration "github.com/anggriawanrilda88/myboilerplate/config"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// UsersUseCase interface
type UsersUseCase interface {
	Create(ctx *fiber.Ctx, Body *models.User) (err error)
	LoginUser(ctx *fiber.Ctx, User *models.User, UserLogin *models.UserLogin) (err error)
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
func (fn *usersUseCase) LoginUser(ctx *fiber.Ctx, User *models.User, UserLogin *models.UserLogin) (err error) {
	if response := fn.service.FindOne(User, UserLogin); response.Error != nil {
		return helper.ErrorHandler(ctx, fiber.ErrForbidden, 400, response.Error.Error())
	}

	// set error when user login not found
	if User.Id == 0 {
		return helper.ErrorHandler(ctx, fiber.ErrForbidden, 400, "User Not Found")
	}

	// Match role to user
	if User.RoleID != 0 {
		Role := new(models.Role)
		_ = fn.serviceRole.FindOne(Role, User.RoleID)
		User.Role = *Role
	}

	config := configuration.New().Viper
	secretKey := config.GetString("JWT_SECRET")

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, fiber.ErrInternalServerError.Message)
	}

	// set token to struct
	UserLogin.Token = t

	return err
}

// Create usecase Users
func (fn *usersUseCase) Create(ctx *fiber.Ctx, Body *models.User) (err error) {
	Body.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	Body.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if response := fn.service.Create(Body); response.Error != nil {
		return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, response.Error.Error())
	}

	// Match role to user
	if Body.RoleID != 0 {
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
// 	tx := fn.service.Transaction()
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

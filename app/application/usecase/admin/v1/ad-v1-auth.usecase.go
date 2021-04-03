package usecases

import (
	"errors"
	"time"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/services"
	redisService "github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/redis/services"
	configuration "github.com/anggriawanrilda88/myboilerplate/config"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthUseCase interface
type AuthUseCase interface {
	Login(ctx *fiber.Ctx, User *models.User, UserLogin *models.UserLogin) (err error)
}

// NewAuthUseCase Instantiate the UseCase
func NewAuthUseCase() AuthUseCase {
	return &authUseCase{
		service:           services.NewUsersService(),
		serviceRole:       services.NewRoleService(),
		serviceRedisUsers: redisService.NewUsersServiceRedis(),
	}
}

type authUseCase struct {
	service           services.UsersService
	serviceRole       services.RoleService
	serviceRedisUsers redisService.UsersServiceRedis
}

// Create usecase Users
func (fn *authUseCase) Login(ctx *fiber.Ctx, User *models.User, UserLogin *models.UserLogin) (err error) {
	// get user info from service
	if response := fn.service.FindOne(User, UserLogin); response.Error != nil {
		return response.Error
	}

	// set error when user login not found
	if User.Id == 0 {
		err = errors.New("User Not Found")
		return
	}

	// Match role to user
	if User.RoleID != 0 {
		Role := new(models.Role)
		_ = fn.serviceRole.FindOne(Role, User.RoleID)
		// User.Role = *Role
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
		return err
	}

	// set token to struct
	UserLogin.Token = t

	return err
}

package auth

import (
	usecases "github.com/anggriawanrilda88/myboilerplate/app/application/usecase/admin/v1"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/helper"
	"github.com/go-playground/validator"

	"github.com/gofiber/fiber/v2"
)

// AuthController interface
type AuthController interface {
	Login(api fiber.Router) fiber.Handler
}

// NewAuthController Instantiate the Controller
func NewAuthController() AuthController {
	return &authController{
		authUsecase:   usecases.NewAuthUseCase(),
		authTransform: NewAuthTransform(),
	}
}

type authController struct {
	authUsecase   usecases.AuthUseCase
	authTransform AuthTransform
}

var validate *validator.Validate

// Create a single user to the database
func (fn *authController) Login(api fiber.Router) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// set parameter
		UserLogin := new(models.UserLogin)
		User := new(models.User)

		if err := ctx.BodyParser(UserLogin); err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrForbidden, 400, "Cannot unmarshal request body, wrong type data.")
		}

		// validate request body
		if err := validator.New().Struct(UserLogin); err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrForbidden, 400, err.Error())
		}

		// get usecase auth
		if err := fn.authUsecase.Login(ctx, User, UserLogin); err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrForbidden, 400, err.Error())
		}

		transform := fn.authTransform.DetailTransform(ctx, fiber.Map{"user": User, "token": UserLogin.Token})
		return ctx.JSON(transform)
	}
}

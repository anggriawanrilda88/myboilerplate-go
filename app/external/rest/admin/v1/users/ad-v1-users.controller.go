package users

import (
	usecases "github.com/anggriawanrilda88/myboilerplate/app/application/usecase/admin/v1"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/helper"
	"github.com/go-playground/validator"

	"github.com/gofiber/fiber/v2"
)

// UsersController interface
type UsersController interface {
	Create(api fiber.Router) fiber.Handler
	Find(api fiber.Router) fiber.Handler
	FindOne(api fiber.Router) fiber.Handler
}

// NewUsersController Instantiate the Controller
func NewUsersController() UsersController {
	return &usersController{
		usersUsecase:  usecases.NewUsersUseCase(),
		userTransform: NewUsersTransform(),
	}
}

type usersController struct {
	usersUsecase  usecases.UsersUseCase
	userTransform UsersTransform
}

var validate *validator.Validate

// Create a single user to the database
func (fn *usersController) Create(api fiber.Router) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//set body to struct
		Body := new(models.User)
		err := ctx.BodyParser(Body)
		if err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, "Cannot unmarshal request body, wrong type data.")
		}

		//validate request body
		err = validator.New().Struct(Body)
		if err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, err.Error())
		}

		//get usecase users
		err = fn.usersUsecase.Create(ctx, Body)
		if err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, err.Error())
		}

		transform := fn.userTransform.DetailTransform(Body, 10, 0, 1)
		return ctx.JSON(transform)
	}
}

// Create a single user to the database
func (fn *usersController) Find(api fiber.Router) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//get usecase users
		var Users []models.User
		response, count, err := fn.usersUsecase.Find(ctx, Users)
		if err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, err.Error())
		}

		transform := fn.userTransform.DetailTransform(response, 10, 0, count)
		return ctx.JSON(transform)
	}
}

// Create a single user to the database
func (fn *usersController) FindOne(api fiber.Router) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//get usecase users
		Users := new(models.User)
		response, err := fn.usersUsecase.FindOne(ctx, Users)
		if err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, err.Error())
		}

		transform := fn.userTransform.DetailTransform(response, 10, 0, 1)
		return ctx.JSON(transform)
	}
}

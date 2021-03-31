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
	// GetAllUsers(DB *database.Database) fiber.Handler
	// GetUser(DB *database.Database) fiber.Handler
	Create(api fiber.Router) fiber.Handler
	LoginUser(api fiber.Router) fiber.Handler
	Find(api fiber.Router) fiber.Handler
	// EditUser(DB *database.Database) fiber.Handler
	// DeleteUser(DB *database.Database) fiber.Handler
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
func (fn *usersController) Find(api fiber.Router) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//get usecase users
		var Users []models.User
		response, err := fn.usersUsecase.Find(ctx, Users)
		if err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrInternalServerError, 400, err.Error())
		}

		err = ctx.JSON(response)

		return err
	}
}

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

		return ctx.JSON(Body)
	}
}

// Create a single user to the database
func (fn *usersController) LoginUser(api fiber.Router) fiber.Handler {
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

		// get usecase users
		if err := fn.usersUsecase.LoginUser(ctx, User, UserLogin); err != nil {
			return helper.ErrorHandler(ctx, fiber.ErrForbidden, 400, err.Error())
		}

		transform := fn.userTransform.DetailTransform(ctx, fiber.Map{"user": User, "token": UserLogin.Token})
		return ctx.JSON(transform)
	}
}

// // GetAllUsers Return all users as JSON
// func (_c *usersController) GetAllUsers(DB *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		var Users []models.User
// 		if response := DB.Find(&Users); response.Error != nil {
// 			panic("Error occurred while retrieving users from the database: " + response.Error.Error())
// 		}
// 		// Match roles to users

// 		for index, User := range Users {
// 			if User.RoleID != 0 {
// 				Role := new(models.Role)
// 				if response := DB.Find(&Role, User.RoleID); response.Error != nil {
// 					panic("An error occurred when retrieving the role: " + response.Error.Error())
// 				}
// 				if Role.ID != 0 {
// 					Users[index].Role = *Role
// 				}
// 			}
// 		}
// 		err := ctx.JSON(Users)
// 		if err != nil {
// 			panic("Error occurred when returning JSON of users: " + err.Error())
// 		}
// 		return err
// 	}
// }

// // Return a single user as JSON
// func (_c *usersController) GetUser(DB *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		User := new(models.User)
// 		id := ctx.Params("id")
// 		if response := DB.Find(&User, id); response.Error != nil {
// 			panic("An error occurred when retrieving the user: " + response.Error.Error())
// 		}
// 		if User.ID == 0 {
// 			err := ctx.SendStatus(fiber.StatusNotFound)
// 			if err != nil {
// 				panic("Cannot return status not found: " + err.Error())
// 			}
// 			err = ctx.JSON(fiber.Map{
// 				"ID": id,
// 			})
// 			if err != nil {
// 				panic("Error occurred when returning JSON of a role: " + err.Error())
// 			}
// 			return err
// 		}
// 		// Match role to user
// 		if User.RoleID != 0 {
// 			Role := new(models.Role)
// 			if response := DB.Find(&Role, User.RoleID); response.Error != nil {
// 				panic("An error occurred when retrieving the role: " + response.Error.Error())
// 			}
// 			if Role.ID != 0 {
// 				User.Role = *Role
// 			}
// 		}
// 		err := ctx.JSON(User)
// 		if err != nil {
// 			panic("Error occurred when returning JSON of a user: " + err.Error())
// 		}
// 		return err
// 	}
// }

// // Edit a single user
// func (_c *usersController) EditUser(DB *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		id := ctx.Params("id")
// 		EditUser := new(models.User)
// 		User := new(models.User)
// 		if err := ctx.BodyParser(EditUser); err != nil {
// 			panic("An error occurred when parsing the edited user: " + err.Error())
// 		}
// 		if response := DB.Find(&User, id); response.Error != nil {
// 			panic("An error occurred when retrieving the existing user: " + response.Error.Error())
// 		}
// 		// User does not exist
// 		if User.ID == 0 {
// 			err := ctx.SendStatus(fiber.StatusNotFound)
// 			if err != nil {
// 				panic("Cannot return status not found: " + err.Error())
// 			}
// 			err = ctx.JSON(fiber.Map{
// 				"ID": id,
// 			})
// 			if err != nil {
// 				panic("Error occurred when returning JSON of a user: " + err.Error())
// 			}
// 			return err
// 		}
// 		User.Name = EditUser.Name
// 		User.Email = EditUser.Email
// 		User.RoleID = EditUser.RoleID
// 		// Match role to user
// 		if User.RoleID != 0 {
// 			Role := new(models.Role)
// 			if response := DB.Find(&Role, User.RoleID); response.Error != nil {
// 				panic("An error occurred when retrieving the role" + response.Error.Error())
// 			}
// 			if Role.ID != 0 {
// 				User.Role = *Role
// 			}
// 		}
// 		// Save user
// 		DB.Save(&User)

// 		err := ctx.JSON(User)
// 		if err != nil {
// 			panic("Error occurred when returning JSON of a user: " + err.Error())
// 		}
// 		return err
// 	}
// }

// // Delete a single user
// func (_c *usersController) DeleteUser(DB *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		id := ctx.Params("id")
// 		var User models.User
// 		DB.Find(&User, id)
// 		if response := DB.Find(&User); response.Error != nil {
// 			panic("An error occurred when finding the user to be deleted" + response.Error.Error())
// 		}
// 		DB.Delete(&User)

// 		err := ctx.JSON(fiber.Map{
// 			"ID":      id,
// 			"Deleted": true,
// 		})
// 		if err != nil {
// 			panic("Error occurred when returning JSON of a user: " + err.Error())
// 		}
// 		return err
// 	}
// }

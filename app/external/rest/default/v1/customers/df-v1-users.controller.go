package customers

import (
	"errors"
	"log"

	usecases "github.com/anggriawanrilda88/myboilerplate/app/application/usecase/default/v1"
	"github.com/gofiber/fiber/v2"
)

// UsersController interface
type UsersController interface {
	FindAll() fiber.Handler
}

// NewUsersController Instantiate the Controller
func NewUsersController() UsersController {
	return &usersController{
		useCase: usecases.NewUsersUseCase(),
	}
}

type usersController struct {
	useCase usecases.UsersUseCase
}

// FindAll interface
func (_c *usersController) FindAll() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		_ = _c.useCase.FindAll(ctx)
		log.Println("Aku Controller")
		users := fiber.Map{
			"success": true,
			"user":    1,
		}
		err := ctx.JSON(users)
		err = errors.New("hahahahah")
		// if err != nil {
		// 	err = ctx.JSON(fiber.Map{
		// 		"status":  500,
		// 		"error":   "Internal Server Error",
		// 		"message": "No message available",
		// 		"code":    "/api/book/1",
		// 	})
		// 	// panic("Error occurred when returning JSON of users: " + err.Error())
		// }
		return err
	}
}

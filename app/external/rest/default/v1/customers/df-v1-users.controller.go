package customers

import (
	"log"

	usecases "github.com/anggriawanrilda88/myboilerplate/app/application/usecase/default/v1"
	"github.com/gofiber/fiber/v2"
)

// UsersController interface
type UsersController interface {
	FindAll(c *fiber.Ctx) error
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
func (_c *usersController) FindAll(c *fiber.Ctx) (err error) {
	err = _c.useCase.FindAll(c)
	log.Println("Aku Controller")
	return c.JSON(fiber.Map{
		"success": true,
		"user":    1,
	})
}

package usecases

import (
	"log"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/services"
	"github.com/gofiber/fiber/v2"
)

// UsersUseCase interface
type UsersUseCase interface {
	FindAll(c *fiber.Ctx) (err error)
}

// NewUsersUseCase Instantiate the UseCase
func NewUsersUseCase() UsersUseCase {
	return &usersUseCase{
		m: services.NewUsersService(),
	}
}

type usersUseCase struct {
	m services.UsersService
}

// FindAll interface
func (_c *usersUseCase) FindAll(c *fiber.Ctx) (err error) {
	err = _c.m.FindAll(c)
	log.Println("Aku Use Case")
	return nil
}

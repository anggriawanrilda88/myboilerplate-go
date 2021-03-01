package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// UsersService interface
type UsersService interface {
	FindAll(c *fiber.Ctx) (err error)
}

// NewUsersService Instantiate Model of Users
func NewUsersService() UsersService {
	return &usersService{}
}

type usersService struct{}

// FindAll function to get all users list
func (_m *usersService) FindAll(c *fiber.Ctx) (err error) {
	log.Println("Aku Service")
	return nil
}

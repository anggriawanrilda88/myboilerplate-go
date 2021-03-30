package users

import "github.com/gofiber/fiber/v2"

// UsersTransform interface
type UsersTransform interface {
	DetailTransform(ctx *fiber.Ctx, data map[string]interface{}) map[string]interface{}
}

// NewUsersTransform Instantiate the Transform
func NewUsersTransform() UsersTransform {
	return &usersTransform{}
}

type usersTransform struct {
}

func (fn *usersTransform) DetailTransform(ctx *fiber.Ctx, data map[string]interface{}) map[string]interface{} {
	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"limit": 10,
			"total": 1,
			"skip":  0,
		},
	}
}

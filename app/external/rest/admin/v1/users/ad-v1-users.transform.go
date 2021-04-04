package users

import "github.com/gofiber/fiber/v2"

// UsersTransform interface
type UsersTransform interface {
	DetailTransform(data interface{}, limit uint, skip uint, total uint) map[string]interface{}
}

// NewUsersTransform Instantiate the Transform
func NewUsersTransform() UsersTransform {
	return &usersTransform{}
}

type usersTransform struct {
}

func (fn *usersTransform) DetailTransform(data interface{}, limit uint, skip uint, total uint) map[string]interface{} {
	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"limit": limit,
			"total": total,
			"skip":  skip,
		},
	}
}

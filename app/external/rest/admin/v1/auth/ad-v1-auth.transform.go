package auth

import "github.com/gofiber/fiber/v2"

// AuthTransform interface
type AuthTransform interface {
	DetailTransform(ctx *fiber.Ctx, data map[string]interface{}) map[string]interface{}
}

// NewAuthTransform Instantiate the Transform
func NewAuthTransform() AuthTransform {
	return &authTransform{}
}

type authTransform struct {
}

func (fn *authTransform) DetailTransform(ctx *fiber.Ctx, data map[string]interface{}) map[string]interface{} {
	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"limit": 10,
			"total": 1,
			"skip":  0,
		},
	}
}

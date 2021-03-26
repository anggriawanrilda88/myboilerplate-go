package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// AuthBasic login users
func AuthBasic() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("ANGGRI===================", ctx.Get("Authorization"))

		err := ctx.Next()
		return err
	}
}

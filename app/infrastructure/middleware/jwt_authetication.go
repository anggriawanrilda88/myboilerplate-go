package middleware

import (
	configuration "github.com/anggriawanrilda88/myboilerplate/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JWTAuthentication login users
func JWTAuthentication() (err fiber.Handler) {
	config := configuration.New().Viper
	secretKey := config.GetString("JWT_SECRET")

	if config.GetBool("MW_FIBER_AUTHENTICATION_ENABLED") {
		// jwt middleware function, handle filter, SigningKey and ErrorHandler
		err = jwtware.New(jwtware.Config{
			Filter: func(ctx *fiber.Ctx) (skip bool) {
				//skiped list route and method
				var allowedRPC = []string{
					"/api/v1/users.POST",
					"/api/v1/users/login.POST",
				}

				//fungsi untuk allowed api berdasarkan list
				for _, val := range allowedRPC {
					if ctx.OriginalURL()+"."+ctx.Method() == val {
						skip = true
						break
					}
				}
				return skip
			},
			SigningKey: []byte(secretKey),
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				return err
			},
		})
	}
	return err
}

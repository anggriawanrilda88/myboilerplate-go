package module

import (
	customers "github.com/anggriawanrilda88/myboilerplate/app/external/rest/default/v1/users"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

//RegisterRoute for register all route created
func RegisterRoute(app *fiber.App, config *viper.Viper) {
	// set route group
	api := app.Group("/api")

	// interceptor for auth basic
	if config.GetBool("MW_FIBER_AUTHENTICATION_ENABLED") {
		api.Use(middleware.AuthBasic())
	}

	// register route users
	registerUsersV1(api, app)
}

func registerUsersV1(api fiber.Router, app *fiber.App) {
	route := api.Group("/v1/users")
	// route.Get("/", customers.NewUsersController().GetAllUsers(DB))
	// route.Get("/:id", customers.NewUsersController().GetUser(DB))
	route.Post("/", customers.NewUsersController().AddUser(api))
}

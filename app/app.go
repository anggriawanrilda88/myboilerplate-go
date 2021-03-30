package module

import (
	customers "github.com/anggriawanrilda88/myboilerplate/app/external/rest/admin/v1/users"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
)

//RegisterRoute for register all route created
func RegisterRoute(app *fiber.App) {
	// set route group
	api := app.Group("/api")

	// interceptor for auth basic
	api.Use(middleware.JWTAuthentication())

	// register route users
	registerUsersV1(api, app)
}

func registerUsersV1(api fiber.Router, app *fiber.App) {
	route := api.Group("/v1/users")
	// route.Get("/", customers.NewUsersController().GetAllUsers(DB))
	// route.Get("/:id", customers.NewUsersController().GetUser(DB))
	route.Post("/", customers.NewUsersController().Create(api))
	route.Post("/login", customers.NewUsersController().LoginUser(api))
}

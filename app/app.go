package module

import (
	"github.com/anggriawanrilda88/myboilerplate/app/external/rest/admin/v1/auth"
	"github.com/anggriawanrilda88/myboilerplate/app/external/rest/admin/v1/users"
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

	// register route users
	registerAuthV1(api, app)
}

func registerUsersV1(api fiber.Router, app *fiber.App) {
	route := api.Group("/v1/users")
	route.Post("/", users.NewUsersController().Create(api))
	route.Get("/", users.NewUsersController().Find(api))
	route.Get("/:id", users.NewUsersController().FindOne(api))
}

func registerAuthV1(api fiber.Router, app *fiber.App) {
	route := api.Group("/v1/auth")
	route.Post("/login", auth.NewAuthController().Login(api))
}

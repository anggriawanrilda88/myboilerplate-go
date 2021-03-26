package module

import (
	customers "github.com/anggriawanrilda88/myboilerplate/app/external/rest/default/v1/users"
	"github.com/gofiber/fiber/v2"
)

//RegisterRoute for register all route created
func RegisterRoute(api fiber.Router) {
	registerUsersV1(api)
}

func registerUsersV1(api fiber.Router) {
	route := api.Group("/v1/users")
	// route.Get("/", customers.NewUsersController().GetAllUsers(DB))
	// route.Get("/:id", customers.NewUsersController().GetUser(DB))
	route.Post("/", customers.NewUsersController().AddUser())
}

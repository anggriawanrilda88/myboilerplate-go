package module

import (
	"github.com/anggriawanrilda88/myboilerplate/app/external/rest/default/v1/customers"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/gofiber/fiber/v2"
)

//RegisterRoute for register all route created
func RegisterRoute(api fiber.Router, DB *database.Database) {
	registerUsersV1(api, DB)
}

func registerUsersV1(api fiber.Router, DB *database.Database) {
	route := api.Group("/v1/users")
	route.Get("/", customers.NewUsersController().GetAllUsers(DB))
}

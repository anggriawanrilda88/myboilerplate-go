package main

import (
	module "github.com/anggriawanrilda88/myboilerplate/app"
	"github.com/anggriawanrilda88/myboilerplate/database"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":5000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// config := &Config{
	// 	Viper: viper.New(),
	// }

	// config.SetConfigFile(".env")
	// // Find and read the config file
	// err := config.ReadInConfig()

	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }

	// value, ok := config.Get("APP_ADDR").(string)

	// if !ok {
	// 	log.Fatalf("Invalid type assertion")
	// }

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api endpoint from app module
	api := app.Group("/api")
	module.RegisterRoute(api)

	// // Bind handlers
	// v1.Get("/users", handlers.UserList)
	// v1.Post("/users", handlers.UserCreate)

	//

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}

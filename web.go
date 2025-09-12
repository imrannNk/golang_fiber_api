package main

import (
	"go-fiber-api/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${time}] - ${latency} - ${bytesSent} ${ip} - ${status} - ${method} ${path}\n",
	}))

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:userId", handlers.GetUser)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:userId", handlers.UpdateUser)
	app.Delete("/users/:userId", handlers.DeleteUser)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	log.Fatal(app.Listen(":3000"))
}

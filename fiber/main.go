package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type fake struct {
	id   string
	data string
}

func main() {
	app := fiber.New()
	data := fake{id: "9012", data: "Joker"}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(data.data)
	})

	log.Fatal(app.Listen(":3000"))
}

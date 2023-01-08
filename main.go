package main

import "github.com/gofiber/fiber/v2"

type Ninja struct {
	Name   string
	Weapon string
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("halo riyan")
	})
	app.Listen(":8080")
}

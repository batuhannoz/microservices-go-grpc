package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/starwars", func(ctx *fiber.Ctx) error {
		// return random starwars character
		return nil
	})
	app.Get("/harrypotter", func(ctx *fiber.Ctx) error {
		// return random harrypotter character
		return nil
	})
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

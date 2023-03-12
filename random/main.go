package main

import (
	"github.com/batuhannoz/microservices-go-grpc/random/client"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	StarWarsClient := client.NewStarWarsClient()
	app := fiber.New()
	app.Get("/starwars", func(ctx *fiber.Ctx) error {
		character := StarWarsClient.GetRandomCharacter()
		return ctx.Status(http.StatusOK).JSON(character)
	})
	app.Get("/harrypotter", func(ctx *fiber.Ctx) error {
		// return random harrypotter character
		return nil
	})
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

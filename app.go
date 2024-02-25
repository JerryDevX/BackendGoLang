package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	println("Hello, World!")

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("TEST") // => Hello john 👋!
    })

    log.Fatal(app.Listen(":3000"))


	app.Listen(":3000")

}
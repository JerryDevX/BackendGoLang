package main

import (
	"backend/configs"
	"backend/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	println("Hello, World!")

	config := configs.FiberConfig()


	app := fiber.New(config)
	

	if os.Getenv("STATUS") == "prod" {
		utils.Running(app)
	} else {
		utils.Server(app)
	}

}

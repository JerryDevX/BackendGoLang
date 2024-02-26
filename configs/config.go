package configs

import (
	"os"

	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("FIBER_READ_TIMEOUT_SECONDS"))

	return fiber.Config{
		ReadTimeout: time.Duration(readTimeoutSecondsCount) * time.Second,
	}
}
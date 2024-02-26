package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
)


func Server(a *fiber.App) {
	// Create a channel to receive OS signals
    sigs := make(chan os.Signal, 1)

    // Register the channel to receive specific signals
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Start a goroutine that will do something when a signal is received
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        os.Exit(0)
    }()

	fiberConfig, _ := ConfigType("fiber")

	if err := a.Listen(fiberConfig); err != nil {
		log.Printf("Oops... Server is not running: %v", err)
	}

	<-sigs

}

func Running(a *fiber.App) {
	ConnURL, _ := ConfigType("fiber")

	if err := a.Listen(ConnURL); err != nil {
		log.Printf("Oops... Server is not running: %v", err)
	}
}
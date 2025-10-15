package main

import (
	"github.com/enxg/skyticket/internal/router"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app, router.Controllers{})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}

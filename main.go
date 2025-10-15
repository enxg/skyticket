package main

import (
	"github.com/enxg/skyticket/internal/router"
	"github.com/enxg/skyticket/pkg/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()

	app := fiber.New(fiber.Config{
		StructValidator: validator.NewStructValidator(),
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}

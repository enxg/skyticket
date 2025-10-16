package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bytedance/sonic"
	"github.com/enxg/skyticket/internal/controllers"
	"github.com/enxg/skyticket/internal/repositories"
	"github.com/enxg/skyticket/internal/responses"
	"github.com/enxg/skyticket/internal/router"
	"github.com/enxg/skyticket/internal/services"
	"github.com/enxg/skyticket/pkg/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/gofiber/fiber/v3/middleware/cors"
)

// @title			SkyTicket
// @version		1.0
// @description	This is the API documentation for SkyTicket.
// @license.name	MIT
// @license.url	https://github.com/enxg/skyticket/blob/main/LICENSE
// @schemes		https
// @host			skyticket.enesgenc.dev
func main() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal().Msg("MONGODB_URI environment variable not set")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to MongoDB")
	}

	db := client.Database("skyticket")

	eventRepository := repositories.NewEventRepository(db)
	eventService := services.NewEventService(eventRepository)
	eventController := controllers.NewEventController(eventService)

	app := fiber.New(fiber.Config{
		StructValidator: validator.NewStructValidator(),
		JSONEncoder:     sonic.Marshal,
		JSONDecoder:     sonic.Unmarshal,
		ErrorHandler:    errorHandler,
	})

	app.Use(cors.New())

	router.SetupRoutes(app, router.Controllers{
		EventController: eventController,
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}

func errorHandler(ctx fiber.Ctx, err error) error {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responses.ValidationErrorResponse{
				Errors: validator.ParseValidationErrors(validationErrors),
			})
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return ctx.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
			Message: "Resource not found",
		})
	}

	log.Error().
		Str("path", ctx.Path()).
		Str("type", fmt.Sprintf("%T", err)).
		Err(err).
		Send()
	return ctx.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
		Message: "Internal server error",
	})
}

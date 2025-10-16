package main

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/enxg/skyticket/internal/controllers"
	"github.com/enxg/skyticket/internal/repositories"
	"github.com/enxg/skyticket/internal/router"
	"github.com/enxg/skyticket/internal/services"
	"github.com/enxg/skyticket/pkg/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

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
	})

	router.SetupRoutes(app, router.Controllers{
		EventController: eventController,
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}

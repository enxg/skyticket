package router

import (
	"github.com/enxg/skyticket/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type Controllers struct {
	EventController controllers.EventController
}

func SetupRoutes(app *fiber.App, c Controllers) {
	app.Group("/events").
		Post("", c.EventController.CreateEvent).
		Get("/:id", c.EventController.GetEventByID).
		Get("", c.EventController.GetAllEvents).
		Patch("/:id", c.EventController.UpdateEvent).
		Delete("/:id", c.EventController.DeleteEvent)

}

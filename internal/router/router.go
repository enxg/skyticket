package router

import (
	"github.com/enxg/skyticket/docs"
	"github.com/enxg/skyticket/internal/controllers"
	"github.com/gofiber/fiber/v3"
	"github.com/yokeTH/gofiber-scalar/scalar/v3"
)

type Controllers struct {
	EventController       controllers.EventController
	TicketController      controllers.TicketController
	ReservationController controllers.ReservationController
}

func SetupRoutes(app *fiber.App, c Controllers) {
	app.Group("/events").
		Post("/", c.EventController.CreateEvent).
		Get("/:id", c.EventController.GetEventByID).
		Get("/", c.EventController.GetAllEvents).
		Patch("/:id", c.EventController.UpdateEvent).
		Delete("/:id", c.EventController.DeleteEvent)

	app.Group("/events/:eventId/tickets").
		Post("/", c.TicketController.CreateTicket).
		Get("/:id", c.TicketController.GetTicketByID).
		Get("/", c.TicketController.GetAllTickets).
		Patch("/:id", c.TicketController.UpdateTicket).
		Delete("/:id", c.TicketController.DeleteTicket)

	app.Group("/events/:eventId/tickets/:ticketId/reservation").
		Post("/", c.ReservationController.CreateReservation).
		Get("/", c.ReservationController.GetReservationByID).
		Patch("/", c.ReservationController.UpdateReservation).
		Delete("/", c.ReservationController.DeleteReservation)

	app.Group("/docs").
		Use(scalar.New(scalar.Config{
			FileContentString: docs.SwaggerInfo.ReadDoc(),
			Path:              "",
			Title:             "SkyTicket API Documentation",
		}))
}

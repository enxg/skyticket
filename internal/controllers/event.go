package controllers

import (
	"time"

	"github.com/enxg/skyticket/internal/requests"
	"github.com/enxg/skyticket/internal/services"
	"github.com/gofiber/fiber/v3"
)

type EventController interface {
	CreateEvent(c fiber.Ctx) error
	GetEventByID(c fiber.Ctx) error
	GetAllEvents(c fiber.Ctx) error
	UpdateEvent(c fiber.Ctx) error
	DeleteEvent(c fiber.Ctx) error
}

type eventController struct {
	eventService services.EventService
}

func NewEventController(eventService services.EventService) EventController {
	return &eventController{
		eventService: eventService,
	}
}

func (s *eventController) CreateEvent(c fiber.Ctx) error {
	var data requests.CreateEventRequest

	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	date, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		return err
	}

	resp, err := s.eventService.CreateEvent(c.Context(), data.Name, date, data.Venue)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (s *eventController) GetEventByID(c fiber.Ctx) error {
	id := c.Params("id")
	resp, err := s.eventService.GetEventByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

func (s *eventController) GetAllEvents(c fiber.Ctx) error {
	resp, err := s.eventService.GetAllEvents(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

func (s *eventController) UpdateEvent(c fiber.Ctx) error {
	id := c.Params("id")

	var data requests.UpdateEventRequest
	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	date, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		return err
	}

	resp, err := s.eventService.UpdateEvent(c.Context(), id, data.Name, date, data.Venue)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

func (s *eventController) DeleteEvent(c fiber.Ctx) error {
	id := c.Params("id")
	err := s.eventService.DeleteEvent(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

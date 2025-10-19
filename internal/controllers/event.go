package controllers

import (
	"time"

	"github.com/enxg/skyticket/internal/requests"
	"github.com/enxg/skyticket/internal/responses"
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

// CreateEvent godoc
//
//	@Summary		Create a new event
//	@Description	Create a new event with the provided details
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			event	body		requests.CreateEventRequest	true	"Event details"
//	@Success		201		{object}	models.Event
//	@Failure		400		{object}	responses.ValidationErrorResponse
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/events [post]
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

	if time.Now().After(date) {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "Event date cannot be in the past",
		})
	}

	resp, err := s.eventService.CreateEvent(c.Context(), data.Name, date, data.Venue)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetEventByID godoc
//
//	@Summary		Get event by ID
//	@Description	Get details of an event by its ID
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Event ID"
//	@Success		200	{object}	models.Event
//	@Failure		400	{object}	responses.ValidationErrorResponse
//	@Failure		404	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/events/{id} [get]
func (s *eventController) GetEventByID(c fiber.Ctx) error {
	id := c.Params("id")
	resp, err := s.eventService.GetEventByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// GetAllEvents godoc
//
//	@Summary		Get all events
//	@Description	Retrieve a list of all events with their details
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Event
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/events [get]
func (s *eventController) GetAllEvents(c fiber.Ctx) error {
	resp, err := s.eventService.GetAllEvents(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// UpdateEvent godoc
//
//	@Summary		Update an existing event
//	@Description	Update the details of an existing event by its ID
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Event ID"
//	@Param			event	body		requests.UpdateEventRequest	true	"Updated event details"
//	@Success		200		{object}	models.Event
//	@Failure		400		{object}	responses.ValidationErrorResponse
//	@Failure		404		{object}	responses.ErrorResponse
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/events/{id} [patch]
func (s *eventController) UpdateEvent(c fiber.Ctx) error {
	id := c.Params("id")

	var data requests.UpdateEventRequest
	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	var date time.Time
	if data.Date != "" {
		date, err = time.Parse(time.RFC3339, data.Date)
		if err != nil {
			return err
		}
	}

	if !date.IsZero() && time.Now().After(date) {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "Event date cannot be in the past",
		})
	}

	resp, err := s.eventService.UpdateEvent(c.Context(), id, data.Name, date, data.Venue)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// DeleteEvent godoc
//
//	@Summary		Delete an event
//	@Description	Delete an event by its ID
//	@Tags			Events
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Event ID"
//	@Success		204	{object}	nil		"Deleted"
//	@Failure		404	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/events/{id} [delete]
func (s *eventController) DeleteEvent(c fiber.Ctx) error {
	id := c.Params("id")
	err := s.eventService.DeleteEvent(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

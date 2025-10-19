package controllers

import (
	"errors"

	"github.com/enxg/skyticket/internal/requests"
	"github.com/enxg/skyticket/internal/responses"
	"github.com/enxg/skyticket/internal/services"
	"github.com/gofiber/fiber/v3"
)

type TicketController interface {
	CreateTicket(c fiber.Ctx) error
	GetTicketByID(c fiber.Ctx) error
	GetAllTickets(c fiber.Ctx) error
	UpdateTicket(c fiber.Ctx) error
	DeleteTicket(c fiber.Ctx) error
}

type ticketController struct {
	ticketService services.TicketService
}

func NewTicketController(ticketService services.TicketService) TicketController {
	return &ticketController{
		ticketService: ticketService,
	}
}

// CreateTicket godoc
//
//	@Summary		Create a ticket
//	@Description	Create a ticket with the provided details
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path		string							true	"Event ID"
//	@Param			ticket	body		requests.CreateTicketRequest	true	"Ticket details"
//	@Success		201		{object}	models.Ticket
//	@Failure		400		{object}	responses.ValidationErrorResponse
//	@Failure		404		{object}	responses.ErrorResponse	"Event not found"
//	@Failure		409		{object}	responses.ErrorResponse	"Seat number is already taken"
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets [post]
func (t *ticketController) CreateTicket(c fiber.Ctx) error {
	var data requests.CreateTicketRequest
	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	eventId := c.Params("eventId")

	resp, err := t.ticketService.CreateTicket(c.Context(), eventId, data.SeatNumber, data.Price)
	if err != nil {
		if errors.Is(err, services.ErrEventNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Event not found",
			})
		}

		if errors.Is(err, services.ErrSeatNumberTaken) {
			return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{
				Message: "Seat number is already taken",
			})
		}

		if errors.Is(err, services.ErrEventAlreadyPassed) {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Event date has already passed",
			})
		}

		return err
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetTicketByID godoc
//
//	@Summary		Get ticket by ID
//	@Description	Get details of a ticket by its ID
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path		string	true	"Event ID"
//	@Param			id		path		string	true	"Ticket ID"
//	@Success		200		{object}	models.Ticket
//	@Failure		404		{object}	responses.ErrorResponse	"Ticket not found for the given event"
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{id} [get]
func (t *ticketController) GetTicketByID(c fiber.Ctx) error {
	eventId := c.Params("eventId")
	ticketId := c.Params("id")

	resp, err := t.ticketService.GetTicket(c.Context(), ticketId, eventId)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// GetAllTickets godoc
//
//	@Summary		Get all tickets for an event
//	@Description	Retrieve a list of all tickets for an event with their details
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path		string	true	"Event ID"
//	@Success		200		{array}		models.Ticket
//	@Failure		404		{object}	responses.ErrorResponse	"Event not found"
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets [get]
func (t *ticketController) GetAllTickets(c fiber.Ctx) error {
	eventId := c.Params("eventId")

	resp, err := t.ticketService.GetTicketsByEvent(c.Context(), eventId)
	if err != nil {
		if errors.Is(err, services.ErrEventNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Event not found",
			})
		}

		return err
	}

	return c.JSON(resp)
}

// UpdateTicket godoc
//
//	@Summary		Update an existing ticket
//	@Description	Update the details of an existing ticket by its ID
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path		string							true	"Event ID"
//	@Param			id		path		string							true	"Ticket ID"
//	@Param			ticket	body		requests.UpdateTicketRequest	true	"Updated ticket details"
//	@Success		200		{object}	models.Ticket
//	@Failure		400		{object}	responses.ValidationErrorResponse
//	@Failure		404		{object}	responses.ErrorResponse	"Ticket not found for the given event"
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{id} [patch]
func (t *ticketController) UpdateTicket(c fiber.Ctx) error {
	eventId := c.Params("eventId")
	ticketId := c.Params("id")

	var data requests.UpdateTicketRequest
	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	resp, err := t.ticketService.UpdateTicket(c.Context(), ticketId, eventId, data.SeatNumber, data.Price)
	if err != nil {
		if errors.Is(err, services.ErrSeatNumberTaken) {
			return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{
				Message: "Seat number is already taken",
			})
		}

		return err
	}

	return c.JSON(resp)
}

// DeleteTicket godoc
//
//	@Summary		Delete a ticket
//	@Description	Delete a ticket by its ID
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//	@Param			eventId	path	string	true	"Event ID"
//	@Param			id		path	string	true	"Ticket ID"
//	@Success		204
//	@Failure		404	{object}	responses.ErrorResponse	"Ticket not found for the given event"
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{id} [delete]
func (t *ticketController) DeleteTicket(c fiber.Ctx) error {
	eventId := c.Params("eventId")
	ticketId := c.Params("id")

	err := t.ticketService.DeleteTicket(c.Context(), ticketId, eventId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

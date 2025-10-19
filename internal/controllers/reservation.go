package controllers

import (
	"errors"

	"github.com/enxg/skyticket/internal/requests"
	"github.com/enxg/skyticket/internal/responses"
	"github.com/enxg/skyticket/internal/services"
	"github.com/gofiber/fiber/v3"
)

type ReservationController interface {
	CreateReservation(c fiber.Ctx) error
	GetReservationByID(c fiber.Ctx) error
	UpdateReservation(c fiber.Ctx) error
	DeleteReservation(c fiber.Ctx) error
}

type reservationController struct {
	reservationService services.ReservationService
}

func NewReservationController(reservationService services.ReservationService) ReservationController {
	return &reservationController{
		reservationService: reservationService,
	}
}

// CreateReservation godoc
//
//	@Summary		Create a reservation
//	@Description	Create a reservation for a ticket
//	@Tags			Reservations
//	@Accept			json
//	@Produce		json
//	@Param			eventId		path		string								true	"Event ID"
//	@Param			ticketId	path		string								true	"Ticket ID"
//	@Param			reservation	body		requests.CreateReservationRequest	true	"Reservation details"
//	@Success		201			{object}	models.Reservation
//	@Failure		400			{object}	responses.ValidationErrorResponse
//	@Failure		404			{object}	responses.ErrorResponse	"Ticket/event not found"
//	@Failure		409			{object}	responses.ErrorResponse	"Ticket is already reserved / Event date has already passed"
//	@Failure		500			{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{ticketId}/reservation [post]
func (r *reservationController) CreateReservation(c fiber.Ctx) error {
	var data requests.CreateReservationRequest
	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	eventID := c.Params("eventId")
	ticketID := c.Params("ticketId")

	resp, err := r.reservationService.CreateReservation(c.Context(), eventID, ticketID, data.CustomerName)
	if err != nil {
		if errors.Is(err, services.ErrTicketNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{
				Message: "Ticket not found",
			})
		}

		if errors.Is(err, services.ErrTicketAlreadyReserved) {
			return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{
				Message: "Ticket is already reserved",
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

// GetReservationByID godoc
//
//	@Summary		Get reservation
//	@Description	Get reservation details for a ticket
//	@Tags			Reservations
//	@Accept			json
//	@Produce		json
//	@Param			eventId		path		string	true	"Event ID"
//	@Param			ticketId	path		string	true	"Ticket ID"
//	@Success		200			{object}	models.Reservation
//	@Failure		404			{object}	responses.ErrorResponse
//	@Failure		500			{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{ticketId}/reservation [get]
func (r *reservationController) GetReservationByID(c fiber.Ctx) error {
	eventID := c.Params("eventId")
	ticketID := c.Params("ticketId")

	resp, err := r.reservationService.GetReservation(c.Context(), eventID, ticketID)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// UpdateReservation godoc
//
//	@Summary		Update a reservation
//	@Description	Update details of an existing reservation
//	@Tags			Reservations
//	@Accept			json
//	@Produce		json
//	@Param			eventId		path		string								true	"Event ID"
//	@Param			ticketId	path		string								true	"Ticket ID"
//	@Param			reservation	body		requests.UpdateReservationRequest	true	"Updated reservation details"
//	@Success		200			{object}	models.Reservation
//	@Failure		400			{object}	responses.ValidationErrorResponse
//	@Failure		404			{object}	responses.ErrorResponse
//	@Failure		500			{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{ticketId}/reservation [patch]
func (r *reservationController) UpdateReservation(c fiber.Ctx) error {
	eventID := c.Params("eventId")
	ticketID := c.Params("ticketId")

	var data requests.UpdateReservationRequest
	err := c.Bind().Body(&data)
	if err != nil {
		return err
	}

	resp, err := r.reservationService.UpdateReservation(c.Context(), eventID, ticketID, data.CustomerName)
	if err != nil {
		if errors.Is(err, services.ErrEventAlreadyPassed) {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Event date has already passed",
			})
		}

		return err
	}

	return c.JSON(resp)
}

// DeleteReservation godoc
//
//	@Summary		Cancel a reservation
//	@Description	Cancel and delete an existing reservation
//	@Tags			Reservations
//	@Accept			json
//	@Produce		json
//	@Param			eventId		path	string	true	"Event ID"
//	@Param			ticketId	path	string	true	"Ticket ID"
//	@Success		204
//	@Failure		404	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/events/{eventId}/tickets/{ticketId}/reservation [delete]
func (r *reservationController) DeleteReservation(c fiber.Ctx) error {
	eventID := c.Params("eventId")
	ticketID := c.Params("ticketId")

	err := r.reservationService.CancelReservation(c.Context(), eventID, ticketID)
	if err != nil {
		if errors.Is(err, services.ErrEventAlreadyPassed) {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
				Message: "Event date has already passed",
			})
		}

		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

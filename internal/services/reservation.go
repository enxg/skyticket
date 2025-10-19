package services

import (
	"context"
	"errors"
	"time"

	"github.com/enxg/skyticket/internal/models"
	"github.com/enxg/skyticket/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, eventID string, ticketID string, customerName string) (models.Reservation, error)
	GetReservation(ctx context.Context, reservationID string, ticketID string) (models.Reservation, error)
	UpdateReservation(ctx context.Context, reservationID string, ticketID string, customerName string) (models.Reservation, error)
	CancelReservation(ctx context.Context, reservationID string, ticketID string) error
}

type reservationService struct {
	reservationRepository repositories.ReservationRepository
	ticketRepository      repositories.TicketRepository
	eventRepository       repositories.EventRepository
	mongoClient           *mongo.Client
}

var (
	ErrTicketNotFound        = errors.New("ticket not found")
	ErrTicketAlreadyReserved = errors.New("ticket already reserved")
	ErrEventAlreadyPassed    = errors.New("event date has already passed")
)

func NewReservationService(reservationRepository repositories.ReservationRepository, ticketRepository repositories.TicketRepository, eventRepository repositories.EventRepository, mongoClient *mongo.Client) ReservationService {
	return &reservationService{
		reservationRepository: reservationRepository,
		ticketRepository:      ticketRepository,
		eventRepository:       eventRepository,
		mongoClient:           mongoClient,
	}
}

func (r *reservationService) CreateReservation(ctx context.Context, eventID string, ticketID string, customerName string) (models.Reservation, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return models.Reservation{}, err
	}

	ticketOid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return models.Reservation{}, err
	}

	event, err := r.eventRepository.FindOneByID(ctx, eventOid)
	if err != nil {
		return models.Reservation{}, err
	}

	ti := time.Now()

	if event.Date.Before(ti) {
		return models.Reservation{}, ErrEventAlreadyPassed
	}

	tx, err := r.mongoClient.StartSession()
	if err != nil {
		return models.Reservation{}, err
	}
	defer tx.EndSession(ctx)

	reservation, err := tx.WithTransaction(ctx, func(txCtx context.Context) (interface{}, error) {
		reserveTicket, err := r.ticketRepository.AttemptToReserve(txCtx, event.ID, ticketOid)
		if err != nil {
			return models.Reservation{}, err
		}

		if !reserveTicket.TicketFound {
			return models.Reservation{}, ErrTicketNotFound
		} else if !reserveTicket.Reserved {
			return models.Reservation{}, ErrTicketAlreadyReserved
		}

		return r.reservationRepository.Create(txCtx, models.Reservation{
			TicketID:        ticketOid,
			CustomerName:    customerName,
			Status:          models.ReservationStatusActive,
			ReservationDate: ti,
		})
	})

	return reservation.(models.Reservation), err
}

func (r *reservationService) GetReservation(ctx context.Context, eventID string, ticketID string) (models.Reservation, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return models.Reservation{}, err
	}

	ticketOid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return models.Reservation{}, err
	}

	ticket, err := r.ticketRepository.FindOne(ctx, models.Ticket{
		ID:      ticketOid,
		EventID: eventOid,
	})
	if err != nil {
		return models.Reservation{}, err
	}

	return r.reservationRepository.FindOne(ctx, models.Reservation{
		TicketID: ticket.ID,
	})
}

func (r *reservationService) UpdateReservation(ctx context.Context, eventID string, ticketID string, customerName string) (models.Reservation, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return models.Reservation{}, err
	}

	ticketOid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return models.Reservation{}, err
	}

	ticket, err := r.ticketRepository.FindOne(ctx, models.Ticket{
		ID:      ticketOid,
		EventID: eventOid,
	})
	if err != nil {
		return models.Reservation{}, err
	}

	return r.reservationRepository.Update(ctx, models.Reservation{
		TicketID:     ticket.ID,
		CustomerName: customerName,
	})
}

func (r *reservationService) CancelReservation(ctx context.Context, eventID string, ticketID string) error {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}

	ticketOid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return err
	}

	ticket, err := r.ticketRepository.FindOne(ctx, models.Ticket{
		ID:      ticketOid,
		EventID: eventOid,
	})
	if err != nil {
		return err
	}

	return r.reservationRepository.Delete(ctx, models.Reservation{
		TicketID: ticket.ID,
	})
}

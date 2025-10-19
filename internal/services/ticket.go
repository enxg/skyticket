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

type TicketService interface {
	CreateTicket(ctx context.Context, eventID string, seatNumber string, price int) (models.Ticket, error)
	GetTicket(ctx context.Context, ticketID string, eventID string) (models.Ticket, error)
	GetTicketsByEvent(ctx context.Context, eventID string) ([]models.Ticket, error)
	UpdateTicket(ctx context.Context, ticketID string, eventID string, seatNumber string, price int) (models.Ticket, error)
	DeleteTicket(ctx context.Context, ticketID string, eventID string) error
}

type ticketService struct {
	ticketRepository      repositories.TicketRepository
	eventRepository       repositories.EventRepository
	reservationRepository repositories.ReservationRepository
	mongoClient           *mongo.Client
}

var (
	ErrEventNotFound   = errors.New("event not found")
	ErrSeatNumberTaken = errors.New("seat number is already taken")
)

func NewTicketService(ticketRepository repositories.TicketRepository, eventRepository repositories.EventRepository, reservationRepository repositories.ReservationRepository, mongoClient *mongo.Client) TicketService {
	return &ticketService{
		ticketRepository:      ticketRepository,
		eventRepository:       eventRepository,
		reservationRepository: reservationRepository,
		mongoClient:           mongoClient,
	}
}

func (t *ticketService) CreateTicket(ctx context.Context, eventID string, seatNumber string, price int) (models.Ticket, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return models.Ticket{}, err
	}

	event, err := t.eventRepository.FindOneByID(ctx, eventOid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Ticket{}, ErrEventNotFound
		}
		return models.Ticket{}, err
	}

	if time.Now().After(event.Date) {
		return models.Ticket{}, ErrEventAlreadyPassed
	}

	existingTickets, err := t.ticketRepository.Find(ctx, models.Ticket{
		EventID:    event.ID,
		SeatNumber: seatNumber,
	})
	if err != nil {
		return models.Ticket{}, err
	}
	if len(existingTickets) > 0 {
		return models.Ticket{}, ErrSeatNumberTaken
	}

	return t.ticketRepository.Create(ctx, models.Ticket{
		EventID:    event.ID,
		SeatNumber: seatNumber,
		Price:      price,
		Status:     models.TicketStatusAvailable,
	})
}

func (t *ticketService) GetTicket(ctx context.Context, ticketID string, eventID string) (models.Ticket, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return models.Ticket{}, err
	}

	oid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return models.Ticket{}, err
	}

	return t.ticketRepository.FindOne(ctx, models.Ticket{
		ID:      oid,
		EventID: eventOid,
	})
}

func (t *ticketService) GetTicketsByEvent(ctx context.Context, eventID string) ([]models.Ticket, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return nil, err
	}

	_, err = t.eventRepository.FindOneByID(ctx, eventOid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrEventNotFound
		}
		return nil, err
	}

	return t.ticketRepository.Find(ctx, models.Ticket{
		EventID: eventOid,
	})
}

func (t *ticketService) UpdateTicket(ctx context.Context, ticketID string, eventID string, seatNumber string, price int) (models.Ticket, error) {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return models.Ticket{}, err
	}

	oid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return models.Ticket{}, err
	}

	if seatNumber != "" {
		searchSeat, err := t.ticketRepository.FindOne(ctx, models.Ticket{
			EventID:    eventOid,
			SeatNumber: seatNumber,
		})
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return models.Ticket{}, err
		}

		if searchSeat.ID != oid && searchSeat.SeatNumber == seatNumber {
			return models.Ticket{}, ErrSeatNumberTaken
		}
	}

	return t.ticketRepository.Update(ctx, models.Ticket{
		ID:         oid,
		EventID:    eventOid,
		SeatNumber: seatNumber,
		Price:      price,
	})
}

func (t *ticketService) DeleteTicket(ctx context.Context, ticketID string, eventID string) error {
	eventOid, err := bson.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}

	oid, err := bson.ObjectIDFromHex(ticketID)
	if err != nil {
		return err
	}

	tx, err := t.mongoClient.StartSession()
	if err != nil {
		return err
	}

	_, err = tx.WithTransaction(ctx, func(txCtx context.Context) (any, error) {
		err := t.ticketRepository.Delete(txCtx, models.Ticket{
			ID:      oid,
			EventID: eventOid,
		})
		if err != nil {
			return nil, err
		}

		return nil, t.reservationRepository.DeleteMany(txCtx, models.Reservation{
			TicketID: oid,
		})
	})

	return err
}

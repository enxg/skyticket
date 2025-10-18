package services

import (
	"context"
	"errors"

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
	ticketRepository repositories.TicketRepository
	eventRepository  repositories.EventRepository
}

var (
	ErrEventNotFound   = errors.New("event not found")
	ErrSeatNumberTaken = errors.New("seat number is already taken")
)

func NewTicketService(ticketRepository repositories.TicketRepository, eventRepository repositories.EventRepository) TicketService {
	return &ticketService{
		ticketRepository: ticketRepository,
		eventRepository:  eventRepository,
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

	return t.ticketRepository.Delete(ctx, models.Ticket{
		ID:      oid,
		EventID: eventOid,
	})
}

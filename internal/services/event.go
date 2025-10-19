package services

import (
	"context"
	"time"

	"github.com/enxg/skyticket/internal/models"
	"github.com/enxg/skyticket/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EventService interface {
	CreateEvent(ctx context.Context, name string, date time.Time, venue string) (models.Event, error)
	GetEventByID(ctx context.Context, id string) (models.Event, error)
	GetAllEvents(ctx context.Context) ([]models.Event, error)
	UpdateEvent(ctx context.Context, id string, name string, date time.Time, venue string) (models.Event, error)
	DeleteEvent(ctx context.Context, id string) error
}

type eventService struct {
	eventRepository       repositories.EventRepository
	ticketRepository      repositories.TicketRepository
	reservationRepository repositories.ReservationRepository
	mongoClient           *mongo.Client
}

func NewEventService(eventRepository repositories.EventRepository, ticketRepository repositories.TicketRepository, reservationRepository repositories.ReservationRepository, mongoClient *mongo.Client) EventService {
	return &eventService{
		eventRepository:       eventRepository,
		ticketRepository:      ticketRepository,
		reservationRepository: reservationRepository,
		mongoClient:           mongoClient,
	}
}

func (e *eventService) CreateEvent(ctx context.Context, name string, date time.Time, venue string) (models.Event, error) {
	res, err := e.eventRepository.Create(ctx, models.Event{
		Name:  name,
		Date:  date,
		Venue: venue,
	})
	if err != nil {
		return models.Event{}, err
	}

	return res, nil
}

func (e *eventService) GetEventByID(ctx context.Context, id string) (models.Event, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	return e.eventRepository.FindOneByID(ctx, oid)
}

func (e *eventService) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	return e.eventRepository.Find(ctx, models.Event{})
}

func (e *eventService) UpdateEvent(ctx context.Context, id string, name string, date time.Time, venue string) (models.Event, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	return e.eventRepository.Update(ctx, models.Event{
		ID:    oid,
		Name:  name,
		Date:  date,
		Venue: venue,
	})
}

func (e *eventService) DeleteEvent(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	tx, err := e.mongoClient.StartSession()
	if err != nil {
		return err
	}
	_, err = tx.WithTransaction(ctx, func(txCtx context.Context) (any, error) {
		err := e.reservationRepository.DeleteMany(txCtx, models.Reservation{
			EventID: oid,
		})
		if err != nil {
			return nil, err
		}

		err = e.ticketRepository.DeleteMany(txCtx, models.Ticket{
			EventID: oid,
		})
		if err != nil {
			return nil, err
		}

		err = e.eventRepository.Delete(ctx, oid)
		return nil, err
	})

	return err
}

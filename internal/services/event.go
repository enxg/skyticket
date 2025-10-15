package services

import (
	"context"
	"time"

	"github.com/enxg/skyticket/internal/models"
	"github.com/enxg/skyticket/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventService interface {
	CreateEvent(ctx context.Context, name string, date time.Time, venue string) (models.Event, error)
	GetEventByID(ctx context.Context, id string) (models.Event, error)
	GetAllEvents(ctx context.Context) ([]models.Event, error)
	UpdateEvent(ctx context.Context, id string, name string, date time.Time, venue string) (models.Event, error)
	DeleteEvent(ctx context.Context, id string) error
}

type eventService struct {
	eventRepository repositories.EventRepository
}

func NewEventService(eventRepository repositories.EventRepository) EventService {
	return &eventService{
		eventRepository: eventRepository,
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
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	return e.eventRepository.GetByID(ctx, oid)
}

func (e *eventService) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	return e.eventRepository.GetAll(ctx)
}

func (e *eventService) UpdateEvent(ctx context.Context, id string, name string, date time.Time, venue string) (models.Event, error) {
	oid, err := primitive.ObjectIDFromHex(id)
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
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return e.eventRepository.Delete(ctx, oid)
}

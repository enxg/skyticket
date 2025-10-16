package repositories

import (
	"context"

	"github.com/enxg/skyticket/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EventRepository interface {
	Create(ctx context.Context, event models.Event) (models.Event, error)
	GetByID(ctx context.Context, id bson.ObjectID) (models.Event, error)
	GetAll(ctx context.Context) ([]models.Event, error)
	Update(ctx context.Context, event models.Event) (models.Event, error)
	Delete(ctx context.Context, id bson.ObjectID) error
}

type eventRepository struct {
	collection mongo.Collection
}

func NewEventRepository(db *mongo.Database) EventRepository {
	return &eventRepository{
		collection: *db.Collection("events"),
	}
}

func (e *eventRepository) Create(ctx context.Context, event models.Event) (models.Event, error) {
	res, err := e.collection.InsertOne(ctx, event)
	if err != nil {
		return models.Event{}, err
	}

	event.ID = res.InsertedID.(bson.ObjectID)
	return event, nil
}

func (e *eventRepository) GetByID(ctx context.Context, id bson.ObjectID) (models.Event, error) {
	var result models.Event
	err := e.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return models.Event{}, err
	}

	return result, nil
}

func (e *eventRepository) GetAll(ctx context.Context) ([]models.Event, error) {
	var events []models.Event
	cursor, err := e.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (e *eventRepository) Update(ctx context.Context, event models.Event) (models.Event, error) {
	filter := bson.D{{"_id", event.ID}}
	update := bson.D{{"$set", event}}

	res, err := e.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.Event{}, err
	}

	if res.MatchedCount == 0 {
		return models.Event{}, mongo.ErrNoDocuments
	}

	return event, nil
}

func (e *eventRepository) Delete(ctx context.Context, id bson.ObjectID) error {
	res, err := e.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

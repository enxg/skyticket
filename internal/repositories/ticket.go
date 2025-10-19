package repositories

import (
	"context"

	"github.com/enxg/skyticket/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket models.Ticket) (models.Ticket, error)
	FindOne(ctx context.Context, filter models.Ticket) (models.Ticket, error)
	Find(ctx context.Context, filter models.Ticket) ([]models.Ticket, error)
	Update(ctx context.Context, ticket models.Ticket) (models.Ticket, error)
	Delete(ctx context.Context, filter models.Ticket) error
	AttemptToReserve(ctx context.Context, eventID bson.ObjectID, ticketID bson.ObjectID) (TicketReservationAttemptResult, error)
}

type ticketRepository struct {
	collection mongo.Collection
}

type TicketReservationAttemptResult struct {
	TicketFound bool
	Reserved    bool
}

func NewTicketRepository(db *mongo.Database) TicketRepository {
	return &ticketRepository{
		collection: *db.Collection("tickets"),
	}
}

func (t *ticketRepository) Create(ctx context.Context, ticket models.Ticket) (models.Ticket, error) {
	res, err := t.collection.InsertOne(ctx, ticket)
	if err != nil {
		return models.Ticket{}, err
	}

	ticket.ID = res.InsertedID.(bson.ObjectID)
	return ticket, nil
}

func (t *ticketRepository) FindOne(ctx context.Context, filter models.Ticket) (models.Ticket, error) {
	var result models.Ticket
	err := t.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.Ticket{}, err
	}

	return result, nil
}

func (t *ticketRepository) Find(ctx context.Context, filter models.Ticket) ([]models.Ticket, error) {
	tickets := make([]models.Ticket, 0)

	cursor, err := t.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (t *ticketRepository) Update(ctx context.Context, ticket models.Ticket) (models.Ticket, error) {
	filter := bson.M{
		"_id":      ticket.ID,
		"event_id": ticket.EventID,
	}
	update := bson.M{"$set": ticket}

	res, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.Ticket{}, err
	}

	if res.MatchedCount == 0 {
		return models.Ticket{}, mongo.ErrNoDocuments
	}

	return t.FindOne(ctx, models.Ticket{
		ID: ticket.ID,
	})
}

func (t *ticketRepository) Delete(ctx context.Context, filter models.Ticket) error {
	res, err := t.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (t *ticketRepository) AttemptToReserve(ctx context.Context, eventID bson.ObjectID, ticketID bson.ObjectID) (TicketReservationAttemptResult, error) {
	filter := bson.M{
		"_id":      ticketID,
		"event_id": eventID,
	}
	update := []bson.M{
		{
			"$set": bson.M{
				"status": bson.M{
					"$cond": bson.M{
						"if":   bson.M{"$eq": bson.A{"$status", models.TicketStatusAvailable}},
						"then": models.TicketStatusReserved,
						"else": "$status",
					},
				},
			},
		},
	}

	res, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return TicketReservationAttemptResult{}, err
	}

	return TicketReservationAttemptResult{
		TicketFound: res.MatchedCount > 0,
		Reserved:    res.ModifiedCount > 0,
	}, nil
}

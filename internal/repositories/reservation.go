package repositories

import (
	"context"

	"github.com/enxg/skyticket/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation models.Reservation) (models.Reservation, error)
	FindOne(ctx context.Context, filter models.Reservation) (models.Reservation, error)
	Update(ctx context.Context, reservation models.Reservation) (models.Reservation, error)
	Delete(ctx context.Context, filter models.Reservation) error
}

type reservationRepository struct {
	collection *mongo.Collection
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{
		collection: db.Collection("reservations"),
	}
}

func (r *reservationRepository) Create(ctx context.Context, reservation models.Reservation) (models.Reservation, error) {
	res, err := r.collection.InsertOne(ctx, reservation)
	if err != nil {
		return models.Reservation{}, err
	}

	reservation.ID = res.InsertedID.(bson.ObjectID)
	return reservation, nil
}

func (r *reservationRepository) FindOne(ctx context.Context, filter models.Reservation) (models.Reservation, error) {
	var result models.Reservation
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.Reservation{}, err
	}

	return result, nil
}

func (r *reservationRepository) Update(ctx context.Context, reservation models.Reservation) (models.Reservation, error) {
	filter := bson.M{
		"ticket_id": reservation.TicketID,
	}
	update := bson.M{"$set": reservation}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.Reservation{}, err
	}

	if res.MatchedCount == 0 {
		return models.Reservation{}, mongo.ErrNoDocuments
	}

	return r.FindOne(ctx, models.Reservation{
		TicketID: reservation.TicketID,
	})
}

func (r *reservationRepository) Delete(ctx context.Context, filter models.Reservation) error {
	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

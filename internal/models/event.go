package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Date  time.Time          `json:"date"`
	Venue string             `json:"venue"`
}

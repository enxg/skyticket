package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Date  time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Venue string             `json:"venue,omitempty" bson:"venue,omitempty"`
}

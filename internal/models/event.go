package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Event struct {
	ID    bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty"`
	Date  time.Time     `json:"date,omitempty" bson:"date,omitempty"`
	Venue string        `json:"venue,omitempty" bson:"venue,omitempty"`
}

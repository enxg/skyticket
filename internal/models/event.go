package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Event struct {
	ID    bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"68f0c6a8f5673dc0ec646731" extensions:"x-order=0"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty" example:"FORMULA 1 ETIHAD AIRWAYS ABU DHABI GRAND PRIX 2025" extensions:"x-order=1"`
	Date  time.Time     `json:"date,omitempty" bson:"date,omitempty" example:"2025-12-07T19:00:00Z" extensions:"x-order=2"`
	Venue string        `json:"venue,omitempty" bson:"venue,omitempty" example:"YTÜ Davutpaşa Tarihi Hamam" extensions:"x-order=3"`
}

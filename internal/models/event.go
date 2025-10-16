package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Event struct {
	ID    bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"68f0c6a8f5673dc0ec646731"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty" example:"FIBA EuroBasket 2025 Finali"`
	Date  time.Time     `json:"date,omitempty" bson:"date,omitempty" example:"2025-09-14T21:00:00+03:00"`
	Venue string        `json:"venue,omitempty" bson:"venue,omitempty" example:"YTÜ Davutpaşa Tarihi Hamam"`
}

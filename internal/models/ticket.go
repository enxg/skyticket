package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TicketStatus string

const (
	TicketStatusAvailable TicketStatus = "available"
	TicketStatusReserved  TicketStatus = "reserved"
)

type Ticket struct {
	ID         bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	EventID    bson.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	SeatNumber string        `json:"seat_number,omitempty" bson:"seat_number,omitempty"`
	Price      float64       `json:"price,omitempty" bson:"price,omitempty"`
	Status     TicketStatus  `json:"status,omitempty" bson:"status,omitempty"`
	Price      int           `json:"price,omitempty" bson:"price,omitempty" example:"4999"`
}

package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TicketStatus string

const (
	TicketStatusAvailable TicketStatus = "AVAILABLE"
	TicketStatusReserved  TicketStatus = "RESERVED"
)

// TODO: Might add an option to specify currency

type Ticket struct {
	ID         bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"68f2ab0516a352dc8f40c543" extensions:"x-order=0"`
	EventID    bson.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty" example:"68f0c6a8f5673dc0ec646731" extensions:"x-order=1"`
	SeatNumber string        `json:"seat_number,omitempty" bson:"seat_number,omitempty" example:"A12" extensions:"x-order=2"`
	Price      int           `json:"price,omitempty" bson:"price,omitempty" example:"4999" extensions:"x-order=3"`
	Status     TicketStatus  `json:"status,omitempty" bson:"status,omitempty" example:"AVAILABLE" extensions:"x-order=4"`
}

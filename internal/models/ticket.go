package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketStatus string

const (
	TicketStatusAvailable TicketStatus = "available"
	TicketStatusReserved  TicketStatus = "reserved"
)

type Ticket struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	EventID    primitive.ObjectID `json:"event_id,omitempty" bson:"event_id,omitempty"`
	SeatNumber string             `json:"seat_number,omitempty" bson:"seat_number,omitempty"`
	Price      float64            `json:"price,omitempty" bson:"price,omitempty"`
	Status     TicketStatus       `json:"status,omitempty" bson:"status,omitempty"`
}

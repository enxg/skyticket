package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketStatus string

const (
	TicketStatusAvailable TicketStatus = "available"
	TicketStatusReserved  TicketStatus = "reserved"
)

type Ticket struct {
	ID         primitive.ObjectID `json:"id"`
	EventID    primitive.ObjectID `json:"event_id"`
	SeatNumber string             `json:"seat_number"`
	Price      float64            `json:"price"`
	Status     TicketStatus       `json:"status"`
}

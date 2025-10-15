package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationStatus string

const (
	ReservationStatusActive    ReservationStatus = "active"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)

type Reservation struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TicketID        primitive.ObjectID `json:"ticket_id,omitempty" bson:"ticket_id,omitempty"`
	CustomerName    string             `json:"customer_name,omitempty" bson:"customer_name,omitempty"`
	ReservationDate time.Time          `json:"reservation_date,omitempty" bson:"reservation_date,omitempty"`
	Status          ReservationStatus  `json:"status,omitempty" bson:"status,omitempty"`
}

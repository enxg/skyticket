package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ReservationStatus string

const (
	ReservationStatusActive    ReservationStatus = "ACTIVE"
	ReservationStatusCancelled ReservationStatus = "CANCELLED"
)

type Reservation struct {
	ID              bson.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	TicketID        bson.ObjectID     `json:"ticket_id,omitempty" bson:"ticket_id,omitempty"`
	CustomerName    string            `json:"customer_name,omitempty" bson:"customer_name,omitempty"`
	ReservationDate time.Time         `json:"reservation_date,omitempty" bson:"reservation_date,omitempty"`
	Status          ReservationStatus `json:"status,omitempty" bson:"status,omitempty"`
}

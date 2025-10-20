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
	ID              bson.ObjectID     `json:"id,omitempty" bson:"_id,omitempty" example:"68f4fea9990e605d6589b5f3" extensions:"x-order=0"`
	EventID         bson.ObjectID     `json:"event_id,omitempty" bson:"event_id,omitempty" example:"68f0c6a8f5673dc0ec646731" extensions:"x-order=1"`
	TicketID        bson.ObjectID     `json:"ticket_id,omitempty" bson:"ticket_id,omitempty" example:"68f2ab0516a352dc8f40c543" extensions:"x-order=2"`
	CustomerName    string            `json:"customer_name,omitempty" bson:"customer_name,omitempty" example:"Lewis Hamilton" extensions:"x-order=3"`
	ReservationDate time.Time         `json:"reservation_date,omitempty" bson:"reservation_date,omitempty" example:"2025-10-19T15:00:00Z" extensions:"x-order=4"`
	Status          ReservationStatus `json:"status,omitempty" bson:"status,omitempty" example:"ACTIVE" extensions:"x-order=5"`
}

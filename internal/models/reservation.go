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
	ID              bson.ObjectID     `json:"id,omitempty" bson:"_id,omitempty" example:"68f4fea9990e605d6589b5f3"`
	TicketID        bson.ObjectID     `json:"ticket_id,omitempty" bson:"ticket_id,omitempty" example:"68f2ab0516a352dc8f40c543"`
	CustomerName    string            `json:"customer_name,omitempty" bson:"customer_name,omitempty" example:"Lewis Hamilton"`
	ReservationDate time.Time         `json:"reservation_date,omitempty" bson:"reservation_date,omitempty" example:"2025-10-1T15:00:00Z"`
	Status          ReservationStatus `json:"status,omitempty" bson:"status,omitempty" example:"ACTIVE"`
}

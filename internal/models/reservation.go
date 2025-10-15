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
	ID              primitive.ObjectID `json:"id"`
	TicketID        primitive.ObjectID `json:"ticket_id"`
	CustomerName    string             `json:"customer_name"`
	ReservationDate time.Time          `json:"reservation_date"`
	Status          ReservationStatus  `json:"status"`
}

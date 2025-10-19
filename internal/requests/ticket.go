package requests

type CreateTicketRequest struct {
	SeatNumber string `json:"seat_number" validate:"required,lt=256" example:"A12"`
	Price      int    `json:"price" validate:"required,gt=0" example:"4999"`
}

type UpdateTicketRequest struct {
	SeatNumber string `json:"seat_number" validate:"omitempty,lt=256" example:"A12"`
	Price      int    `json:"price" validate:"omitempty,gt=0" example:"4999"`
}

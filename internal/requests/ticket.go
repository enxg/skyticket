package requests

type CreateTicketRequest struct {
	SeatNumber string `json:"seat_number" validate:"required" example:"A12"`
	Price      int    `json:"price" validate:"required,gt=0" example:"4999"`
}

type UpdateTicketRequest struct {
	SeatNumber string `json:"seat_number" validate:"required" example:"A12"`
	Price      int    `json:"price" validate:"required,gt=0" example:"4999"`
}

package requests

type CreateReservationRequest struct {
	CustomerName string `json:"customer_name" validate:"required,lt=256" example:"Enes Genç"`
}

type UpdateReservationRequest struct {
	CustomerName string `json:"customer_name,omitempty" validate:"omitempty,lt=256" example:"Enes Genç"`
}

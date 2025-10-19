package requests

type CreateReservationRequest struct {
	CustomerName string `json:"customer_name" validate:"required" example:"Enes Genç"`
}

type UpdateReservationRequest struct {
	CustomerName string `json:"customer_name" example:"Enes Genç"`
}

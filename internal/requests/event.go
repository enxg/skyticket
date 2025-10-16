package requests

type CreateEventRequest struct {
	Name  string `json:"name" validate:"required" example:"FIBA EuroBasket 2025 Finali"`
	Date  string `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00" example:"2025-09-14T21:00:00+03:00"`
	Venue string `json:"venue" validate:"required" example:"YTÜ Davutpaşa Tarihi Hamam"`
}

type UpdateEventRequest struct {
	Name  string `json:"name" validate:"" example:"FIBA EuroBasket 2025 Finali"`
	Date  string `json:"date" validate:"datetime=2006-01-02T15:04:05Z07:00" example:"2025-09-14T21:00:00+03:00"`
	Venue string `json:"venue" validate:"" example:"YTÜ Davutpaşa Tarihi Hamam"`
}

package requests

type CreateEventRequest struct {
	Name  string `json:"name" validate:"required"`
	Date  string `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Venue string `json:"venue" validate:"required"`
}

type UpdateEventRequest struct {
	Name  string `json:"name" validate:""`
	Date  string `json:"date" validate:"datetime=2006-01-02T15:04:05Z07:00"`
	Venue string `json:"venue" validate:""`
}

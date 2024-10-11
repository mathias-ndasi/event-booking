package dtos

import "time"

type GetSingleEventParam struct {
	EventId int64 `json:"eventId" binding:"required"`
}

type CreateEventDto struct {
	Name        string    `json:"name" binding:"required" validate:"required|string|min_length:5" message:"The event name is invalid" label:"Event name"`
	Description string    `json:"description" binding:"required" validate:"required|string|min_length:5" message:"The event description is invalid" label:"Event description"`
	Location    string    `json:"location" binding:"required" validate:"required|string|min_length:3" message:"The event location is invalid" label:"Event location"`
	DateTime    time.Time `json:"dateTime" binding:"required" validate:"customTimeIsoValidator" message:"The date of the event is invalid" label:"Event timestamp"`
}

type UpdateEventDto struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
}

func (data CreateEventDto) CustomTimeIsoValidator(value string) bool {
	_, error := time.Parse(time.RFC3339, value)
	return error == nil
}

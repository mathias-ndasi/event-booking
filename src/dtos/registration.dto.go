package dtos

import "example.com/event-booking/prisma/db"

type EventRegistrationDto struct {
	EventId int64 `json:"eventId" binding:"required"`
}

type UpdateRegistrationDto struct {
	Status db.RegistrationStatus `json:"status"`
}

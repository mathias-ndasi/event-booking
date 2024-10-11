package models

import (
	"fmt"
	"time"

	"example.com/event-booking/prisma/db"
	prisma "example.com/event-booking/src/database"
	"example.com/event-booking/src/dtos"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	CustomerId  int64     `json:"customerId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func SaveEvent(authenticatedCustomerId int64, dto *dtos.CreateEventDto) (*db.EventModel, error) {
	// create an event
	fmt.Println(authenticatedCustomerId)
	fmt.Printf("%T", authenticatedCustomerId)
	prismaClient, context := prisma.GetClient()

	createdEvent, error := prismaClient.Event.CreateOne(
		db.Event.Name.Set(dto.Name),
		db.Event.Description.Set(dto.Description),
		db.Event.EventDate.Set(dto.DateTime),
		db.Event.Customer.Link(
			db.Customer.ID.Equals(int(authenticatedCustomerId)),
		),
		db.Event.Location.Set(dto.Location),
		db.Event.CustomerID.Set(int(authenticatedCustomerId)),
	).Exec(context)
	if error != nil {
		return nil, error
	}

	return createdEvent, nil
}

func GetAllEvents() ([]db.EventModel, error) {
	prismaClient, context := prisma.GetClient()

	events, error := prismaClient.Event.FindMany().Exec(context)
	if error != nil {
		return nil, error
	}

	return events, nil
}

func GetEvent(eventId int64) (*db.EventModel, error) {
	prismaClient, context := prisma.GetClient()

	event, error := prismaClient.Event.FindUnique(
		db.Event.ID.Equals(int(eventId)),
	).Exec(context)
	if error != nil {
		if error == db.ErrNotFound {
			return nil, fmt.Errorf("Event with ID %d not found", eventId)
		}

		return nil, error
	}

	return event, nil
}

func DeleteEvent(eventId int64) error {
	prismaClient, context := prisma.GetClient()

	_, error := prismaClient.Event.FindUnique(
		db.Event.ID.Equals(int(eventId)),
	).Delete().Exec(context)
	if error != nil {
		return error
	}

	return nil
}

func UpdateEvent(eventId int, dto dtos.UpdateEventDto) (*db.EventModel, error) {
	prismaClient, context := prisma.GetClient()

	_, error := prismaClient.Event.FindUnique(
		db.Event.ID.Equals(int(eventId)),
	).Exec(context)
	if error != nil {
		if error == db.ErrNotFound {
			return nil, fmt.Errorf("Event with ID %d not found", eventId)
		}

		return nil, error
	}

	event, error := prismaClient.Event.FindUnique(
		db.Event.ID.Equals(int(eventId)),
	).Update(
		db.Event.Name.Set(dto.Name),
		db.Event.Description.Set(dto.Description),
		db.Event.EventDate.Set(dto.DateTime),
		db.Event.Location.Set(dto.Location),
	).Exec(context)
	if error != nil {
		return nil, error
	}

	return event, nil
}

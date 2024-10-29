package models

import (
	"time"

	"example.com/event-booking/prisma/db"
	prisma "example.com/event-booking/src/database"
	"example.com/event-booking/src/dtos"
)

type Registration struct {
	Id         int64     `json:"id"`
	EventId    int64     `json:"eventId"`
	CustomerId int64     `json:"customerId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func New(authenticatedCustomerId int64, dto *dtos.EventRegistrationDto) (*db.RegistrationUniqueTxResult, error) {
	prismaClient, context := prisma.GetClient()

	registration := prismaClient.Registration.CreateOne(
		db.Registration.Status.Set(db.RegistrationStatusActive),
		db.Registration.Event.Link(
			db.Event.ID.Equals(int(dto.EventId)),
		),
		db.Registration.Customer.Link(
			db.Customer.ID.Equals(int(authenticatedCustomerId)),
		),
	).Tx()
	registrationHistory := prismaClient.RegistrationHistory.CreateOne(
		db.RegistrationHistory.Status.Set(db.RegistrationStatusActive),
		db.RegistrationHistory.Registration.Link(
			db.Registration.ID.Equals(registration.Result().ID),
		),
		db.RegistrationHistory.UpdateDoneBy.Link(
			db.Customer.ID.Equals(int(authenticatedCustomerId)),
		),
		db.RegistrationHistory.StartDate.Set(time.Now()),
	).Tx()

	if error := prismaClient.Prisma.Transaction(registration, registrationHistory).Exec(context); error != nil {
		return nil, error
	}

	return &registration, nil
}

func GetRegistration(registrationId int64) (*db.RegistrationModel, error) {
	prismaClient, context := prisma.GetClient()

	registration, error := prismaClient.Registration.FindFirst(
		db.Registration.ID.Equals(int(registrationId)),
	).Exec(context)

	if error != nil {
		return nil, error
	}

	return registration, nil
}

func UpdateRegistration(authenticatedCustomerId int64, registrationId int64, dto *dtos.UpdateRegistrationDto) (*db.RegistrationUniqueTxResult, error) {
	prismaClient, context := prisma.GetClient()

	updatedRegistration := prismaClient.Registration.UpsertOne(
		db.Registration.ID.Equals(int(registrationId)),
	).Update(
		db.Registration.Status.Set(dto.Status),
	).Tx()

	registrationHistory, error := prismaClient.RegistrationHistory.FindFirst(
		db.RegistrationHistory.RegistrationID.Equals(int(registrationId)),
	).Exec(context)
	if error != nil || registrationHistory == nil {
		return nil, error
	}

	updatedRegistrationHistory := prismaClient.RegistrationHistory.UpsertOne(
		db.RegistrationHistory.ID.Equals(int(registrationHistory.ID)),
	).Update(
		db.RegistrationHistory.EndDate.Set(time.Now()),
	).Tx()

	newRegistrationHistory := prismaClient.RegistrationHistory.CreateOne(
		db.RegistrationHistory.Status.Set(db.RegistrationStatusActive),
		db.RegistrationHistory.Registration.Link(
			db.Registration.ID.Equals(int(registrationId)),
		),
		db.RegistrationHistory.UpdateDoneBy.Link(
			db.Customer.ID.Equals(int(authenticatedCustomerId)),
		),
		db.RegistrationHistory.StartDate.Set(time.Now()),
	).Tx()

	error = prismaClient.Prisma.Transaction(updatedRegistration, updatedRegistrationHistory, newRegistrationHistory).Exec(context)
	if error != nil {
		return nil, error
	}

	return &updatedRegistration, nil
}

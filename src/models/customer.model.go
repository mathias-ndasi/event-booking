package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/event-booking/prisma/db"
	prisma "example.com/event-booking/src/database"
	"example.com/event-booking/src/dtos"
)

type Customer struct {
	Id           int64     `json:"id"`
	EmailAddress string    `json:"emailAddress" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func SignUp(dto dtos.SignUpDto) (*db.CustomerModel, error) {
	prismaClient, context := prisma.GetClient()

	if customer, _ := prismaClient.Customer.FindUnique(
		db.Customer.EmailAddress.Equals(dto.EmailAddress),
	).Exec(context); customer != nil {
		return nil, errors.New("ErrUniqueConstraint: Customer with email address already exists")
	}

	customer, error := prismaClient.Customer.CreateOne(
		db.Customer.EmailAddress.Set(dto.EmailAddress),
		db.Customer.PasswordHash.Set(dto.Password),
	).Exec(context)
	if error != nil {
		return nil, error
	}

	return customer, nil
}

func GetCustomers() ([]db.CustomerModel, error) {
	prismaClient, context := prisma.GetClient()
	customers, error := prismaClient.Customer.FindMany().Exec(context)
	if error != nil {
		return nil, error
	}

	return customers, nil
}

func GetCustomerFromEmailAddress(emailAddress string) (*db.CustomerModel, error) {
	prismaClient, context := prisma.GetClient()
	customer, error := prismaClient.Customer.FindUnique(
		db.Customer.EmailAddress.Equals(emailAddress),
	).Exec(context)
	if error != nil {
		if error == db.ErrNotFound {
			return nil, fmt.Errorf("ErrNotFound: %v", error)
		}
		return nil, error
	}

	return customer, nil
}

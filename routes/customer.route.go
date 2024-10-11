package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"example.com/event-booking/src/dtos"
	"example.com/event-booking/src/models"
	"example.com/event-booking/utils"
)

func signUp(context *gin.Context) {
	var signUpDto dtos.SignUpDto
	error := context.ShouldBindJSON(&signUpDto)
	if error != nil {
		log.Printf("Error occurred during customer sign up: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create customer due to client request error"})
		return
	}

	hashedPassword, error := utils.HashPassword(signUpDto.Password)
	if error != nil {
		log.Printf("Error occurred during customer sign up: %v", error)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create customer"})
		return
	}

	signUpDto.Password = hashedPassword
	customer, error := models.SignUp(signUpDto)
	if error != nil {
		if strings.Split(fmt.Sprintf("%v", error), ":")[0] == "ErrUniqueConstraint" {
			log.Printf("Customer with email address already exists: %v", signUpDto.EmailAddress)
			context.JSON(http.StatusConflict, gin.H{"message": "Customer with email address already exists"})
			return
		}

		log.Printf("Error occurred during customer sign up: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create customer"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully", "data": customer})
}

func getCustomers(context *gin.Context) {
	customers, error := models.GetCustomers()
	if error != nil {
		log.Printf("Error occurred while getting customers: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get customers"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Customers retrieved successfully", "data": customers})
}

func login(context *gin.Context) {
	var loginDto dtos.LoginDto
	error := context.ShouldBindJSON(&loginDto)
	if error != nil {
		log.Printf("Error occurred during customer login: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not login customer due to client request error"})
		return
	}

	customer, error := models.GetCustomerFromEmailAddress(loginDto.EmailAddress)
	if error != nil {
		if strings.Split(fmt.Sprintf("%v", error), ":")[0] == "ErrNotFound" {
			log.Printf("Error occurred during customer login: %v", error)
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Email or password incorrect"})
			return
		}

		log.Printf("Error occurred during customer login: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not login customer"})
		return
	}

	isPasswordValid := utils.IsPasswordHashValid(loginDto.Password, customer.PasswordHash)
	if !isPasswordValid {
		log.Printf("Error occurred during customer login: %v", error)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Email or password incorrect"})
		return
	}

	jwtToken, error := utils.GenerateToken(customer.EmailAddress, int64(customer.ID))
	if error != nil {
		log.Printf("Error occurred during customer login: %v", error)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login customer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Customer logged in successfully", "data": map[string]string{"token": jwtToken}})
}

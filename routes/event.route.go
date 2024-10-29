package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/event-booking/src/dtos"
	"example.com/event-booking/src/models"
)

func getEvents(context *gin.Context) {
	events, error := models.GetAllEvents()
	if error != nil {
		log.Printf("Error occurred while retrieving events: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve events"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Events retrieved successfully", "data": events})
}

func getEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if error != nil {
		log.Printf("Error occurred while parsing event ID: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	eventModel, error := models.GetEvent(eventId)
	if error != nil {
		log.Printf("Error occurred while retrieving event ID: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event retrieved successfully", "data": eventModel})
}

func deleteEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if error != nil {
		log.Printf("Error occurred while parsing event ID: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	error = models.DeleteEvent(eventId)
	if error != nil {
		log.Printf("Error occurred deleting event ID: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func createEvent(context *gin.Context) {
	var eventDto dtos.CreateEventDto

	error := context.ShouldBindJSON(&eventDto)
	if error != nil {
		log.Printf("Error occurred while binding JSON input: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event due to client request error"})
		return
	}

	customerId := context.GetInt64("customerId")
	eventModel, error := models.SaveEvent(customerId, &eventDto)
	if error != nil {
		log.Printf("Error occurred while creating event: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "data": &eventModel})
}

func updateEvent(context *gin.Context) {
	eventId, error := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if error != nil {
		log.Printf("Error occurred while parsing event ID: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	var eventDto dtos.UpdateEventDto
	error = context.ShouldBindJSON(&eventDto)
	if error != nil {
		log.Printf("Error occurred while updating event: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event due to client request error"})
		return
	}

	_, error = models.UpdateEvent(int(eventId), eventDto)
	if error != nil {
		log.Printf("Error occurred updating event ID: %v", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

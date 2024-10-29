package routes

import (
	"github.com/gin-gonic/gin"

	"example.com/event-booking/prisma/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/events", getEvents)
	authenticated.POST("/events", createEvent)
	authenticated.GET("/events/:eventId", getEvent)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
	authenticated.GET("/customers", getCustomers)
}

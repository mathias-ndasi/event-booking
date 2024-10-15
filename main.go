package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"example.com/event-booking/routes"
	"example.com/event-booking/utils"
)

func main() {
	error := utils.LoadEnv()
	if error != nil {
		log.Fatalf("Error loading .env file: %v", error)
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run("localhost:3000")
}

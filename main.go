package main

import (
	"github.com/gin-gonic/gin"

	"example.com/event-booking/routes"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run("localhost:3000")
}

package main

import (
	"fmt"
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

	response := twoSum([]int{1, 2, 4, 5, 2}, 4)
	fmt.Println(response)

	gin.ForceConsoleColor()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run("localhost:3000")
}

func twoSum(numbers []int, target int) []int {
	numberHashTable := make(map[int]int, len(numbers))
	response := make([]int, 2)

	for index, number := range numbers {
		numberHashTable[number] = index
		numberDifference := target - number

		if target > number {
			continue
		}

		if _, ok := numberHashTable[numberDifference]; !ok {
			continue
		}

		response = append(response, index, numberHashTable[number])
	}

	return response
}

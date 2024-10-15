package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	currentWorkingDirectory, _ := os.Getwd()
	error := godotenv.Load(fmt.Sprintf("%s/.env", currentWorkingDirectory))
	return error
}

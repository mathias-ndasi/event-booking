package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type environment struct {
	DATABASE_URL string
	GO_ENV       string
	APP_SECRET   string
}

func getEnvironmentVariable(key string) string {
	cwd, _ := os.Getwd()
	viper.SetConfigFile(fmt.Sprintf("%v/.env", cwd))

	error := viper.ReadInConfig()
	if error != nil {
		log.Fatalf("Error while reading config file %s", error)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion for environment variable: %v", key)
	}

	return value
}

func Environment() *environment {
	environment := &environment{
		DATABASE_URL: getEnvironmentVariable("DATABASE_URL"),
		GO_ENV:       getEnvironmentVariable("GO_ENV"),
		APP_SECRET:   getEnvironmentVariable("APP_SECRET"),
	}

	return environment
}

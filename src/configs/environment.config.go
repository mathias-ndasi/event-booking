package configs

import (
	"os"
)

type environment struct {
	DATABASE_URL string
	GO_ENV       string
	APP_SECRET   string
}

func Environment() *environment {
	return &environment{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		GO_ENV:       os.Getenv("GO_ENV"),
		APP_SECRET:   os.Getenv("APP_SECRET"),
	}
}

package prisma

import (
	goContext "context"
	"fmt"

	"example.com/event-booking/prisma/db"
	"example.com/event-booking/src/configs"
	"example.com/event-booking/src/constants"
)

func GetClient() (*db.PrismaClient, goContext.Context) {
	goEnvironment := configs.Environment().GO_ENV
	client := &db.PrismaClient{}

	switch goEnvironment {
	case constants.Environment()["test"]:
		client = db.NewClient(db.WithDatasourceURL(fmt.Sprintf("%v_test", configs.Environment().DATABASE_URL)))
	default:
		client = db.NewClient()
	}

	if error := client.Prisma.Connect(); error != nil {
		panic(error)
	}

	// defer func() {
	// 	if error := client.Prisma.Disconnect(); error != nil {
	// 		panic(error)
	// 	}
	// }()

	context := goContext.Background()

	return client, context
}

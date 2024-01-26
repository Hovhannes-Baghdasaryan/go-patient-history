package configuration

import (
	"context"
	constant "crud-go-api/libs/common/constant/logger"
	logger "crud-go-api/libs/common/logger/main"
	"crud-go-api/libs/data-layer/entity/tags/ent"
	"crud-go-api/libs/data-layer/entity/tags/ent/migrate"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

func DatabaseConnection() *ent.Client {
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
	)
	portParse, err := strconv.Atoi(port)

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, portParse, user, dbName, password))

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.DatabaseConfiguration, Message: fmt.Sprintf("Failed Opening Connection: %#v", err.Error())})
		client.Close()
	}

	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.DatabaseConfiguration, Message: fmt.Sprintf("Failed Creating Schema Resources: %v", err)})
		client.Close()
	}

	return client
}

package config

import (
	"fmt"
	config "github.com/Hovhannes-Baghdasaryan/go-patient-history/config/main"
	constant "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/constant/logger"
	logger "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/logger/main"
	"github.com/joho/godotenv"
	"os"
)

type DBMainConfig struct {
	Host     string
	Password string
	User     string
	DBName   string
	SslMode  string
	Port     string
}

type DBConfig struct {
	DBMainConfig
	config.MainConfig
}

func DBConfigLoad() *DBConfig {
	err := godotenv.Load("local.env")
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.DbConfig, Message: fmt.Sprintf("Some error occured. Err: %s", err)})
		return nil
	}

	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
		sslMode  = os.Getenv("SSL_MODE")
	)

	var cfg DBConfig

	cfg.DBName = dbName
	cfg.SslMode = sslMode
	cfg.Host = host
	cfg.Port = port
	cfg.Password = password
	cfg.User = user

	return &cfg

}

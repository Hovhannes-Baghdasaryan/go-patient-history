package logger

import (
	"crud-go-api/config"
	"crud-go-api/libs/common/constant/environment"
	"crud-go-api/libs/common/logger/slogpretty"
	"log/slog"
	"os"
)

type LoggerPayload struct {
	FuncName string
	Message  string
}

func SetupLogger() *slog.Logger {
	cfg := config.ConfigLoad()

	var log *slog.Logger

	switch cfg.Env {
	case environment.EnvLocal:
		log = slogpretty.SetupPrettySlog()
	case environment.EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case environment.EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func LogError(loggerPayload LoggerPayload) {
	log := SetupLogger()
	log.Error("", slog.String("Function Name", loggerPayload.FuncName), slog.String("Message", loggerPayload.Message))
}

func LogInfo(loggerPayload LoggerPayload) {
	log := SetupLogger()
	log.Info("", slog.String("Function Name", loggerPayload.FuncName), slog.String("Message", loggerPayload.Message))
}

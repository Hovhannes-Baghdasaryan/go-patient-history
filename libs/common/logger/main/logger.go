package logger

import (
	"go-patient-history/config"
	"go-patient-history/libs/common/constant/environment"
	"go-patient-history/libs/common/logger/slogpretty"
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

func LogDebug(loggerPayload LoggerPayload) {
	log := SetupLogger()
	log.Debug("", slog.String("Function Name", loggerPayload.FuncName), slog.String("Message", loggerPayload.Message))
}

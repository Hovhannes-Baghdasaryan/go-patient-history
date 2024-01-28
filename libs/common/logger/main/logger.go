package logger

import (
	config "github.com/Hovhannes-Baghdasaryan/go-patient-history/config/main"
	envconstant "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/constant/environment"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/logger/slogpretty"
	"log/slog"
	"os"
)

type LoggerPayload struct {
	FuncName string
	Message  string
}

func SetupLogger() *slog.Logger {
	cfg := config.MainConfigLoad()

	var log *slog.Logger

	switch cfg.Env {
	case envconstant.EnvLocal:
		log = slogpretty.SetupPrettySlog()
	case envconstant.EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envconstant.EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func LogError(loggerPayload LoggerPayload) {
	log := SetupLogger()
	log.Error("", slog.String("Function Name", loggerPayload.FuncName), slog.String("Message", loggerPayload.Message))
}

func LogDebug(loggerPayload LoggerPayload) {
	log := SetupLogger()
	log.Debug("", slog.String("Function Name", loggerPayload.FuncName), slog.String("Message", loggerPayload.Message))
}

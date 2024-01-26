package main

import (
	"fmt"
	"go-patient-history/config"
	router "go-patient-history/internal/router/patient"
	constant "go-patient-history/libs/common/constant/logger"
	logger "go-patient-history/libs/common/logger/main"
	common "go-patient-history/libs/common/router"
	configuration "go-patient-history/libs/data-layer/configuration/ent"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.ConfigLoad()

	logger.LogDebug(logger.LoggerPayload{Message: slog.String("env", cfg.Env).String(), FuncName: constant.MainBoostrap})

	clientDB := configuration.DatabaseConnection()

	baseRoute := common.BaseRouter()
	routes := router.InjectPatientRouter(baseRoute, clientDB)

	logger.LogDebug(logger.LoggerPayload{FuncName: constant.MainBoostrap, Message: fmt.Sprintf("Starting Server %s", slog.String("address", cfg.Address))})

	server := &http.Server{
		Addr:    cfg.Address,
		Handler: *routes,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.MainBoostrap, Message: fmt.Sprintf("Starting Server %#v", err)})
	}
}

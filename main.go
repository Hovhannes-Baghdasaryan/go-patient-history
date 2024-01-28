package main

import (
	"fmt"
	mainconfig "github.com/Hovhannes-Baghdasaryan/go-patient-history/config/main"
	_ "github.com/Hovhannes-Baghdasaryan/go-patient-history/docs"
	router "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/router/patient"
	constant "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/constant/logger"
	logger "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/logger/main"
	common "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/router"
	configuration "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/data-layer/config/ent"
	"log/slog"
	"net/http"
)

// @title 				Patient Service API
// @version				1.0
// @description 		Patient service API we predicting patient data by name

// @host 				localhost:8080
// @BasePath			/v1
func main() {
	cfg := mainconfig.MainConfigLoad()

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

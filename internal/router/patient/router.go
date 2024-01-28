package router

import (
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent"
	controller "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/controller/patients"
	services "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/services/patients"
	repository "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/data-layer/repository/patients"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InjectPatientRouter(router *gin.Engine, clientDB *ent.Client) **gin.Engine {
	// Repo
	patientRepository := repository.InjectPatientsRepositoryImpl(clientDB)

	// Service
	validate := validator.New()
	patientService := services.InjectPatientsServiceImpl(patientRepository, validate)

	// Controller
	patientController := controller.InjectPatientsController(patientService)

	patientsRouter := router.Group("v1/patients")
	{
		patientsRouter.GET("", patientController.FindAll)
		patientsRouter.GET("/:patientId", patientController.FindById)
		patientsRouter.POST("", patientController.Create)
		patientsRouter.PATCH("/:patientId", patientController.Update)
		patientsRouter.DELETE("/:patientId", patientController.Delete)
	}

	return &router
}

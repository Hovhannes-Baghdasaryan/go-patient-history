package controller

import (
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/config/main"
	reqconvert "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/converter/request"
	outputresponse "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/converter/response"
	services "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/services/patients"
	logconstant "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/constant/logger"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/exception"
	helper "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/helper/error"
	logger "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/logger/main"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PatientsControllerImpl struct {
	patientsService services.PatientsService
	Config          config.MainConfig
}

func InjectPatientsController(service services.PatientsService) PatientsController {
	cfg := config.MainConfigLoad()

	return &PatientsControllerImpl{
		patientsService: service,
		Config:          *cfg,
	}
}

// CreatePatient		godoc
// @Summary				Create Patient
// @Description			Save Patiens and predict data by name
// @Param				patient body reqconvert.CreatePatientRequest true "Create Patient"
// @Produce				application/json
// @Tags				Patients
// @Success				200 {object} response.Response[CreateSinglePatientOutputResponse]
// @Router				/patients [post]
func (controller PatientsControllerImpl) Create(ctx *gin.Context) {
	createPatientsRequest := reqconvert.CreatePatientRequest{}
	err := ctx.ShouldBindJSON(&createPatientsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsController, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return
	}

	patientCreateData, err := controller.patientsService.Create(ctx, createPatientsRequest)
	if err != nil {
		return
	}

	webResponse := response.Response[*outputresponse.PatientSingleResponse]{
		Status:  http.StatusCreated,
		Message: "Patient Created Successfully",
		Data:    patientCreateData,
	}
	webResponse.ActionSucceeded(ctx)
}

// UpdatePatient		godoc
// @Summary				Update Patient
// @Param				patientId path string true "update patient by id"
// @Param				patient body reqconvert.UpdatePatientRequest true "Create Patient"
// @Produce				application/json
// @Tags				Patients
// @Success				200 {object} response.Response[UpdateSinglePatientOutputResponse]
// @Router				/patients/{patientId} [patch]
func (controller PatientsControllerImpl) Update(ctx *gin.Context) {
	updatePatientRequest := reqconvert.UpdatePatientRequest{}
	err := ctx.ShouldBindJSON(&updatePatientRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsController, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return
	}

	patientId := ctx.Param("patientId")
	parsedUUID, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsController, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return
	}
	updatePatientRequest.Id = parsedUUID

	updatedPatient, err := controller.patientsService.Update(ctx, updatePatientRequest)
	if err != nil {
		return
	}

	webResponse := response.Response[UpdateSinglePatientOutputResponse]{
		Message: "Updated Successfully",
		Data:    updatedPatient,
	}
	webResponse.ActionSucceeded(ctx)
}

// DeletePatient		godoc
// @Summary				Delete Patient
// @Param				patientId path string true "delete patient by id"
// @Produce				application/json
// @Tags				Patients
// @Success				200 {object} response.Response[DeleteSinglePatientOutputResponse]
// @Router				/patients/{patientId} [delete]
func (controller PatientsControllerImpl) Delete(ctx *gin.Context) {
	patientId := ctx.Param("patientId")
	uuidParse, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DeletePatientsController, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return
	}

	resData, err := controller.patientsService.Delete(ctx, uuidParse)
	if err != nil {
		return
	}

	webResponse := response.Response[DeleteSinglePatientOutputResponse]{
		Message: "Deleted successfully",
		Data:    resData,
	}
	webResponse.ActionSucceeded(ctx)
}

// GetPatient			godoc
// @Summary				Get Patient
// @Param				patientId path string true "get patient by id"
// @Produce				application/json
// @Tags				Patients
// @Success				200 {object} response.Response[GetSinglePatientOutputResponse]
// @Router				/patients/{patientId} [get]
func (controller PatientsControllerImpl) FindById(ctx *gin.Context) {
	patientId := ctx.Param("patientId")

	parsedUUID, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindByIdPatientsController, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return
	}

	tagResponse, err := controller.patientsService.FindById(ctx, parsedUUID)
	if err != nil {
		return
	}

	webResponse := response.Response[GetSinglePatientOutputResponse]{
		Message: "Patient Single Find",
		Data:    tagResponse,
	}
	webResponse.ActionSucceeded(ctx)
}

// GetAllPatient		godoc
// @Summary				Get All Patients
// @Produce				application/json
// @Param 				name query string false "Name filter"
// @Param 				surname query string false "Surname filter"
// @Param 				patronymic query string false "Patronymic filter"
// @Tags				Patients
// @Success				200 {object} response.Response[GetAllPatientsOutputResponse]
// @Router				/patients [get]
func (controller PatientsControllerImpl) FindAll(ctx *gin.Context) {
	page := ctx.Query("page")
	perPage := ctx.Query("perPage")
	nameFilter := ctx.Query("name")
	surnameFilter := ctx.Query("surname")
	patronymicFilter := ctx.Query("patronymic")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	perPageInt, err := strconv.Atoi(perPage)
	if err != nil {
		perPageInt = 10
	}

	patientResponseAll, err := controller.patientsService.FindAll(ctx, pageInt, perPageInt, nameFilter, surnameFilter, patronymicFilter)
	if err != nil {
		return
	}

	webResponse := response.Response[GetAllPatientsOutputResponse]{
		Message: "Patients All Find",
		Data:    patientResponseAll,
	}
	webResponse.ActionSucceeded(ctx)
}

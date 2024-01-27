package controller

import (
	"github.com/gin-gonic/gin"
	"go-patient-history/config/main"
	reqconvert "go-patient-history/internal/converter/request"
	services "go-patient-history/internal/services/patients"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	helper "go-patient-history/libs/common/helper/error"
	logger "go-patient-history/libs/common/logger/main"
	"go-patient-history/libs/common/response"
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

	tagCreateData, err := controller.patientsService.Create(ctx, createPatientsRequest)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Status:  http.StatusCreated,
		Message: "Patient Created Successfully",
		Data:    tagCreateData,
	}
	webResponse.ActionSucceeded(ctx)
}

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

	updatedTagId, err := controller.patientsService.Update(ctx, updatePatientRequest)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Message: "Updated Successfully",
		Data:    updatedTagId,
	}
	webResponse.ActionSucceeded(ctx)
}

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

	webResponse := response.Response{
		Message: "Deleted successfully",
		Data:    resData,
	}
	webResponse.ActionSucceeded(ctx)
}

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

	webResponse := response.Response{
		Message: "Patient Single Find",
		Data:    tagResponse,
	}
	webResponse.ActionSucceeded(ctx)
}

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

	webResponse := response.Response{
		Message: "Patients All Find",
		Data:    patientResponseAll,
	}
	webResponse.ActionSucceeded(ctx)
}

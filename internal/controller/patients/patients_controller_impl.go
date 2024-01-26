package controller

import (
	"github.com/gin-gonic/gin"
	"go-patient-history/config"
	services "go-patient-history/internal/services/patients"
	"go-patient-history/libs/common/response"
	"net/http"
	"strconv"
)

type PatientsControllerImpl struct {
	patientsService services.PatientsService
	Config          config.Config
}

func InjectPatientsController(service services.PatientsService) PatientsController {
	cfg := config.ConfigLoad()

	return &PatientsControllerImpl{
		patientsService: service,
		Config:          *cfg,
	}
}

func (controller PatientsControllerImpl) Create(ctx *gin.Context) {
	tagCreateData, err := controller.patientsService.Create(ctx)
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
	updatedTagId, err := controller.patientsService.Update(ctx)
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
	resData, err := controller.patientsService.Delete(ctx)
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
	tagResponse, err := controller.patientsService.FindById(ctx)
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

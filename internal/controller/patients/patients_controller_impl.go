package controller

import (
	"github.com/gin-gonic/gin"
	services "go-patient-history/internal/services/patients"
	"go-patient-history/libs/common/response"
	"net/http"
)

type PatientsControllerImpl struct {
	patientsService services.PatientsService
}

func InjectTagsController(service services.PatientsService) PatientController {
	return &PatientsControllerImpl{
		patientsService: service,
	}
}

func (controller PatientsControllerImpl) Create(ctx *gin.Context) {
	tagCreateData, err := controller.patientsService.Create(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Status:  http.StatusCreated,
		Message: "Tag Created Successfully",
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
		Message: "Tag Single Find",
		Data:    tagResponse,
	}
	webResponse.ActionSucceeded(ctx)
}

func (controller PatientsControllerImpl) FindAll(ctx *gin.Context) {
	tagResponseAll, err := controller.patientsService.FindAll(ctx)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Message: "Patients All Find",
		Data:    tagResponseAll,
	}
	webResponse.ActionSucceeded(ctx)
}

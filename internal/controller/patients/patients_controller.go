package controller

import (
	"github.com/gin-gonic/gin"
	outputresponse "go-patient-history/internal/converter/response"
	"go-patient-history/libs/common/response"
)

type PatientsController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

// For swagger generics
type GetAllPatientsOutputResponse *response.PaginatedOutputResponse[[]outputresponse.PatientOutputResponse]
type GetSinglePatientOutputResponse *outputresponse.PatientSingleResponse
type DeleteSinglePatientOutputResponse *outputresponse.PatientOutputResponse
type UpdateSinglePatientOutputResponse *outputresponse.PatientSingleResponse
type CreateSinglePatientOutputResponse *outputresponse.PatientSingleResponse

package services

import (
	"github.com/gin-gonic/gin"
	resconverter "go-patient-history/internal/converter/response"
	repository "go-patient-history/libs/data-layer/repository/patients"
)

type PatientsService interface {
	Create(ctx *gin.Context) (resconverter.PatientSingleResponse, error)
	Update(ctx *gin.Context) (resconverter.PatientOutputResponse, error)
	Delete(ctx *gin.Context) (resconverter.PatientOutputResponse, error)
	FindById(ctx *gin.Context) (resconverter.PatientSingleResponse, error)
	FindAll(ctx *gin.Context, page int, perPage int, nameFilter string, surnameFilter string, patronymicFilter string) (resconverter.PatientPaginatedOutputResponse[[]resconverter.PatientOutputResponse], error)
	getPatientProviderData(ctx *gin.Context, patientPredictName string) (repository.PatientProviderData, error)
}

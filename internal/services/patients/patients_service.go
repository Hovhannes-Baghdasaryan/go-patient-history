package services

import (
	converter "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/converter/request"
	resconverter "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/converter/response"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/response"
	repository "github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/data-layer/repository/patients"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientsService interface {
	Create(ctx *gin.Context, createPatientRequest converter.CreatePatientRequest) (*resconverter.PatientSingleResponse, error)
	Update(ctx *gin.Context, updatePatientRequest converter.UpdatePatientRequest) (*resconverter.PatientSingleResponse, error)
	Delete(ctx *gin.Context, parsedUUId uuid.UUID) (*resconverter.PatientOutputResponse, error)
	FindById(ctx *gin.Context, parsedUUId uuid.UUID) (*resconverter.PatientSingleResponse, error)
	FindAll(ctx *gin.Context, page int, perPage int, nameFilter string, surnameFilter string, patronymicFilter string) (*response.PaginatedOutputResponse[[]resconverter.PatientOutputResponse], error)
	getPatientProviderData(patientPredictName string) (*repository.PatientProviderData, error)
}

package repository

import (
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent"
	converter "github.com/Hovhannes-Baghdasaryan/go-patient-history/internal/converter/request"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientsRepository interface {
	Save(createPatient converter.CreatePatientRequest, patientProviderData PatientProviderData, ctx *gin.Context) (*ent.PatientEntity, error)
	Update(ctx *gin.Context, updatePayload converter.UpdatePatientRequest) (*ent.PatientEntity, error)
	UpdateWithProviderData(ctx *gin.Context, updatePayload converter.UpdatePatientRequest, providerData PatientProviderData) (*ent.PatientEntity, error)
	Delete(patientId uuid.UUID, ctx *gin.Context) error
	FindById(patientId uuid.UUID, ctx *gin.Context) (*ent.PatientEntity, error)
	FindAll(ctx *gin.Context, page int, perPage int, nameFilter string, surnameFilter string, patronymicFilter string) (response.PaginatedOutputResponse[[]*ent.PatientEntity], error)
}

type PatientProviderData struct {
	Age     *int    `json:"age,omitempty"`
	Country string  `json:"country,omitempty"`
	Gender  *string `json:"gender,omitempty"`
}

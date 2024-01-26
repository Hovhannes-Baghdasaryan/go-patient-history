package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-patient-history/ent"
	converter "go-patient-history/internal/converter/request"
)

type PatientsRepository interface {
	Save(tags converter.CreatePatientRequest, ctx *gin.Context) (*ent.PatientEntity, error)
	Update(updatePayload converter.UpdatePatientRequest, ctx *gin.Context) (uuid.UUID, error)
	Delete(tagId uuid.UUID, ctx *gin.Context) (uuid.UUID, error)
	FindById(tagId uuid.UUID, ctx *gin.Context) (*ent.PatientEntity, error)
	FindAll() ([]*ent.PatientEntity, error)
}

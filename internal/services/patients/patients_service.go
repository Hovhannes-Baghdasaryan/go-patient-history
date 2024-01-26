package services

import (
	"github.com/gin-gonic/gin"
	converter "go-patient-history/internal/converter/response"
)

type TagsService interface {
	Create(ctx *gin.Context) (converter.PatientOutputResponse, error)
	Update(ctx *gin.Context) (converter.PatientOutputResponse, error)
	Delete(ctx *gin.Context) (converter.PatientOutputResponse, error)
	FindById(ctx *gin.Context) (converter.PatientResponse, error)
	FindAll(ctx *gin.Context) ([]converter.PatientResponse, error)
}

package response

import (
	"github.com/google/uuid"
	"go-patient-history/ent"
	"go-patient-history/libs/common/response"
)

type PatientOutputResponse struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic,omitempty"`
}

type PatientSingleResponse struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic,omitempty"`
	Age        int       `json:"age"`
	Country    string    `json:"country"`
	Gender     string    `json:"gender"`
}

type PatientFindPagination interface {
	[]PatientOutputResponse | []*ent.PatientEntity
}

type PatientPaginatedOutputResponse[T PatientFindPagination] struct {
	Patients   T                   `json:"patients"`
	Pagination response.Pagination `json:"pagination"`
}

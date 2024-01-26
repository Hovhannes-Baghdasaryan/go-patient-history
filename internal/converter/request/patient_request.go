package converter

import (
	"github.com/google/uuid"
)

type CreatePatientRequest struct {
	Name       string `validate:"required,min=2,max=15" json:"name"`
	Surname    string `validate:"required,min=2,max=15" json:"surname"`
	Patronymic string `validate:"omitempty,min=2,max=15" json:"patronymic"`
}

type UpdatePatientRequest struct {
	Id         uuid.UUID `json:"id,omitempty"`
	Name       *string   `validate:"max=200,min=2" json:"name,omitempty"`
	Surname    *string   `validate:"max=200,min=2" json:"surname,omitempty"`
	Patronymic *string   `validate:"min=2,max=15" json:"patronymic,omitempty"`
}

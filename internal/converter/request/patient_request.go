package converter

import "github.com/google/uuid"

type CreatePatientRequest struct {
	Name       string `validate:"required,min=1,max=15" json:"name"`
	Surname    string `validate:"required,min=1,max=15" json:"surname"`
	Patronymic string `validate:"required,min=1,max=15" json:"patronymic"`
}

type UpdatePatientRequest struct {
	Id   uuid.UUID `validate:"required"`
	Name string    `validate:"required,max=200,min=1" json:"name"`
}

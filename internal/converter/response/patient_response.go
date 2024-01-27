package response

import (
	"github.com/google/uuid"
)

type PatientOutputResponse struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name,omitempty"`
	Surname    string    `json:"surname,omitempty"`
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

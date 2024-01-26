package response

import "github.com/google/uuid"

type PatientResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
}

type PatientOutputResponse struct {
	Id uuid.UUID `json:"id"`
}

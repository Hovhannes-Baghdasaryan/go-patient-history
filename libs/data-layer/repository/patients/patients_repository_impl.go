package repository

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-patient-history/ent"
	"go-patient-history/ent/patiententity"
	"go-patient-history/ent/predicate"
	converter "go-patient-history/internal/converter/request"
	errconstant "go-patient-history/libs/common/constant/error"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	logger "go-patient-history/libs/common/logger/main"
	"go-patient-history/libs/common/response"
)

type PatientsRepositoryImpl struct {
	clientDB *ent.Client
}

func InjectPatientsRepositoryImpl(clientDB *ent.Client) *PatientsRepositoryImpl {
	return &PatientsRepositoryImpl{
		clientDB: clientDB,
	}
}

func (repo *PatientsRepositoryImpl) Save(createPatientPayload converter.CreatePatientRequest, providerData PatientProviderData, ctx *gin.Context) (*ent.PatientEntity, error) {
	result, err := repo.clientDB.PatientEntity.Create().SetName(createPatientPayload.Name).SetSurname(createPatientPayload.Surname).SetPatronymic(createPatientPayload.Patronymic).SetAge(*providerData.Age).SetGender(patiententity.Gender(*providerData.Gender)).SetCountry(providerData.Country).Save(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsRepository, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return nil, err
	}

	return result, nil
}

func (repo *PatientsRepositoryImpl) UpdateWithProviderData(ctx *gin.Context, updatePayload converter.UpdatePatientRequest, providerData PatientProviderData) (*ent.PatientEntity, error) {
	responsePatient, err := repo.clientDB.PatientEntity.UpdateOneID(updatePayload.Id).SetNillableName(updatePayload.Name).SetNillableSurname(updatePayload.Surname).SetNillablePatronymic(updatePayload.Patronymic).SetAge(*providerData.Age).SetCountry(providerData.Country).SetGender(patiententity.Gender(*providerData.Gender)).Save(context.Background())
	if responsePatient == nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsProviderRepository, Message: errconstant.PatientNotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.PatientNotFound).Error(),
		}
		webError.BadRequestException(ctx)
		return nil, errors.New(errconstant.PatientNotFound)
	}

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsProviderRepository, Message: errconstant.DBInternalError})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return nil, errors.New(errconstant.DBInternalError)
	}

	return responsePatient, nil
}

func (repo *PatientsRepositoryImpl) Update(ctx *gin.Context, updatePayload converter.UpdatePatientRequest) (*ent.PatientEntity, error) {
	responsePatient, err := repo.clientDB.PatientEntity.UpdateOneID(updatePayload.Id).SetNillableName(updatePayload.Name).SetNillableSurname(updatePayload.Surname).SetNillablePatronymic(updatePayload.Patronymic).Where(patiententity.ID(updatePayload.Id)).Save(context.Background())
	if responsePatient == nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsRepository, Message: errconstant.PatientNotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.PatientNotFound).Error(),
		}
		webError.BadRequestException(ctx)
		return nil, errors.New(errconstant.PatientNotFound)
	}

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsRepository, Message: errconstant.DBInternalError})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return nil, errors.New(errconstant.DBInternalError)
	}

	return responsePatient, nil
}

func (repo *PatientsRepositoryImpl) Delete(patientUUId uuid.UUID, ctx *gin.Context) error {
	err := repo.clientDB.PatientEntity.DeleteOneID(patientUUId).Exec(context.Background())

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DeletePatientsRepository, Message: errconstant.PatientNotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.PatientNotFound).Error(),
		}
		webError.InternalException(ctx)
		return errors.New(errconstant.PatientNotFound)
	}

	return nil
}

func (repo *PatientsRepositoryImpl) FindById(PatientId uuid.UUID, ctx *gin.Context) (*ent.PatientEntity, error) {
	result, err := repo.clientDB.PatientEntity.Query().Where(patiententity.ID(PatientId)).Only(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindByIdPatientsRepository, Message: errconstant.PatientNotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.PatientNotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return nil, err
	}

	return result, nil
}

func (repo *PatientsRepositoryImpl) FindAll(ctx *gin.Context, page int, perPage int, nameFilter string, surnameFilter string, patronymicFilter string) (response.PaginatedOutputResponse[[]*ent.PatientEntity], error) {
	skipPages := (page - 1) * perPage

	whereOption := []predicate.PatientEntity{patiententity.NameContainsFold(nameFilter), patiententity.SurnameContainsFold(surnameFilter), patiententity.PatronymicContainsFold(patronymicFilter)}

	// get all count for total items
	patientFullPagesCount, err := repo.clientDB.PatientEntity.Query().Where(whereOption...).Count(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindAllPatientsRepository, Message: err.Error()})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return response.PaginatedOutputResponse[[]*ent.PatientEntity]{}, errors.New(errconstant.DBInternalError)
	}

	// (Field)ContainsFold is for filtering not case-sensitive
	result, err := repo.clientDB.PatientEntity.Query().Offset(skipPages).Limit(perPage).Where(whereOption...).All(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindAllPatientsRepository, Message: err.Error()})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return response.PaginatedOutputResponse[[]*ent.PatientEntity]{}, errors.New(errconstant.DBInternalError)
	}

	return response.PaginatedOutputResponse[[]*ent.PatientEntity]{
		Pagination: response.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   patientFullPagesCount,
		},
		Items: result,
	}, nil
}

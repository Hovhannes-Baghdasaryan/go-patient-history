package repository

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-patient-history/ent"
	"go-patient-history/ent/patiententity"
	converter "go-patient-history/internal/converter/request"
	errconstant "go-patient-history/libs/common/constant/error"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	logger "go-patient-history/libs/common/logger/main"
)

type PatientsRepositoryImpl struct {
	clientDB *ent.Client
}

func InjectPatientsRepositoryImpl(clientDB *ent.Client) *PatientsRepositoryImpl {
	return &PatientsRepositoryImpl{
		clientDB: clientDB,
	}
}

func (repo *PatientsRepositoryImpl) Save(Patient converter.CreatePatientRequest, ctx *gin.Context) (*ent.PatientEntity, error) {
	result, err := repo.clientDB.PatientEntity.Create().SetName(Patient.Name).SetSurname(Patient.Surname).SetPatronymic(Patient.Patronymic).Save(context.Background())
	if err != nil {
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return &ent.PatientEntity{}, err
	}

	return result, nil
}

func (repo *PatientsRepositoryImpl) Update(updatePayload converter.UpdatePatientRequest, ctx *gin.Context) (uuid.UUID, error) {
	ok, err := repo.clientDB.PatientEntity.Update().SetName(updatePayload.Name).Where(patiententity.ID(updatePayload.Id)).Save(context.Background())

	if ok == 0 {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsRepository, Message: errconstant.NotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.NotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return uuid.UUID{}, errors.New(errconstant.NotFound)
	}

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsRepository, Message: errconstant.DBInternalError})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return uuid.UUID{}, errors.New(errconstant.DBInternalError)
	}

	return updatePayload.Id, nil
}

func (repo *PatientsRepositoryImpl) Delete(PatientUUId uuid.UUID, ctx *gin.Context) (uuid.UUID, error) {
	ok, err := repo.clientDB.PatientEntity.Delete().Where(patiententity.ID(PatientUUId)).Exec(context.Background())
	if ok == 0 {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DeletePatientsRepository, Message: errconstant.NotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.NotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return uuid.UUID{}, errors.New(errconstant.NotFound)
	}

	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DeletePatientsRepository, Message: errconstant.DBInternalError})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return uuid.UUID{}, errors.New(errconstant.DBInternalError)
	}

	return PatientUUId, nil
}

func (repo *PatientsRepositoryImpl) FindById(PatientId uuid.UUID, ctx *gin.Context) (*ent.PatientEntity, error) {
	result, err := repo.clientDB.PatientEntity.Query().Where(patiententity.ID(PatientId)).Only(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindByIdPatientsRepository, Message: errconstant.NotFound})
		webError := exception.Error{
			Message: errors.New(errconstant.NotFound).Error(),
		}
		webError.NotFoundException(ctx)
		return nil, err
	}

	return result, nil
}

func (repo *PatientsRepositoryImpl) FindAll() ([]*ent.PatientEntity, error) {
	result, err := repo.clientDB.PatientEntity.Query().All(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}

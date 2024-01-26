package repository

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go-patient-history/ent"
	"go-patient-history/ent/predicate"
	errconstant "go-patient-history/libs/common/constant/error"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	logger "go-patient-history/libs/common/logger/main"
	"go-patient-history/libs/common/response"
)

// 1. Entity model

type AllEntities interface {
	[]*ent.PatientEntity
}

// 2. For query limit and offset client methods

type AllEntitiesClient interface {
	ent.PatientEntityClient
}

// 3. For where option

type AllPredicts interface {
	[]predicate.PatientEntity
}

type BaseRepositoryImpl[T1 AllEntities, T2 AllEntitiesClient, T3 AllPredicts] struct {
	Repo   BaseRepository
	Entity T2
}

func (repo BaseRepositoryImpl[T1, T2, T3]) FindPaginated(ctx *gin.Context, entityClient ent.PatientEntityClient, page int, perPage int, whereOption T3) (PaginatedOutputResponse[T1], error) {
	skipPages := (page - 1) * perPage

	// get all count for total items
	patientFullPagesCount, err := entityClient.Query().Count(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindAllPatientsRepository, Message: err.Error()})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return PaginatedOutputResponse[T1]{}, errors.New(errconstant.DBInternalError)
	}

	// (Field)ContainsFold is for filtering not case-sensitive
	result, err := entityClient.Query().Offset(skipPages).Limit(perPage).Where(whereOption...).All(context.Background())
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindAllPatientsRepository, Message: err.Error()})
		webError := exception.Error{
			Message: errors.New(errconstant.DBInternalError).Error(),
		}
		webError.InternalException(ctx)
		return PaginatedOutputResponse[T1]{}, errors.New(errconstant.DBInternalError)
	}

	return PaginatedOutputResponse[T1]{
		Pagination: response.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   patientFullPagesCount,
		},
		Items: result,
	}, nil
}

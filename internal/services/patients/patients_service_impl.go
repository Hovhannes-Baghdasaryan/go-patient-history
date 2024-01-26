package services

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	reqconvert "go-patient-history/internal/converter/request"
	resconverter "go-patient-history/internal/converter/response"
	errconstant "go-patient-history/libs/common/constant/error"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	helper "go-patient-history/libs/common/helper/error"
	logger "go-patient-history/libs/common/logger/main"
	repository "go-patient-history/libs/data-layer/repository/tags"
)

type PatientsServiceImpl struct {
	PatientsRepository repository.PatientsRepository
	Validate           *validator.Validate
}

func InjectPatientsServiceImpl(patientRepository repository.PatientsRepository, validate *validator.Validate) TagsService {
	return &PatientsServiceImpl{
		PatientsRepository: patientRepository,
		Validate:           validate,
	}
}

func (service PatientsServiceImpl) Create(ctx *gin.Context) (resconverter.PatientOutputResponse, error) {
	createTagsRequest := reqconvert.CreatePatientRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	err = service.Validate.Struct(createTagsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	patientCreateData, err := service.PatientsRepository.Save(createTagsRequest, ctx)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsService, Message: err.Error()})
		return resconverter.PatientOutputResponse{}, err
	}

	return resconverter.PatientOutputResponse{Id: patientCreateData.ID}, nil
}

func (service PatientsServiceImpl) Update(ctx *gin.Context) (resconverter.PatientOutputResponse, error) {
	updateTagsRequest := reqconvert.UpdatePatientRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	tagId := ctx.Param("tagId")
	parsedUUID, err := helper.IsValidUUID(tagId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	updateTagsRequest.Id = parsedUUID

	result, err := service.PatientsRepository.Update(updateTagsRequest, ctx)
	if err != nil {
		return resconverter.PatientOutputResponse{}, err
	}

	return resconverter.PatientOutputResponse{Id: result}, nil
}

func (service PatientsServiceImpl) Delete(ctx *gin.Context) (resconverter.PatientOutputResponse, error) {
	tagId := ctx.Param("tagId")
	uuidParse, err := helper.IsValidUUID(tagId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DeletePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	deletedTagId, err := service.PatientsRepository.Delete(uuidParse, ctx)
	if err != nil {
		return resconverter.PatientOutputResponse{}, err
	}

	return resconverter.PatientOutputResponse{Id: deletedTagId}, nil
}

func (service PatientsServiceImpl) FindById(ctx *gin.Context) (resconverter.PatientResponse, error) {
	patientId := ctx.Param("tagId")

	parsedUUID, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindByIdPatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientResponse{}, err
	}

	patientData, err := service.PatientsRepository.FindById(parsedUUID, ctx)
	if err != nil {
		return resconverter.PatientResponse{}, err
	}

	return resconverter.PatientResponse{
		Id:   patientData.ID,
		Name: patientData.Name,
	}, nil
}

func (service PatientsServiceImpl) FindAll(ctx *gin.Context) ([]resconverter.PatientResponse, error) {
	result, err := service.PatientsRepository.FindAll()
	if err != nil {
		webError := exception.Error{
			Message: errconstant.DBInternalError,
		}
		webError.InternalException(ctx)
		return nil, err
	}

	var patients []resconverter.PatientResponse
	for _, value := range result {
		patient := resconverter.PatientResponse{
			Id:   value.ID,
			Name: value.Name,
		}

		patients = append(patients, patient)
	}
	return patients, nil
}

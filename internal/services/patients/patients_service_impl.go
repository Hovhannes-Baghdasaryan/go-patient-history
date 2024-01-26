package services

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-patient-history/ent"
	reqconvert "go-patient-history/internal/converter/request"
	resconverter "go-patient-history/internal/converter/response"
	providers "go-patient-history/internal/providers/patients"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	helper "go-patient-history/libs/common/helper/error"
	logger "go-patient-history/libs/common/logger/main"
	repository "go-patient-history/libs/data-layer/repository/patients"
)

type PatientsServiceImpl struct {
	PatientsRepository repository.PatientsRepository
	Validate           *validator.Validate
	PatientProvider    providers.PatientProviderImpl
}

func InjectPatientsServiceImpl(patientsRepository repository.PatientsRepository, validate *validator.Validate) PatientsService {
	return &PatientsServiceImpl{
		PatientsRepository: patientsRepository,
		Validate:           validate,
	}
}

func (service PatientsServiceImpl) Create(ctx *gin.Context) (resconverter.PatientSingleResponse, error) {
	createPatientsRequest := reqconvert.CreatePatientRequest{}
	err := ctx.ShouldBindJSON(&createPatientsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientSingleResponse{}, err
	}

	err = service.Validate.Struct(createPatientsRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientSingleResponse{}, err
	}

	providerData, err := service.getPatientProviderData(ctx, createPatientsRequest.Name)
	if err != nil {
		return resconverter.PatientSingleResponse{}, err
	}

	patientCreateData, err := service.PatientsRepository.Save(createPatientsRequest, providerData, ctx)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.CreatePatientsService, Message: err.Error()})
		return resconverter.PatientSingleResponse{}, err
	}

	return resconverter.PatientSingleResponse{Id: patientCreateData.ID, Name: patientCreateData.Name, Surname: patientCreateData.Surname, Patronymic: *patientCreateData.Patronymic, Gender: patientCreateData.Gender.String(), Age: patientCreateData.Age, Country: patientCreateData.Country}, nil
}

func (service PatientsServiceImpl) Update(ctx *gin.Context) (resconverter.PatientOutputResponse, error) {
	updatePatientRequest := reqconvert.UpdatePatientRequest{}
	err := ctx.ShouldBindJSON(&updatePatientRequest)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	patientId := ctx.Param("patientId")
	parsedUUID, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.UpdatePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}
	updatePatientRequest.Id = parsedUUID

	// If I update surname I do not need to update age,gender,country because they depend only on name
	var providerData repository.PatientProviderData
	var result *ent.PatientEntity
	if updatePatientRequest.Name != nil {
		providerData, err = service.getPatientProviderData(ctx, *updatePatientRequest.Name)
		if err != nil {
			return resconverter.PatientOutputResponse{}, err
		}
		repoResult, err := service.PatientsRepository.UpdateWithProviderData(ctx, updatePatientRequest, providerData)
		if err != nil {
			return resconverter.PatientOutputResponse{}, err
		}

		result = repoResult
	} else {
		repoResult, err := service.PatientsRepository.Update(ctx, updatePatientRequest)
		if err != nil {
			return resconverter.PatientOutputResponse{}, err
		}

		result = repoResult
	}

	return resconverter.PatientOutputResponse{Id: result.ID, Name: result.Name, Surname: result.Surname, Patronymic: *result.Patronymic}, nil
}

func (service PatientsServiceImpl) Delete(ctx *gin.Context) (resconverter.PatientOutputResponse, error) {
	patientId := ctx.Param("patientId")
	uuidParse, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.DeletePatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientOutputResponse{}, err
	}

	err = service.PatientsRepository.Delete(uuidParse, ctx)
	if err != nil {
		return resconverter.PatientOutputResponse{}, err
	}

	return resconverter.PatientOutputResponse{Id: uuidParse}, nil
}

func (service PatientsServiceImpl) FindById(ctx *gin.Context) (resconverter.PatientSingleResponse, error) {
	patientId := ctx.Param("patientId")

	parsedUUID, err := helper.IsValidUUID(patientId)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.FindByIdPatientsService, Message: err.Error()})
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return resconverter.PatientSingleResponse{}, err
	}

	patientData, err := service.PatientsRepository.FindById(parsedUUID, ctx)
	if err != nil {
		return resconverter.PatientSingleResponse{}, err
	}

	return resconverter.PatientSingleResponse{
		Id:         patientData.ID,
		Name:       patientData.Name,
		Surname:    patientData.Surname,
		Patronymic: *patientData.Patronymic,
		Gender:     patientData.Gender.String(),
		Country:    patientData.Country,
		Age:        patientData.Age,
	}, nil
}

func (service PatientsServiceImpl) FindAll(ctx *gin.Context, page int, perPage int, nameFilter string, surnameFilter string, patronymicFilter string) (resconverter.PatientPaginatedOutputResponse[[]resconverter.PatientOutputResponse], error) {
	result, err := service.PatientsRepository.FindAll(ctx, page, perPage, nameFilter, surnameFilter, patronymicFilter)
	if err != nil {
		return resconverter.PatientPaginatedOutputResponse[[]resconverter.PatientOutputResponse]{}, err
	}

	var patients []resconverter.PatientOutputResponse
	for _, value := range result.Items {
		newValue := resconverter.PatientOutputResponse{
			Id:         value.ID,
			Name:       value.Name,
			Surname:    value.Surname,
			Patronymic: *value.Patronymic,
		}
		patients = append(patients, newValue)
	}

	return resconverter.PatientPaginatedOutputResponse[[]resconverter.PatientOutputResponse]{
		Patients:   patients,
		Pagination: result.Pagination,
	}, nil
}

func (service PatientsServiceImpl) getPatientProviderData(ctx *gin.Context, patientPredictName string) (repository.PatientProviderData, error) {
	ageResponse, err := service.PatientProvider.GetAge(patientPredictName)
	if err != nil {
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return repository.PatientProviderData{}, err
	}

	genderResponse, err := service.PatientProvider.GetGender(patientPredictName)
	if err != nil {
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return repository.PatientProviderData{}, err
	}

	countryResponse, err := service.PatientProvider.GetCountry(patientPredictName)
	if err != nil {
		webError := exception.Error{
			Message: err.Error(),
		}
		webError.BadRequestException(ctx)
		return repository.PatientProviderData{}, err
	}

	return repository.PatientProviderData{
		Age:     ageResponse.Age,
		Country: countryResponse.Country,
		Gender:  genderResponse.Gender,
	}, nil
}

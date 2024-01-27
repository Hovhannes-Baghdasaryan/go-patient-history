package services

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go-patient-history/ent"
	reqconvert "go-patient-history/internal/converter/request"
	resconverter "go-patient-history/internal/converter/response"
	providers "go-patient-history/internal/providers/patients"
	logconstant "go-patient-history/libs/common/constant/logger"
	"go-patient-history/libs/common/exception"
	logger "go-patient-history/libs/common/logger/main"
	"go-patient-history/libs/common/response"
	repository "go-patient-history/libs/data-layer/repository/patients"
	"sync"
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

func (service PatientsServiceImpl) Create(ctx *gin.Context, createPatientsRequest reqconvert.CreatePatientRequest) (resconverter.PatientSingleResponse, error) {
	err := service.Validate.Struct(createPatientsRequest)
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

// If I update surname I do not need to update age,gender,country because they depend only on name change

func (service PatientsServiceImpl) Update(ctx *gin.Context, updatePatientRequest reqconvert.UpdatePatientRequest) (*resconverter.PatientSingleResponse, error) {
	var result *ent.PatientEntity
	if updatePatientRequest.Name != nil {
		providerData, err := service.getPatientProviderData(ctx, *updatePatientRequest.Name)
		if err != nil {
			return nil, err
		}

		repoResult, err := service.PatientsRepository.UpdateWithProviderData(ctx, updatePatientRequest, providerData)
		if err != nil {
			return nil, err
		}

		result = repoResult
	} else {
		repoResult, err := service.PatientsRepository.Update(ctx, updatePatientRequest)
		if err != nil {
			return nil, err
		}

		result = repoResult
	}

	return &resconverter.PatientSingleResponse{
		Id:         result.ID,
		Name:       result.Name,
		Surname:    result.Surname,
		Patronymic: *result.Patronymic,
		Gender:     result.Gender.String(),
		Age:        result.Age,
		Country:    result.Country,
	}, nil
}

func (service PatientsServiceImpl) Delete(ctx *gin.Context, parsedUUId uuid.UUID) (resconverter.PatientOutputResponse, error) {
	if err := service.PatientsRepository.Delete(parsedUUId, ctx); err != nil {
		return resconverter.PatientOutputResponse{}, err
	}

	return resconverter.PatientOutputResponse{Id: parsedUUId}, nil
}

func (service PatientsServiceImpl) FindById(ctx *gin.Context, parsedUUID uuid.UUID) (resconverter.PatientSingleResponse, error) {
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

func (service PatientsServiceImpl) FindAll(ctx *gin.Context, page int, perPage int, nameFilter string, surnameFilter string, patronymicFilter string) (response.PaginatedOutputResponse[[]resconverter.PatientOutputResponse], error) {
	result, err := service.PatientsRepository.FindAll(ctx, page, perPage, nameFilter, surnameFilter, patronymicFilter)
	if err != nil {
		return response.PaginatedOutputResponse[[]resconverter.PatientOutputResponse]{}, err
	}
	patients := make([]resconverter.PatientOutputResponse, 0)
	for _, value := range result.Items {
		newValue := resconverter.PatientOutputResponse{
			Id:         value.ID,
			Name:       value.Name,
			Surname:    value.Surname,
			Patronymic: *value.Patronymic,
		}
		patients = append(patients, newValue)
	}

	return response.PaginatedOutputResponse[[]resconverter.PatientOutputResponse]{
		Items:      patients,
		Pagination: result.Pagination,
	}, nil
}

func (service PatientsServiceImpl) getPatientProviderData(ctx *gin.Context, patientPredictName string) (repository.PatientProviderData, error) {
	var wg sync.WaitGroup

	fetchData := func(provider providers.PatientProviderImpl, fetchType string, result *repository.PatientProviderData) {
		defer wg.Done()

		switch fetchType {
		case providers.Age:
			ageResp, err := provider.GetAge(patientPredictName)
			result.Age = ageResp.Age
			if err != nil {
				return
			}
		case providers.Country:
			countryResp, err := provider.GetCountry(patientPredictName)
			result.Country = countryResp.Country
			if err != nil {
				return
			}
		case providers.Gender:
			genderResp, err := provider.GetGender(patientPredictName)
			result.Gender = genderResp.Gender
			if err != nil {
				return
			}
		}
	}

	var patientProviderResponse repository.PatientProviderData

	wg.Add(3)
	go fetchData(service.PatientProvider, providers.Age, &patientProviderResponse)
	go fetchData(service.PatientProvider, providers.Gender, &patientProviderResponse)
	go fetchData(service.PatientProvider, providers.Country, &patientProviderResponse)

	wg.Wait()

	return patientProviderResponse, nil
}

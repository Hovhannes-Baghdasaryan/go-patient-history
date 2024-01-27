package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	constant "go-patient-history/libs/common/constant/error"
	logconstant "go-patient-history/libs/common/constant/logger"
	logger "go-patient-history/libs/common/logger/main"
	"net/http"
)

type PatientProviderImpl struct {
	provider PatientsProvider
}

func (provider PatientProviderImpl) GetAge(name string) (GetAgeResponse, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {

		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientAgeProvider, Message: err.Error()})
		return GetAgeResponse{}, errors.New(constant.GetPatientAgeProviderRequest)
	}
	defer response.Body.Close()

	var finalResponse GetAgeResponse
	if err := json.NewDecoder(response.Body).Decode(&finalResponse); err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientAgeProvider, Message: err.Error()})
		return GetAgeResponse{}, errors.New(constant.GetPatientAgeProviderJsonDecoding)
	}

	if finalResponse.Age == nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientAgeProvider, Message: constant.AgeWithGivenNameNotExist})
		return GetAgeResponse{}, errors.New(constant.AgeWithGivenNameNotExist)
	}

	return GetAgeResponse{
		Age:   finalResponse.Age,
		Name:  finalResponse.Name,
		Count: finalResponse.Count,
	}, nil
}

func (provider PatientProviderImpl) GetGender(name string) (GetGenderResponse, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientGenderProvider, Message: err.Error()})
		return GetGenderResponse{}, errors.New(constant.GetPatientGenderProviderRequest)
	}
	defer response.Body.Close()

	var finalResponse GetGenderResponse
	if err := json.NewDecoder(response.Body).Decode(&finalResponse); err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientGenderProvider, Message: err.Error()})
		return GetGenderResponse{}, errors.New(constant.GetPatientGenderProviderJsonDecoding)
	}

	if finalResponse.Gender == nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientGenderProvider, Message: constant.GenderWithGivenNameNotExist})
		return GetGenderResponse{}, errors.New(constant.GenderWithGivenNameNotExist)
	}

	return GetGenderResponse{
		Name:        finalResponse.Name,
		Count:       finalResponse.Count,
		Gender:      finalResponse.Gender,
		Probability: finalResponse.Probability,
	}, nil
}

func (provider PatientProviderImpl) GetCountry(name string) (GetCountryResponse[string], error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientCountryProvider, Message: err.Error()})
		return GetCountryResponse[string]{}, errors.New(constant.GetPatientCountryProviderRequest)
	}
	defer response.Body.Close()

	var finalResponse GetCountryResponse[[]CountryItem]
	if err := json.NewDecoder(response.Body).Decode(&finalResponse); err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientCountryProvider, Message: err.Error()})
		return GetCountryResponse[string]{}, errors.New(constant.GetPatientCountryProviderJsonDecoding)
	}

	if len(finalResponse.Country) == 0 {
		logger.LogError(logger.LoggerPayload{FuncName: logconstant.GetPatientCountryProvider, Message: constant.CountryWithGivenNameNotExist})
		return GetCountryResponse[string]{}, errors.New(constant.CountryWithGivenNameNotExist)
	}

	var mostLikelyCountry string
	var mostLikelyCountryProbability float64
	for _, country := range finalResponse.Country {
		if country.Probability > mostLikelyCountryProbability {
			mostLikelyCountry = country.Country_id
			mostLikelyCountryProbability = country.Probability
		}
	}

	return GetCountryResponse[string]{
		Country: mostLikelyCountry,
		Count:   finalResponse.Count,
		Name:    finalResponse.Name,
	}, nil
}

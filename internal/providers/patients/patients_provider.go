package providers

type PatientsProvider interface {
	GetAge(name string) int
	GetGender(name string) string
	GetNationalization(name string) string
}

type GetAgeResponse struct {
	Count int    `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   *int   `json:"age,omitempty"`
}

type GetGenderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      *string `json:"gender,omitempty"`
	Probability float64 `json:"probability"`
}

type CountryGeneric interface {
	string | []CountryItem
}

type CountryItem struct {
	Country_id  string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type GetCountryResponse[T CountryGeneric] struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country T
}

package dateservice

import (
	"context"
	"fmt"

	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

type Service struct {
	dateapi *dateapi.APIClient
}

func New() *Service {
	config := dateapi.NewConfiguration()
	config.Scheme = "https"
	config.Host = "date.nager.at"
	apiClient := dateapi.NewAPIClient(config)
	return &Service{apiClient}
}

func (s *Service) GetCountries(ctx context.Context) ([]dateapi.CountryV3Dto, error) {
	resp, _, err := s.dateapi.CountryAPI.CountryAvailableCountries(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `CountryAPI.CountryAvailableCountries``: %v", err)
	}
	return resp, nil
}

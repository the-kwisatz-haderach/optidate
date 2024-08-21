package dateservice

import (
	"context"
	"fmt"

	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

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
		return nil, fmt.Errorf("error when calling `CountryAvailableCountries``: %v", err)
	}
	return resp, nil
}

func (s *Service) GetCalendar(ctx context.Context, countryCode string, opts CreateCalendarOptions) ([]FormattedDate, error) {
	resp, _, err := s.dateapi.PublicHolidayAPI.PublicHolidayPublicHolidaysV3(ctx, int32(opts.Year), countryCode).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `PublicHolidayNextPublicHolidays``: %v", err)
	}
	formatted, err := createCalendar(resp, opts)
	if err != nil {
		return nil, fmt.Errorf("error when formatting dates`: %v", err)
	}
	return formatted, nil
}

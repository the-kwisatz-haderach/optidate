package dateservice

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

func New(cache ServiceCache) *Service {
	config := dateapi.NewConfiguration()
	config.Scheme = "https"
	config.Host = "date.nager.at"
	apiClient := dateapi.NewAPIClient(config)
	return &Service{apiClient, cache}
}

func createCacheKey(year int32, countryCode string) string {
	return fmt.Sprintf("%d%s", year, countryCode)
}

func (s *Service) GetCountries(ctx context.Context) ([]dateapi.CountryV3Dto, error) {
	var err error
	ck := "countries"
	resp := s.cache.Get(ck)
	if resp == nil {
		resp, _, err = s.dateapi.CountryAPI.CountryAvailableCountries(ctx).Execute()
		if err != nil {
			log.Error().Err(err).Msg("error when calling `CountryAvailableCountries`")
			return nil, err
		}
		s.cache.Set(ck, resp)
	}
	return resp.([]dateapi.CountryV3Dto), nil
}

func (s *Service) GetCalendar(ctx context.Context, countryCode string, opts CreateCalendarOptions) ([]FormattedDate, error) {
	var err error
	ck := createCacheKey(int32(opts.Year), countryCode)
	resp := s.cache.Get(ck)
	if resp == nil {
		resp, _, err = s.dateapi.PublicHolidayAPI.PublicHolidayPublicHolidaysV3(ctx, int32(opts.Year), countryCode).Execute()
		if err != nil {
			log.Error().Err(err).Msg("error when calling `PublicHolidayNextPublicHolidays`")
			return nil, err
		}
		s.cache.Set(ck, resp)
	}
	ckNext := createCacheKey(int32(opts.Year)+1, countryCode)
	respNext := s.cache.Get(ckNext)
	if respNext == nil {
		respNext, _, err = s.dateapi.PublicHolidayAPI.PublicHolidayPublicHolidaysV3(ctx, int32(opts.Year)+1, countryCode).Execute()
		if err != nil {
			log.Error().Err(err).Msg("error when calling `PublicHolidayNextPublicHolidays`")
			return nil, err
		}
		s.cache.Set(ckNext, respNext)
	}
	rn := respNext.([]dateapi.PublicHolidayV3Dto)
	r := resp.([]dateapi.PublicHolidayV3Dto)
	if len(rn) > 0 {
		resp = append(r, rn[0])
	}
	formatted, err := createCalendar(r, opts)
	if err != nil {
		log.Error().Err(err).Msg("error when formatting dates")
		return nil, err
	}
	return formatted, nil
}

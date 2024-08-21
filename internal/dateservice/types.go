package dateservice

import (
	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

type Service struct {
	dateapi *dateapi.APIClient
}

type FormattedDate struct {
	IsSqueeze    bool                   `json:"isSqueeze"`
	IsHoliday    bool                   `json:"isHoliday"`
	IsWeekend    bool                   `json:"isWeekend"`
	HolidayName  string                 `json:"holidayName"`
	Timestamp    string                 `json:"timestamp"`
	HolidayTypes []dateapi.HolidayTypes `json:"holidayTypes"`
	GapSize      int                    `json:"gapSize"`
	Week         int                    `json:"week"`
}

type CreateCalendarOptions struct {
	Threshold int
	Year      int
}

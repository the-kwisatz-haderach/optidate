package dateservice

import (
	"time"

	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

type Service struct {
	dateapi *dateapi.APIClient
}

type FormattedDate struct {
	IsSqueeze       bool                   `json:"isSqueeze"`
	IsHoliday       bool                   `json:"isHoliday"`
	IsWeekend       bool                   `json:"isWeekend"`
	HolidayName     string                 `json:"holidayName"`
	Timestamp       string                 `json:"timestamp"`
	HolidayTypes    []dateapi.HolidayTypes `json:"holidayTypes"`
	DaysUntilDayOff int                    `json:"daysUntilDayOff"`
	parsedDate      time.Time
}

package dateservice

import (
	"math"
	"time"

	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

const timeLayout string = "2006-01-02"

func createCalendar(dates []dateapi.PublicHolidayV3Dto, threshold int) ([]FormattedDate, error) {
	fds := make([]FormattedDate, 0, 365)

	nextHoliday, dates := dates[0], dates[1:]
	start := time.Now()
	end := start.AddDate(1, 0, 0)
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		var fd FormattedDate

		fd.Timestamp = d.Format(timeLayout)
		fd.parsedDate = d
		fd.IsWeekend = d.Weekday() == time.Saturday || d.Weekday() == time.Sunday

		if nextHoliday.GetDate() == fd.Timestamp {
			fd.IsHoliday = true
			fd.HolidayName = nextHoliday.GetLocalName()
			fd.HolidayTypes = nextHoliday.GetTypes()
			if len(dates) > 0 {
				nextHoliday, dates = dates[0], dates[1:]
			}
		}

		if !fd.IsWeekend && !fd.IsHoliday {
			// Check how long until next holiday vs until Saturday
			daysUntilWeekend := float64(6 - d.Weekday())
			t, err := time.Parse(timeLayout, nextHoliday.GetDate())
			if err != nil {
				panic(err)
			}
			diff := t.Sub(d)
			daysUntilNextHoliday := float64(math.Ceil(diff.Hours() / 24))
			minDiff := math.Min(daysUntilWeekend, daysUntilNextHoliday)
			fd.DaysUntilDayOff = int(minDiff)

			if fd.DaysUntilDayOff <= threshold {
				if len(fds) > 0 {
					prev := fds[len(fds)-1]
					fd.IsSqueeze = prev.IsSqueeze || prev.IsHoliday || prev.IsWeekend
				} else {
					fd.IsSqueeze = true
				}
			}
		}

		fds = append(fds, fd)
	}

	return fds, nil
}

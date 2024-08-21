package dateservice

import (
	"errors"
	"math"
	"time"

	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

const timeLayout string = "2006-01-02"

func createCalendar(holidays []dateapi.PublicHolidayV3Dto, opts CreateCalendarOptions) ([]FormattedDate, error) {
	currYear := time.Now().UTC().Year()
	if opts.Year < currYear-1 {
		return nil, errors.New("year can't be prior to last year")
	}
	if len(holidays) == 0 {
		return []FormattedDate{}, nil
	}

	fds := make([]FormattedDate, 0, len(holidays)*7)

	currIndex := 0
	t, err := time.Parse(timeLayout, holidays[currIndex].GetDate())
	if err != nil {
		panic(err)
	}

	start := t.UTC().Truncate(24 * time.Hour)
	start = start.AddDate(0, 0, int(start.Weekday()-1)*-1)
	end := time.Date(opts.Year, 12, 31, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7)
	end = end.AddDate(0, 0, int(end.Weekday())*-1)

	weekCount := 0
	hasHoliday := false
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		var currDate FormattedDate
		currDate.Timestamp = d.Format(timeLayout)
		currDate.IsWeekend = d.Weekday() == time.Saturday || d.Weekday() == time.Sunday
		_, w := d.ISOWeek()
		currDate.Week = w

		if len(holidays) > currIndex {
			currHoliday := holidays[currIndex]
			if currHoliday.GetDate() == currDate.Timestamp {
				currDate.IsHoliday = true
				currDate.HolidayName = currHoliday.GetLocalName()
				currDate.HolidayTypes = currHoliday.GetTypes()
				currIndex += 1
				hasHoliday = true
			}
		}

		if !currDate.IsWeekend && !currDate.IsHoliday {
			distanceToPreviousWeekend := float64(d.Weekday())
			distanceToNextWeekend := float64(math.Abs(float64(d.Weekday() - 6)))

			var distanceToPreviousHoliday float64 = 5
			if currIndex > 0 {
				prevHoliday := holidays[currIndex-1]
				t, err := time.Parse(timeLayout, prevHoliday.GetDate())
				if err != nil {
					panic(err)
				}
				diff := t.Sub(d).Abs().Hours()
				distanceToPreviousHoliday = math.Floor(diff / 24)
			}

			var distanceToNextHoliday float64 = 5
			if len(holidays) > currIndex {
				nextHoliday := holidays[currIndex]
				t, err = time.Parse(timeLayout, nextHoliday.GetDate())
				if err != nil {
					panic(err)
				}
				diff := t.Sub(d).Abs().Hours()
				distanceToNextHoliday = math.Floor(diff / 24)
			}

			prevDiff := math.Min(distanceToPreviousHoliday, distanceToPreviousWeekend)
			nextDiff := math.Min(distanceToNextHoliday, distanceToNextWeekend)
			minDiff := math.Max(prevDiff, nextDiff)
			currDate.GapSize = int(minDiff)
			if len(fds) > 0 {
				prevDate := fds[len(fds)-1]
				currDate.GapSize = int(math.Max(float64(currDate.GapSize), float64(prevDate.GapSize)))
			}
			currDate.IsSqueeze = currDate.GapSize <= opts.Threshold
		}

		fds = append(fds, currDate)
		weekCount += 1
		if weekCount == 7 {
			if !hasHoliday {
				fds = fds[0 : len(fds)-7]
			}
			hasHoliday = false
			weekCount = 0
		}
	}

	return fds, nil
}

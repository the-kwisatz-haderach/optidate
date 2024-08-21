package dateservice

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-kwisatz-haderach/optidate/dateapi"
)

func newDefaultOptions() CreateCalendarOptions {
	return CreateCalendarOptions{Threshold: 3, NumberOfDays: 50}
}

func loadTestData(t *testing.T) []dateapi.PublicHolidayV3Dto {
	jsonFile, err := os.Open("./testdata.json")
	if err != nil {
		t.Fatal("error while parsing test data")
	}
	defer jsonFile.Close()
	var dates []dateapi.PublicHolidayV3Dto
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &dates)
	return dates
}

func TestEmptyInput(t *testing.T) {
	var dates []dateapi.PublicHolidayV3Dto
	_, err := createCalendar(dates, newDefaultOptions())
	assert.NoError(t, err)
}

func TestZeroThreshold(t *testing.T) {
	dates := loadTestData(t)
	opts := newDefaultOptions()
	opts.Threshold = 0
	actual, err := createCalendar(dates, opts)
	assert.NoError(t, err)

	for _, a := range actual {
		assert.False(t, a.IsSqueeze)
	}
}

func TestHighThreshold(t *testing.T) {
	dates := loadTestData(t)
	opts := newDefaultOptions()
	opts.Threshold = 5
	actual, err := createCalendar(dates, opts)
	assert.NoError(t, err)
	assert.Equal(t, opts.NumberOfDays, len(actual))

	for _, a := range actual {
		assert.True(t, a.IsSqueeze || a.IsHoliday || a.IsWeekend)
	}
}

func TestTimestampFormat(t *testing.T) {
	dates := loadTestData(t)
	opts := newDefaultOptions()
	opts.Threshold = 5
	actual, err := createCalendar(dates, opts)
	assert.NoError(t, err)

	for _, a := range actual {
		assert.Equal(t, 10, len(a.Timestamp))
	}
}

func TestDaysUntilDayOff(t *testing.T) {
	dates := loadTestData(t)
	opts := newDefaultOptions()
	opts.Threshold = 2
	actual, err := createCalendar(dates, opts)
	assert.NoError(t, err)

	for _, a := range actual {
		if a.IsHoliday || a.IsWeekend {
			assert.Equal(t, 0, a.DistanceToDayOff)
		}
	}
}

func TestOneWeek(t *testing.T) {
	opts := newDefaultOptions()
	opts.Threshold = 2
	opts.NumberOfDays = 8

	var dates = []dateapi.PublicHolidayV3Dto{
		{
			Date:      dateapi.PtrString("2024-01-01"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("New Year's Day")),
			Types:     []dateapi.HolidayTypes{"Public", "Cool"},
		},
		{
			Date:      dateapi.PtrString("2024-01-02"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("Yoyo Tuesday")),
			Types:     []dateapi.HolidayTypes{},
		},
		{
			Date:      dateapi.PtrString("2024-01-05"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("Happy Friday")),
			Types:     []dateapi.HolidayTypes{"Public"},
		},
	}

	var expected = []FormattedDate{
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        false,
			HolidayName:      "New Year's Day",
			Timestamp:        "2024-01-01",
			HolidayTypes:     []dateapi.HolidayTypes{"Public", "Cool"},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        false,
			HolidayName:      "Yoyo Tuesday",
			Timestamp:        "2024-01-02",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        true,
			IsHoliday:        false,
			IsWeekend:        false,
			HolidayName:      "",
			Timestamp:        "2024-01-03",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 2,
		},
		{
			IsSqueeze:        true,
			IsHoliday:        false,
			IsWeekend:        false,
			HolidayName:      "",
			Timestamp:        "2024-01-04",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 2,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        false,
			HolidayName:      "Happy Friday",
			Timestamp:        "2024-01-05",
			HolidayTypes:     []dateapi.HolidayTypes{"Public"},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        false,
			IsWeekend:        true,
			HolidayName:      "",
			Timestamp:        "2024-01-06",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        false,
			IsWeekend:        true,
			HolidayName:      "",
			Timestamp:        "2024-01-07",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        false,
			IsWeekend:        false,
			HolidayName:      "",
			Timestamp:        "2024-01-08",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 5,
		},
	}

	actual, err := createCalendar(dates, opts)
	assert.NoError(t, err)

	for i := range actual {
		assert.Equal(t, expected[i].DistanceToDayOff, actual[i].DistanceToDayOff)
		assert.Equal(t, expected[i].HolidayName, actual[i].HolidayName)
		assert.Equal(t, expected[i].IsHoliday, actual[i].IsHoliday)
		assert.Equal(t, expected[i].IsSqueeze, actual[i].IsSqueeze)
		assert.Equal(t, expected[i].IsWeekend, actual[i].IsWeekend)
		assert.Equal(t, expected[i].Timestamp, actual[i].Timestamp)
	}
}

func TestAnotherWeek(t *testing.T) {
	opts := newDefaultOptions()
	opts.Threshold = 2
	opts.NumberOfDays = 7

	var dates = []dateapi.PublicHolidayV3Dto{
		{
			Date:      dateapi.PtrString("2024-12-21"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("Weekend Before Xmas")),
			Types:     []dateapi.HolidayTypes{"Public"},
		},
		{
			Date:      dateapi.PtrString("2024-12-24"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("Julafton")),
			Types:     []dateapi.HolidayTypes{"Public"},
		},
		{
			Date:      dateapi.PtrString("2024-12-25"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("Juldagen")),
			Types:     []dateapi.HolidayTypes{"Public"},
		},
		{
			Date:      dateapi.PtrString("2024-12-26"),
			LocalName: *dateapi.NewNullableString(dateapi.PtrString("Annandag jul")),
			Types:     []dateapi.HolidayTypes{"Public"},
		},
	}

	var expected = []FormattedDate{
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        true,
			HolidayName:      "Weekend Before Xmas",
			Timestamp:        "2024-12-21",
			HolidayTypes:     []dateapi.HolidayTypes{"Public"},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        false,
			IsWeekend:        true,
			HolidayName:      "",
			Timestamp:        "2024-12-22",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        true,
			IsHoliday:        false,
			IsWeekend:        false,
			HolidayName:      "",
			Timestamp:        "2024-12-23",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 1,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        false,
			HolidayName:      "Julafton",
			Timestamp:        "2024-12-24",
			HolidayTypes:     []dateapi.HolidayTypes{"Public"},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        false,
			HolidayName:      "Juldagen",
			Timestamp:        "2024-12-25",
			HolidayTypes:     []dateapi.HolidayTypes{"Public"},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        true,
			IsWeekend:        false,
			HolidayName:      "Annandag jul",
			Timestamp:        "2024-12-26",
			HolidayTypes:     []dateapi.HolidayTypes{"Public"},
			DistanceToDayOff: 0,
		},
		{
			IsSqueeze:        true,
			IsHoliday:        false,
			IsWeekend:        false,
			HolidayName:      "",
			Timestamp:        "2024-12-27",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 1,
		},
		{
			IsSqueeze:        false,
			IsHoliday:        false,
			IsWeekend:        true,
			HolidayName:      "",
			Timestamp:        "2024-12-28",
			HolidayTypes:     []dateapi.HolidayTypes{},
			DistanceToDayOff: 0,
		},
	}

	actual, err := createCalendar(dates, opts)
	assert.NoError(t, err)

	for i := range actual {
		assert.Equal(t, expected[i].DistanceToDayOff, actual[i].DistanceToDayOff)
		assert.Equal(t, expected[i].HolidayName, actual[i].HolidayName)
		assert.Equal(t, expected[i].IsHoliday, actual[i].IsHoliday)
		assert.Equal(t, expected[i].IsSqueeze, actual[i].IsSqueeze)
		assert.Equal(t, expected[i].IsWeekend, actual[i].IsWeekend)
		assert.Equal(t, expected[i].Timestamp, actual[i].Timestamp)
	}
}

package converttyme

import (
	"time"

	"github.com/seiyab/tyme"
)

// FromLocalYear is used to convert from LocalYear
func FromLocalYear(year tyme.LocalYear) IntermediateTime {
	dateTime := time.Date(year.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
	return IntermediateTime(dateTime)
}

// FromLocalMonth is used to convert from LocalMonth
func FromLocalMonth(month tyme.LocalMonth) IntermediateTime {
	dateTime := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.Local)
	return IntermediateTime(dateTime)
}

// FromLocalDate is used to convert from LocalDate
func FromLocalDate(date tyme.LocalDate) IntermediateTime {
	dateTime := time.Date(date.Year(), date.Month(), date.Date(), 0, 0, 0, 0, time.Local)
	return IntermediateTime(dateTime)
}

// FromLocalHour is used to convert from LocalHour
func FromLocalHour(hour tyme.LocalHour) IntermediateTime {
	dateTime := time.Date(hour.Year(), hour.Month(), hour.Date(), hour.Hour(), 0, 0, 0, time.Local)
	return IntermediateTime(dateTime)
}
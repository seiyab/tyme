package tyme

import (
	"fmt"
	"time"
)

// LocalHour represents year, month and daae without Location
// e.g) January 2, 2006
type LocalHour struct {
	date LocalDate
	hour int
}

// NewLocalHour returns instance of LocalHour with given year, month and date
func NewLocalHour(year int, month time.Month, date int, hour int) LocalHour {
	return LocalHour{
		date: NewLocalDate(year, month, date),
		hour: hour,
	}
}

// Year returns year
func (h *LocalHour) Year() int {
	return h.date.Year()
}

// Month returns month
func (h *LocalHour) Month() time.Month {
	return h.date.Month()
}

// Date returns number of date
func (h *LocalHour) Date() int {
	return h.date.Date()
}

// Hour returns number of date
func (h *LocalHour) Hour() int {
	return h.hour
}

// String returns "yyyy-mm-ddThh" format
func (d LocalHour) String() string {
	return fmt.Sprintf("%04d-%02d-%02dT%02d", d.Year(), int(d.Month()), d.Date(), d.Hour())
}

package tyme

import (
	"fmt"
	"time"
)

// LocalMonth represents year and month without Location
// e.g.) January, 2006
type LocalMonth struct {
	year  LocalYear
	month time.Month
}

// NewLocalMonth returns instance of LocalMonth with given year and month
func NewLocalMonth(year int, month time.Month) LocalMonth {
	return LocalMonth{
		year:  NewLocalYear(year),
		month: month,
	}
}

// Year returns year
func (m *LocalMonth) Year() int {
	return m.year.Year()
}

// Month returns time.Month
func (m *LocalMonth) Month() time.Month {
	return m.month
}

// String returns "yyyy-mm" format
func (m LocalMonth) String() string {
	return fmt.Sprintf("%04d-%02d", m.Year(), int(m.month))
}

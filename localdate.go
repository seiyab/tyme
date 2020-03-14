package tyme

import (
	"fmt"
	"time"
)

// LocalDate represents year, month and daae without Location
// e.g) January 2, 2006
type LocalDate struct {
	month LocalMonth
	date  int
}

// NewLocalDate returns instance of LocalDate with given year, month and date
func NewLocalDate(year int, month time.Month, date int) LocalDate {
	return LocalDate{
		month: NewLocalMonth(year, month),
		date:  date,
	}
}

// Year returns year
func (d *LocalDate) Year() int {
	return d.month.Year()
}

// Month returns month
func (d *LocalDate) Month() time.Month {
	return d.month.Month()
}

// Date returns number of date
func (d *LocalDate) Date() int {
	return d.date
}

// String returns "yyyy-mm-dd" format
func (d LocalDate) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year(), int(d.Month()), d.Date())
}

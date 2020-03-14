package converttyme

import (
	"time"

	"github.com/seiyab/tyme"
)

// ToTime is used to convert to time.Time
func (i IntermediateTime) ToTime() time.Time {
	return time.Time(i)
}

// ToLocalYear is used to convert to LocalYear
func (i IntermediateTime) ToLocalYear() tyme.LocalYear {
	return tyme.NewLocalYear(time.Time(i).Year())
}

// ToLocalMonth is used to convert to LocalMonth
func (i IntermediateTime) ToLocalMonth() tyme.LocalMonth {
	t := time.Time(i)
	return tyme.NewLocalMonth(t.Year(), t.Month())
}

// ToLocalDate is used to convert to LocalDate
func (i IntermediateTime) ToLocalDate() tyme.LocalDate {
	t := time.Time(i)
	return tyme.NewLocalDate(t.Year(), t.Month(), t.Day())
}

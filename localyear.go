package tyme

import "fmt"

// LocalYear represents year without Location
// e.g.) 2006
type LocalYear struct {
	year int
}

// NewLocalYear returns instance of LocalYear
func NewLocalYear(year int) LocalYear {
	return LocalYear{year: year}
}

// Year returns number of year
func (y *LocalYear) Year() int {
	return y.year
}

// String returns "yyyy" format
func (y LocalYear) String() string {
	return fmt.Sprintf("%04d", y.Year())
}
